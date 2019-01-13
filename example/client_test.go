package example

import (
	"io/ioutil"
	"testing"

	"github.com/nuczzz/udshttp"
)

func TestClient(t *testing.T) {
	client := udshttp.NewUnixClient(udshttp.DefaultUnixAddr)
	resp, err := client.Get("http://unix/index")
	if err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
