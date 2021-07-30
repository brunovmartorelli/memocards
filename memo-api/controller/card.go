package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/brunovmartorelli/memo-api/domain"
	"github.com/brunovmartorelli/memo-api/repository"
	"github.com/valyala/fasthttp"
	"gopkg.in/validator.v2"
)

type Card struct {
	repository repository.CardRepository
}

type WriteCardBody struct {
	Front string `validate:"max=100"`
	Back  string `validate:"max=100"`
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

func (c *Card) Update() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		frontEncoded := ctx.UserValue("front").(string)
		front, err := url.QueryUnescape(frontEncoded)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		log.Println(front)

		b := ctx.Request.Body()
		card := WriteCardBody{}
		if err := json.Unmarshal(b, &card); err != nil {
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		if err := validator.Validate(card); err != nil {
			log.Println(err)
			ctx.SetBodyString(err.Error())
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		domainCard := domain.Card{
			Front: card.Front,
			Back:  card.Back,
		}

		count, err := c.repository.Update(front, domainCard)
		if err != nil {
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		if count < 1 {
			log.Println("Not modified")
			ctx.SetStatusCode(fasthttp.StatusNotModified)
			return
		}

		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.SetBodyString(`{"message": "Carta atualizada com sucesso."}`)
	})
}

func (c *Card) Post() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		deckEncoded := ctx.UserValue("deckName").(string)
		deckName, err := url.QueryUnescape(deckEncoded)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		b := ctx.Request.Body()
		card := WriteCardBody{}
		if err := json.Unmarshal(b, &card); err != nil {
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		if err := validator.Validate(card); err != nil {
			log.Println(err)
			ctx.SetBodyString(err.Error())
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		_, geterr := c.repository.GetByFront(card.Front, deckName)
		if geterr == nil {
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusConflict)
			ctx.SetBodyString(fmt.Sprintf("A carta %s já existe.", card.Front))

			return
		}

		domainCard := domain.Card{
			Front: card.Front,
			Back:  card.Back,
		}

		if err := c.repository.Create(deckName, domainCard); err != nil {
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.SetBodyString(`{"message": Carta criada com sucesso."}`)
	})
}

func (c *Card) Delete() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		frontEncoded := ctx.UserValue("front").(string)
		front, err := url.QueryUnescape(frontEncoded)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		log.Println(front)

		count, err := c.repository.Delete(front)
		if err != nil {
			ctx.SetBodyString(err.Error())
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		if count == 0 {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBodyString(fmt.Sprintf("%s Not Found.", front))
			return
		}

		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.SetBodyString(`{"message": "Carta deletada com sucesso."}`)
	})
}
