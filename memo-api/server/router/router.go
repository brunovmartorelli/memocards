package router

import (
	"github.com/brunovmartorelli/memo-api/controller"
	"github.com/brunovmartorelli/memo-api/repository"

	fhr "github.com/fasthttp/router"
)

type router struct {
	Router *fhr.Router
	Mongo  *repository.Mongo
}

func (r *router) cardRoutes() {
	repo := repository.NewCard(r.Mongo.Client)
	c := controller.NewCard(repo)
	r.Router.GET("/decks/{deckName}/cards", c.List())
	r.Router.POST("/decks/{deckName}/cards", c.Post())
	r.Router.PUT("/decks/{deckName}/cards/{front}", c.Update())
	r.Router.DELETE("/decks/{deckName}/cards/{front}", c.Delete())

}

func (r *router) deckRoutes() {
	repo := repository.NewDeck(r.Mongo.Client)
	d := controller.NewDeck(repo)
	r.Router.GET("/decks", d.List())
	r.Router.GET("/decks/{name}", d.Get())
	r.Router.POST("/decks", d.Post())
	r.Router.DELETE("/decks/{name}", d.Delete())
	r.Router.PUT("/decks/{name}", d.Update())
}

func New(m *repository.Mongo) *router {
	r := &router{
		Router: fhr.New(),
		Mongo:  m,
	}
	r.cardRoutes()
	r.deckRoutes()
	return r
}
