package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/reche13/echodb/internal/commands"
	"github.com/reche13/echodb/internal/config"
	"github.com/reche13/echodb/internal/persistence"
	"github.com/reche13/echodb/internal/server"
	"github.com/reche13/echodb/internal/store"
)

func Execute() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load configurations: %v\n", err)
		os.Exit(1)
	}
	
	st := store.New()

	var aof *persistence.AOFManager
	if cfg.Persistence.Enabled {
				log.Println("YESS")
		aof, err = persistence.NewAOFManager(&cfg.Persistence)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err := aof.Close(); err != nil {
				log.Println("Failed to close AOF:", err)
			}
		}()

		aof.StartBackgroundFlush()
		if err := aof.Load(st); err != nil {
			log.Println("Failed to restore AOF:", err)
		}
	} else {
		log.Println("Persistence is disabled. Skipping AOF initialization.")
	}

	ex := commands.NewExecutor(st, aof)
	s := server.New(&cfg.Server, ex)

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