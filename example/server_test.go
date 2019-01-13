package example

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nuczzz/udshttp"
)

func echo(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "handler test ok")
}

func TestServer(t *testing.T) {
	server := udshttp.NewUnixServer(udshttp.DefaultUnixAddr)
	server.HandlerFunc("GET", "/echo", echo)
	t.Log(server.ListenAndServe())
}
