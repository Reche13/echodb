package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init(){
	Register("EXISTS", Exists)
}

func Exists(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) == 0 {
		return protocol.NewError("ERR wrong number of arguments for 'EXISTS' command")
	}

	keys := []string{}
	for _, a := range args {
		k, ok := a.GetString()
		if !ok {
			return protocol.NewError("ERR invalid key for 'EXISTS' command")
		}
		keys = append(keys, k)
	}

	count := store.Exists(keys...)
	return protocol.NewInteger(int64(count))
}