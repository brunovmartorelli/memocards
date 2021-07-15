package server

import (
	"log"

	"github.com/brunovmartorelli/memo-api/controller"
	"github.com/brunovmartorelli/memo-api/server/router"
	"github.com/valyala/fasthttp"
)

type server struct {
	httpServer *fasthttp.Server
}

func New(c *controller.Card, d *controller.Deck) *server {
	r := router.New()
	CardRoutes(r, c)
	DeckRoutes(r, d)
	h := r.Router.Handler
	return &server{
		httpServer: &fasthttp.Server{
			Handler: h,
		},
	}
}

func (s *server) Run() {
	log.Println("Starting server")
	if err := s.httpServer.ListenAndServe(":3030"); err != nil {
		log.Fatal(err)
	}
}

func (s *server) Shutdown() error {
	return s.httpServer.Shutdown()
}
