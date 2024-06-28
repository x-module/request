package main

import (
	"fmt"

	"github.com/x-module/request"
)

func main() {
	// Create new request instance
	req := request.NewRequest()
	req.Method("GET")

	// Define target URL
	req.URL("http://httpbin.org/headers")

	// Set a new header field
	req.SetHeader("Client", "request")

	// Set sample context data
	req.Context.Set("foo", "bar")
	req.Context.Set("bar", "baz")

	// Set sample body as string
	req.BodyString("hello, request!")

	// Output all context data
	fmt.Println(req.Context.GetAll())

	// Perform the request
	res, err := req.Do()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	// Set sample context data
	fmt.Println(req.Context.GetString("foo"))
	fmt.Println(req.Context.GetString("bar"))

	// Reads the whole body and returns it as string
	fmt.Printf("Body: %s", res.String())
}
