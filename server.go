package udshttp

import (
	"net"
	"net/http"
)

type UnixServer struct {
	addr   string
	router *Router
}

func NewUnixServer(addr string, r *Router) *UnixServer {
	if addr == "" {
		addr = defaultUnixAddr
	}
	if r == nil {
		r = NewRouter()
	}

	return &UnixServer{
		addr:   addr,
		router: r,
	}
}

func (us *UnixServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	us.router.match(req.Method, req.URL.Path)(resp, req)
}

func (us *UnixServer) ListenAndServe() error {
	ln, err := net.Listen(defaultUnixNetwork, us.addr)
	if err != nil {
		return err
	}
	server := &http.Server{
		Handler: us,
	}
	return server.Serve(ln)
}
