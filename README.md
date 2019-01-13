# udshttp
HTTP with unix domain socket.

Server code example:
```
package main

import (
	"fmt"
	"net/http"

	"github.com/nuczzz/udshttp"
)

func echo(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "handler test ok")
}

func main() {
	server := udshttp.NewUnixServer(udshttp.DefaultUnixAddr)
	server.HandlerFunc("GET", "/echo", echo)
	fmt.Println(server.ListenAndServe())
}
```
Client code example:
```
package main

import (
	"io/ioutil"

	"github.com/nuczzz/udshttp"
)

func main() {
	client := udshttp.NewUnixClient(udshttp.DefaultUnixAddr)
	resp, err := client.Get("http://unix/index")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
```