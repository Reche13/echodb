package main

import (
	"log"

	"github.com/reche13/echodb/internal/server"
	"github.com/reche13/echodb/internal/store"
)

func main() {
	store := store.New()

	s := server.New(":6380", store)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}