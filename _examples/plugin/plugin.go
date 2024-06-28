package main

import (
	"fmt"
	"net/url"

	"github.com/x-module/request"
	"github.com/x-module/request/context"
	"github.com/x-module/request/plugin"
	"github.com/x-module/request/plugins/headers"
)

func main() {
	// Create a new client
	cli := request.New()

	// Define a custom header
	cli.Use(headers.Set("Token", "s3cr3t"))

	// Create a request plugin to define the URL
	cli.Use(plugin.NewRequestPlugin(func(ctx *context.Context, h context.Handler) {
		u, _ := url.Parse("http://httpbin.org/headers")
		ctx.Request.URL = u
		h.Next(ctx)
	}))

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
