package router

import (
	"github.com/brunovmartorelli/memo-api/controller"
	fhr "github.com/fasthttp/router"
)

type router struct {
	Router *fhr.Router
}

func (r *router) cardRoutes(c *controller.Card) {
	r.Router.GET("/card", c.Get())
	r.Router.POST("/card", c.Post())
}

func (r *router) deckRoutes(d *controller.Deck) {
	r.Router.GET("/deck", d.Get())
}

//TODO: Remover a instânciação dos controllers daqui, receber como parâmetros
func (r *router) Routes() {
	cardcontroller := controller.NewCard()
	deckcontroller := controller.NewDeck()
	r.cardRoutes(cardcontroller)
	r.deckRoutes(deckcontroller)

}

func New() *router {
	r := &router{
		Router: fhr.New(),
	}
	return r
}
