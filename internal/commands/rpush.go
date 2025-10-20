package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init(){
	Register("RPUSH", RPush)
}

func RPush(store *store.Store, args []*protocol.RESPValue)*protocol.RESPValue{
	if len(args) < 2 {
		return protocol.NewError("ERR wrong number of arguments for 'RPUSH' command")
	}

	key, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid key for 'RPUSH' command")
	}

	values := make([]string, 0, len(args) - 1)
	for _, arg := range args[1:] {
		val, ok := arg.GetString()
		if !ok {
			return protocol.NewError("ERR invalid list element for 'RPUSH' command")
		}
		values = append(values, val)
	}

	newLen, err := store.RPush(key, values...)
	if err != nil {
		return protocol.NewError(err.Error())
	}

	return protocol.NewInteger(int64(newLen))
}