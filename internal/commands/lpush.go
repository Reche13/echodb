package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init(){
	Register("LPUSH", LPush)
}

func LPush(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) < 2 {
		return protocol.NewError("ERR wrong number of arguments for 'LPUSH' command")
	}

	key, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid key for 'LPUSH' command")
	}

	values := make([]string, 0, len(args) - 1)
	for _, arg := range args[1:] {
		val, ok := arg.GetString()
		if !ok {
			return protocol.NewError("ERR invalid list element for 'LPUSH' command")
		}
		values = append(values, val)
	}

	newLen := store.LPush(key, values...)

	if newLen == -1 {
		return protocol.NewError("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	return protocol.NewInteger(int64(newLen))
}