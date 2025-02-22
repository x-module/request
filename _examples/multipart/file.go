package main

import (
	"fmt"
	"os"

	"github.com/x-module/request"
	"github.com/x-module/request/plugins/multipart"
)

func main() {
	// Create a new client
	cli := request.New()

	// Define the generic base URL
	cli.URL("http://httpbin.org/post")

	// Create a new request
	req := cli.Request()

	// Read a file from disk and post it
	file, _ := os.Open("LICENSE")
	defer file.Close()

	// Register the multipart plugin at request specific middleware level
	req.Use(multipart.File("license", file))

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

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}
