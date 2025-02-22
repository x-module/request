package main

import (
	"fmt"

	"github.com/x-module/request"
)

func main() {
	// Create a new client
	cli := request.New()

	// Create a new request based on the current client
	req := cli.Request().URL("http://httpbin.org/get")

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

	// Read the whole body buffer and return it as string
	fmt.Printf("Body: %s", res.String())
}
