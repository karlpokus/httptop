package httptop

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

var isRequest = regexp.MustCompile(`^GET|^POST|^PUT|^HEAD`)

type request struct {
	Path  string
	Count int
}

func Start(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	var requests []*request
	for scanner.Scan() {
		text := scanner.Text()
		if isRequest.MatchString(text) {
			if len(requests) == 0 {
				requests = append(requests, &request{Path: text, Count: 1})
				continue
			}
			for _, r := range requests {
				if r.Path == text {
					r.Count++
					break
				}
				requests = append(requests, &request{Path: text, Count: 1})
			}
		}
	}
	for _, r := range requests {
		fmt.Fprintf(w, "%d %s\n", r.Count, r.Path)
	}
}
