package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init(){
	Register("PERSIST", Persist)
}

func Persist(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) != 1 {
		return protocol.NewError("ERR wrong number of arguments for 'PERSIST' command")
	}

	key, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid key for 'PERSIST' command")
	}
	
	if store.Persist(key) {
		return protocol.NewInteger(1)
	}
	
	return protocol.NewInteger(0)
}