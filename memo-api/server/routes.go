package server

import (
	"github.com/brunovmartorelli/memo-api/controller"
	"github.com/brunovmartorelli/memo-api/server/router"
)

func CardRoutes(r *router.Router, c *controller.Card) {
	r.Router.GET("/card", c.Get())
}

func DeckRoutes(r *router.Router, d *controller.Deck) {
	r.Router.GET("/deck", d.Get())
}
