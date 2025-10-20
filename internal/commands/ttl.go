package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init(){
	Register("TTL", TTL)
}

func TTL(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) != 1 {
		return protocol.NewError("ERR wrong number of arguments for 'TTL' command")
	}

	key, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid key for 'TTL' command")
	}
	
	ttl := store.TTL(key)
	return protocol.NewInteger(ttl)
}