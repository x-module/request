package main

import (
	"fmt"

	"github.com/x-module/request"
	"github.com/x-module/request/context"
	"github.com/x-module/request/utils"
)

func main() {
	// Create a new client
	cli := request.New()

	// Creates a new request based on the current client
	req := cli.Request().URL("http://httpbin.org/get")

	// Attach a request midddleware function to intercept the request.
	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		// If host matches, intercept the request
		if ctx.Request.URL.Host == "httpbin.org" {
			ctx.Response.StatusCode = 200
			utils.WriteBodyString(ctx.Response, "intercepted\n")
		}
		h.Stop(ctx)
	})

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
