package udshttp

import (
	"io/ioutil"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewUnixClient(defaultUnixAddr)
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
