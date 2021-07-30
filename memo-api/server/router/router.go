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
	r.Router.GET("/deck/{deckName}/card", c.Get())
	r.Router.POST("/deck/{deckName}/card", c.Post())
	r.Router.PUT("/deck/{deckName}/card/{front}", c.Update())
	r.Router.DELETE("/deck/{deckName}/card/{front}", c.Delete())

}

func (r *router) deckRoutes() {
	repo := repository.NewDeck(r.Mongo.Client)
	d := controller.NewDeck(repo)
	r.Router.GET("/deck/{name}", d.Get())
	r.Router.POST("/deck", d.Post())
	r.Router.DELETE("/deck/{name}", d.Delete())
	r.Router.PUT("/deck/{name}", d.Update())
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
