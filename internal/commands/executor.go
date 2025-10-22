package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

type Executor struct {
	store *store.Store
}

func NewExecutor(store *store.Store) *Executor {
	return &Executor{store: store}
}

func (e *Executor) Execute(command *protocol.RESPValue) *protocol.RESPValue {
	return Execute(e.store, command)
}