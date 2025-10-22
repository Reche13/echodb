package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/reche13/echodb/internal/commands"
	"github.com/reche13/echodb/internal/server"
	"github.com/reche13/echodb/internal/store"
)

func main() {
	st := store.New()
	aof, err := store.NewAOFManager("echodb.aof")
	if err != nil {
		log.Fatal(err)
	}
	defer func(){
		if err := aof.Close(); err != nil {
			log.Println("Failed to close AOF:", err)
		}
	}()
	
	log.Println("Restoring AOF data...")
	if err := aof.LoadFromAOF(st); err != nil {
		log.Println("Failed to restore AOF:", err)
	}

	st.Aof = aof
	ex := commands.NewExecutor(st)
	s := server.New(":6380", ex)

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