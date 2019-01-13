package udshttp

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
)

const (
	DefaultUnixNetwork = "unix"
	DefaultUnixAddr    = "/var/run/unix_server.sock"
)

type UnixServer struct {
	addr string
	*httprouter.Router
}

func NewUnixServer(addr string) *UnixServer {
	if addr == "" {
		addr = DefaultUnixAddr
	}

	return &UnixServer{
		addr:   addr,
		Router: httprouter.New(),
	}
}

func (us *UnixServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	us.Router.ServeHTTP(resp, req)
}

func (us *UnixServer) ListenAndServe() error {
	ln, err := net.Listen(DefaultUnixNetwork, us.addr)
	if err != nil {
		return err
	}
	server := &http.Server{
		Handler: us,
	}
	return server.Serve(ln)
}
