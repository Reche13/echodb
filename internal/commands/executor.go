package commands

import (
	"github.com/reche13/echodb/internal/persistence"
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

type Executor struct {
	store *store.Store
	persister persistence.Persistence
}

func NewExecutor(store *store.Store, executor persistence.Persistence) *Executor {
	return &Executor{store: store, persister: executor}
}

func (e *Executor) Execute(command *protocol.RESPValue) *protocol.RESPValue {
	result := Execute(e.store, command)

	if e.persister != nil {
		e.persister.Log(command)
	}
	
	return result
}