package server

import (
	"log"

	"github.com/brunovmartorelli/memo-api/config"
	"github.com/brunovmartorelli/memo-api/server/router"
	"github.com/valyala/fasthttp"
)

type server interface {
	Run()
	Shutdown() error
}

type fastserver struct {
	httpServer *fasthttp.Server
	cfg        *config.Server
}

func New(cfg *config.Server) server {
	r := router.New()
	r.Routes()
	h := r.Router.Handler
	hs := &fasthttp.Server{
		Handler: h,
	}
	return &fastserver{
		httpServer: hs,
		cfg:        cfg,
	}
}

func (s *fastserver) Run() {
	log.Println("Starting server")
	if err := s.httpServer.ListenAndServe(s.cfg.Port); err != nil {
		log.Fatal(err)
	}
}

func (s *fastserver) Shutdown() error {
	return s.httpServer.Shutdown()
}
