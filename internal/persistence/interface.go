package persistence

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

type Persistence interface {
	Log(command *protocol.RESPValue) error
	Load(store *store.Store) error
	Close() error
}