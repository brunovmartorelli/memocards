package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/brunovmartorelli/memo-api/domain"
	"github.com/brunovmartorelli/memo-api/domain/entities"
	"github.com/brunovmartorelli/memo-api/repository"
	"github.com/valyala/fasthttp"
	"gopkg.in/validator.v2"
)

type PostDeckBody struct {
	Name        string `validate:"nonzero, max=20"`
	Description string `validate:"max=100"`
}

type PutDeckBody struct {
	Name        string `validate:"max=40"`
	Description string `validate:"max=100"`
}

type Deck struct {
	repository repository.DeckRepository
	usecase    *domain.UseCase
}

func NewDeck(r repository.DeckRepository, u *domain.UseCase) *Deck {
	return &Deck{
		repository: r,
		usecase:    u,
	}
}

func (d *Deck) List() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		qa := ctx.Request.URI().QueryArgs()
		cards := true
		if qa.Has("cards") {
			cards = qa.GetBool("cards")
		}

		decks, err := d.repository.List(cards)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetBodyString(err.Error())
			return
		}
		body, jerr := json.Marshal(decks)
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

func (d *Deck) Get() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		deckEncoded := ctx.UserValue("name").(string)
		deckName, err := url.QueryUnescape(deckEncoded)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		deck, geterr := d.repository.GetByName(deckName)
		if geterr != nil {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBodyString(fmt.Sprintf("%s Not Found.", deckName))
			return
		}

		body, jerr := json.Marshal(deck)
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

func (d *Deck) GetStudy() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		deckEncoded := ctx.UserValue("name").(string)
		deckName, err := url.QueryUnescape(deckEncoded)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		_, geterr := d.repository.GetByName(deckName)
		if geterr != nil {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBodyString(fmt.Sprintf("%s Not Found.", deckName))
			return
		}

		cards, _ := d.usecase.FilterCardsToStudy(deckName, time.Now())

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

func (d *Deck) Update() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		name := ctx.UserValue("name").(string)
		b := ctx.Request.Body()
		deck := PutDeckBody{}
		if err := json.Unmarshal(b, &deck); err != nil {
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

		domainDeck := entities.Deck{
			Name:        deck.Name,
			Description: deck.Description,
		}

		_, err := d.repository.Update(name, domainDeck)
		if err != nil {
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.SetBodyString(`{"message": "Deck atualizado com sucesso."}`)
	})
}

func (d *Deck) Post() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		b := ctx.Request.Body()
		deck := PostDeckBody{}
		if err := json.Unmarshal(b, &deck); err != nil {
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

		_, err := d.repository.GetByName(deck.Name)
		if err == nil {
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusConflict)
			ctx.SetBodyString(fmt.Sprintf("O deck %s já existe.", deck.Name))
			return
		}

		domainDeck := entities.Deck{
			Name:        deck.Name,
			Description: deck.Description,
			Cards:       &[]entities.Card{},
		}

		if err := d.repository.Create(domainDeck); err != nil {
			log.Println(err)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.SetBodyString(`{"message": "Deck criado com sucesso."}`)
	})
}

func (d *Deck) Delete() fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		name := ctx.UserValue("name").(string)
		count, err := d.repository.Delete(name)
		if err != nil {
			ctx.SetBodyString(err.Error())
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		if count == 0 {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBodyString(fmt.Sprintf("%s Not Found.", name))
			return
		}

		ctx.Response.Header.Add("Content-Type", "application/json; charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.SetBodyString(`{"message": "Deck deletado com sucesso."}`)
	})
}
