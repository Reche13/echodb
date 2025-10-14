package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init() {
	Register("GET", Get)
}

func Get(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) != 1 {
		return protocol.NewError("ERR wrong number of arguments for 'GET' command")
	}

	key, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid key for 'GET' command")
	}

	val, exists := store.Get(key)
	if !exists {
		return protocol.NewNullBulkString()
	}

	return protocol.NewBulkString(val)
}