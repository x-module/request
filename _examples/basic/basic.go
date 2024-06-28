package main

import (
	"fmt"

	"github.com/x-module/request"
)

func main() {
	// Create a new client
	cli := request.New()

	// Define base URL
	cli.BaseURL("http://httpbin.org")

	// Create a new request based on the current client
	req := cli.Request()

	// Define the URL path at request level
	req.Path("/headers")

	// Set a new header field
	req.SetHeader("Client", "request")

	// Perform the request
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	// Reads the whole body and returns it as string
	fmt.Printf("Body: %s", res.String())
}
