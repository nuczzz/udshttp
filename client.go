package udshttp

import (
	"context"
	"net"
	"net/http"
)

type UnixClient struct {
	*http.Client
}

func NewUnixClient(serverSock string) *UnixClient {
	if serverSock == "" {
		serverSock = defaultUnixAddr
	}

	return &UnixClient{
		Client: &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", serverSock)
				},
			},
		},
	}
}
