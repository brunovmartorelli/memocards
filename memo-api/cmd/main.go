package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/brunovmartorelli/memo-api/server"
)

func main() {
	httpServer := server.New()
	go httpServer.Run()

	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)
	<-shutdown
	if err := httpServer.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
