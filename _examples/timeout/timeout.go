package main

import (
	"fmt"
	"time"

	"github.com/x-module/request"
	"github.com/x-module/request/plugins/timeout"
)

func main() {
	// Create a new client
	cli := request.New()

	// Define the max timeout for the whole HTTP request
	cli.Use(timeout.Request(10 * time.Second))

	// Define dial specific timeouts
	cli.Use(timeout.Dial(5*time.Second, 30*time.Second))

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
