package main

import (
	"log"

	"github.com/reche13/echodb/internal/server"
)

func main() {
	s := server.New(":6380")

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}