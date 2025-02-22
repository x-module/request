package main

import (
	"fmt"
	"net/http"

	"github.com/x-module/request"
	"github.com/x-module/request/plugins/transport"
)

func main() {
	// Create a new client
	cli := request.New()

	// Define the default HTTP transport
	cli.Use(transport.Set(http.DefaultTransport))

	// Perform the request
	res, err := cli.Request().URL("http://httpbin.org/headers").Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}
