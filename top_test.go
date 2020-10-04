package httptop

import (
	"testing"
  "strings"

  "github.com/karlpokus/bufw"
)

var data = `T 127.0.0.1:38802 -> 127.0.0.1:3000 [AP] #4
GET /tenant/1 HTTP/1.1.
--
T 127.0.0.1:3000 -> 127.0.0.1:38802 [AP] #6
HTTP/1.1 200 OK.
--
T 127.0.0.1:38802 -> 127.0.0.1:3000 [AP] #4
GET /tenant/1 HTTP/1.1.
--
`

func TestStart(t *testing.T) {
  r := strings.NewReader(data)
  w := bufw.New()
  go Start(r, w)
  err := w.Wait()
  if err != nil {
    t.Fatal("bufw timeout")
  }
  want := "2 GET /tenant/1 HTTP/1.1."
  got := w.String()
  if want != got {
    t.Fatalf("wanted %s, got %s", want, got)
  }
}
