package main

import (
	"fmt"

	"github.com/x-module/request"
	"github.com/x-module/request/plugins/url"
)

func main() {
	// Create a new client
	cli := request.New()

	// Define the base URL
	cli.Use(url.BaseURL("http://httpbin.org"))

	// Define the path with dynamic param
	// Use the :<name> notation
	cli.Use(url.Path("/:resource"))

	// Define the path value
	cli.Use(url.Param("resource", "headers"))

	// Perform the request
	res, err := cli.Request().Send()
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
