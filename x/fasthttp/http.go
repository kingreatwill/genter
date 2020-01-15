package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	fooHandlerFunc := func(ctx *fasthttp.RequestCtx) {
		// set some headers and status code first
		//ctx.SetContentType("foo/bar")
		ctx.SetStatusCode(fasthttp.StatusOK)

		// then write the first part of body
		fmt.Fprintf(ctx, "this is the first part of body\n")

		// then set more headers
		ctx.Response.Header.Set("Foo-Bar", "baz")

		// then write more body
		fmt.Fprintf(ctx, "this is the second part of body\n")

		// then override already written body
		//ctx.SetBody([]byte("this is completely new body contents"))

		// then update status code
		ctx.SetStatusCode(fasthttp.StatusNotFound)

		// basically, anything may be updated many times before
		// returning from RequestHandler.
		//
		// Unlike net/http fasthttp doesn't put response to the wire until
		// returning from RequestHandler.
	}
	// the corresponding fasthttp code
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/foo":
			fooHandlerFunc(ctx)
		case "/bar":
			//barHandlerFunc(ctx)
		case "/baz":
			//bazHandler.HandlerFunc(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	fasthttp.ListenAndServe(":8181", m)
}
