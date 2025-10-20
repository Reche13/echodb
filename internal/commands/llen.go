package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init(){
	Register("LLEN", LLen)
}

func LLen(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) != 1 {
		return protocol.NewError("ERR wrong number of arguments for 'LLEN' command")
	}

	key, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid key for 'LLEN' command")
	}

	length := store.LLen(key)

	if length == -1 {
		return protocol.NewError("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	return protocol.NewInteger(int64(length))
}