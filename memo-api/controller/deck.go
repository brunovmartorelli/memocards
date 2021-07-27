package controller

import (
	"encoding/json"
	"log"

	"github.com/brunovmartorelli/memo-api/domain"
	"github.com/brunovmartorelli/memo-api/repository"
	"github.com/valyala/fasthttp"
	"gopkg.in/validator.v2"
)

type PostDeckBody struct {
	Name        string `validate:"nonzero, max=20"`
	Description string `validate:"max=40"`
}

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
		b := ctx.Request.Body()
		deck := PostDeckBody{}
		if err := json.Unmarshal(b, &deck); err != nil {
			//FIXME: error handling
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		if err := validator.Validate(deck); err != nil {
			log.Println(err)
			ctx.SetBodyString(err.Error())
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		domainDeck := domain.Deck{
			Name:        deck.Name,
			Description: deck.Description,
		}

		if err := d.repository.Create(domainDeck); err != nil {
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.SetBodyString(`{"message": "Deck Criado com sucesso."}`)
	})
}
