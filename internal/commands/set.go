package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init(){
	Register("SET", Set)
}

func Set(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) < 2 {
		return protocol.NewError("ERR wrong number of arguments for 'SET' command")
	}

	key, ok1 := args[0].GetString()
	val, ok2 := args[1].GetString()
	if !ok1 || !ok2 {
		return protocol.NewError("ERR invalid arguments for 'SET' command")
	}

	store.Set(key, val)
	return protocol.NewSimpleString("OK")
}