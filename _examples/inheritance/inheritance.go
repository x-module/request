package main

import (
	"fmt"

	"github.com/x-module/request"
	"github.com/x-module/request/plugins/headers"
)

func main() {
	// Create a parent client
	parent := request.New()

	// Define default URL
	parent.URL("http://httpbin.org")

	// Define a custom header via parent client
	parent.Use(headers.Set("API-Token", "s3cr3t"))

	// Create a new client
	cli := request.New()

	// Bind parent client
	cli.UseParent(parent)

	// Create a new request based on the current client
	req := cli.Request()

	// Perform the request
	res, err := req.Path("/post").JSON(map[string]string{"foo": "bar"}).Send()
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
