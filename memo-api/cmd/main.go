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
)

func main() {
	ctx := context.Background()
	cfg, _ := config.New()

	m := repository.NewMongo(&cfg.Mongo)
	m.Client.Connect(ctx)
	defer m.Disconnect(ctx)

	httpServer := server.New(&cfg.Server)
	go httpServer.Run()

	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)
	<-shutdown
	if err := httpServer.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
