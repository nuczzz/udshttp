package udshttp

import (
	"fmt"
	"net/http"
	"testing"
)

func HandlerTest(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "handler test ok")
}

func myRouter() *Router {
	r := NewRouter()
	r.Get("/test", HandlerTest)
	return r
}

func TestServer(t *testing.T) {
	router := myRouter()
	server := NewUnixServer(defaultUnixAddr, router)
	t.Log(server.ListenAndServe())
}
