package controller

import (
	"github.com/valyala/fasthttp"
)

type Card struct {
}

func NewCard() *Card {
	return &Card{}
}

func (c *Card) Get() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.SetBodyString(`{"message": "Card"}`)
	})
}

func (c *Card) Post() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")

	})
}
