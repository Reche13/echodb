package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/reche13/echodb/internal/server"
	"github.com/reche13/echodb/internal/store"
)

func main() {
	store := store.New()

	s := server.New(":6380", store)

	go func(){
		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh
	s.Stop()
}