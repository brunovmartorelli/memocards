package controller

import (
	"encoding/json"
	"log"

	"github.com/brunovmartorelli/memo-api/domain"
	"github.com/brunovmartorelli/memo-api/repository"
	"github.com/valyala/fasthttp"
)

type Deck struct {
	repository repository.DeckRepository
}

func NewDeck(r repository.DeckRepository) *Deck {
	return &Deck{
		repository: r,
	}
}

func (d *Deck) Get() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.SetBodyString(`{"message": "Deck"}`)
	})
}

func (d *Deck) Post() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		//TODO: Validate deck body
		b := ctx.Request.Body()
		deck := domain.Deck{}
		if err := json.Unmarshal(b, &deck); err != nil {
			//FIXME: error handling
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		if err := d.repository.Create(deck); err != nil {
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.SetBodyString(`{"message": "Deck Criado com sucesso."}`)
	})
}
