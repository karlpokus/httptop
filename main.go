package main

import (
	"fmt"
	"bufio"
	"strings"
	"regexp"
	"text/tabwriter"
	"os"
)

var data = `T 127.0.0.1:38802 -> 127.0.0.1:3000 [AP] #4
GET /tenant/123/thing HTTP/1.1.
--
T 127.0.0.1:3000 -> 127.0.0.1:38802 [AP] #6
HTTP/1.1 200 OK.
--
T 127.0.0.1:38804 -> 127.0.0.1:3000 [AP] #14
POST /tenant/123/otherthing HTTP/1.1.
--
T 127.0.0.1:3000 -> 127.0.0.1:38804 [AP] #16
HTTP/1.1 200 OK.
--
T 172.20.16.1:49768 -> 172.20.16.44:3000 [AP] #24
GET /health HTTP/1.1.
--
T 127.0.0.1:38802 -> 127.0.0.1:3000 [AP] #4
GET /tenant/123/thing HTTP/1.1.
--
`

var isRequest = regexp.MustCompile(`^GET|^POST|^PUT|^HEAD`)

type request struct {
	Path string
	Count int
}

func main() {
	r := strings.NewReader(data)
	scanner := bufio.NewScanner(r)
	var requests []*request
	for scanner.Scan() {
		text := scanner.Text()
		if isRequest.MatchString(text) {
			if len(requests) == 0 {
				requests = append(requests, &request{Path:text, Count:1})
				continue
			}
			for _, r := range requests {
				if r.Path == text {
					r.Count++
					break
				}
				requests = append(requests, &request{Path:text, Count:1})
			}
		}
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, r := range requests {
		fmt.Fprintf(w, "%s\t%d\n", r.Path, r.Count)
	}
	w.Flush()
}
