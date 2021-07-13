package controller

import (
	"github.com/valyala/fasthttp"
)

type Card struct {
}

func (c *Card) Get(ctx *fasthttp.RequestCtx) fasthttp.RequestHandler {
	ctx.Response.SetBodyString("Hello World")
	return func(ctx *fasthttp.RequestCtx) {

	}
}
