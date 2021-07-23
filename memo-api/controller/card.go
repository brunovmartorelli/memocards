package controller

import (
	"github.com/brunovmartorelli/memo-api/repository"
	"github.com/valyala/fasthttp"
)

type Card struct {
	repository repository.CardRepository
}

func NewCard(r repository.CardRepository) *Card {
	return &Card{
		repository: r,
	}
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
