package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/brunovmartorelli/memo-api/domain"
	"github.com/brunovmartorelli/memo-api/domain/entities"
	"github.com/brunovmartorelli/memo-api/repository"
	"github.com/valyala/fasthttp"
	"gopkg.in/validator.v2"
)

type Card struct {
	repository repository.CardRepository
	usecase    *domain.UseCase
}

type WriteCardBody struct {
	Front string `validate:"max=100"`
	Back  string `validate:"max=100"`
}

func NewCard(r repository.CardRepository, u *domain.UseCase) *Card {
	return &Card{
		repository: r,
		usecase:    u,
	}
}

func (c *Card) List() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		deckEncoded := ctx.UserValue("deckName").(string)
		deckName, err := url.QueryUnescape(deckEncoded)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		cards, err := c.repository.List(deckName)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetBodyString(err.Error())
			return
		}

		body, jerr := json.Marshal(cards)
		if jerr != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetBodyString(jerr.Error())
			return
		}

		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString(string(body))
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

		domainCard := entities.Card{
			Front: card.Front,
			Back:  card.Back,
		}

		count, err := c.repository.Update(front, deckName, domainCard)
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

		foundCard, geterr := c.repository.GetByFront(card.Front, deckName)
		if geterr == nil && foundCard != nil {
			log.Println(geterr)
			ctx.SetStatusCode(fasthttp.StatusConflict)
			ctx.SetBodyString(fmt.Sprintf("A carta %s j?? existe.", card.Front))

			return
		}

		domainCard := entities.Card{
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

		deckEncoded := ctx.UserValue("deckName").(string)
		deckName, err := url.QueryUnescape(deckEncoded)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		count, err := c.repository.Delete(front, deckName)
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

func (c *Card) UpdateScore() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		frontEncoded := ctx.UserValue("front").(string)
		front, ferr := url.QueryUnescape(frontEncoded)
		if ferr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		deckEncoded := ctx.UserValue("deckName").(string)
		deckName, derr := url.QueryUnescape(deckEncoded)
		if derr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		resetData := ctx.QueryArgs().Peek("reset")
		var reset bool
		err := json.Unmarshal(resetData, &reset)
		reset = (err == nil)
		fmt.Println(reset)
		newScore, err := c.usecase.UpdateCardScore(front, deckName, reset)

		if err != nil {
			if e, ok := err.(domain.NotFoundError); ok {
				ctx.SetBodyString(fmt.Sprintf("Could not find deck %s or card %s: %s", deckName, front, e.Error()))
				ctx.SetStatusCode(fasthttp.StatusNotFound)
				return
			}

			ctx.SetBodyString(fmt.Sprintf("Could not update card %s from deck %s: %s", front, deckName, err.Error()))
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return

		}

		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		response := fmt.Sprintf(`{"message": "Score atualizado: %d."}`, newScore)
		ctx.Response.SetBodyString(response)
	})
}
