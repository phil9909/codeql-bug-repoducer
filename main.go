package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
)

const testHost = "tls11.browserleaks.com" // any host which is not reachable via TLS1.2 or above, "tls-v1-1.badssl.com:1011" should also work

func main() {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{},
		},
	}

	resp, err := client.Get(fmt.Sprintf("https://%s/", testHost))
	if err == nil {
		panic(fmt.Sprintf("this should not work, but got %q", resp.Status))
	}

	errString := err.Error()
	if !strings.Contains(errString, "tls: protocol version not supported") && !strings.Contains(errString, "tls: server selected unsupported protocol version ") {
		panic("unexpected error: " + errString)
	}

	fmt.Println("[DEBUG] got expected error: ", err)
}
