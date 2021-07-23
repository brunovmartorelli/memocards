package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/brunovmartorelli/memo-api/config"
	"github.com/brunovmartorelli/memo-api/repository"
	"github.com/brunovmartorelli/memo-api/server"
	"github.com/brunovmartorelli/memo-api/server/router"
)

func main() {
	ctx := context.Background()
	cfg, _ := config.New()

	//FIXME: Mensagem de erro caso não conseguir se conectar com o banco de dados
	m := repository.NewMongo(&cfg.Mongo)
	if err := m.Client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer m.Disconnect(ctx)

	r := router.New(m)
	httpServer := server.New(&cfg.Server, r.Router.Handler)
	go httpServer.Run()

	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)
	<-shutdown
	if err := httpServer.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
