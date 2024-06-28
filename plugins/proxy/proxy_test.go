package proxy

import (
	"net/http"
	"strings"
	"testing"

	"github.com/nbio/st"

	"github.com/x-module/request/context"
)

func TestProxy(t *testing.T) {
	ctx := context.New()
	ctx.Request.URL.Scheme = "http"

	fn := newHandler()
	servers := map[string]string{"http": "http://localhost:3128"}

	Set(servers).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)

	transport := ctx.Client.Transport.(*http.Transport)
	url, err := transport.Proxy(ctx.Request)

	st.Expect(t, err, nil)
	st.Expect(t, url.Host, "localhost:3128")
	st.Expect(t, url.Scheme, "http")
}

func TestProxyParseError(t *testing.T) {
	ctx := context.New()
	ctx.Request.URL.Scheme = "http"

	fn := newHandler()
	servers := map[string]string{"http": "://"}

	Set(servers).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)

	transport := ctx.Client.Transport.(*http.Transport)
	_, err := transport.Proxy(ctx.Request)

	st.Expect(t, strings.Contains(err.Error(), "missing protocol scheme"), true)
}

type handler struct {
	fn     context.Handler
	called bool
}

func newHandler() *handler {
	h := &handler{}
	h.fn = context.NewHandler(func(c *context.Context) {
		h.called = true
	})
	return h
}
