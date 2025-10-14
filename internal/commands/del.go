package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init() {
	Register("DEL", Del)
}

func Del(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) == 0 {
		return protocol.NewError("ERR wrong number of arguments for 'DEL' command")
	}

	keys := []string{}
	for _, a := range args {
		k, ok := a.GetString()
		if !ok {
			return protocol.NewError("ERR invalid key for 'DEL' command")
		}
		keys = append(keys, k)
	}

	deleted := store.Del(keys...)
	return protocol.NewInteger(int64(deleted))
}