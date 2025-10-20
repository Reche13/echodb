package commands

import (
	"strconv"

	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init(){
	Register("RPOP", RPop)
}

func RPop(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) < 1 || len(args) > 2 {
		return protocol.NewError("ERR wrong number of arguments for 'RPOP' command")
	}

	key, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid key for 'RPOP' command")
	}

	count := 1
	if len(args) > 1 {
		nStr, ok := args[1].GetString()
		if !ok {
			return protocol.NewError("ERR invalid count argument for 'RPOP' command")
		}

		n, err := strconv.Atoi(nStr)
		if err != nil || n < 1 {
			return protocol.NewError("ERR value is out of range, must be positive")
		}

		count = n
	}

	values := store.RPop(key, count)

	if len(values) == 0 {
		return protocol.NewArray(nil)
	}

	arr := make([]*protocol.RESPValue, len(values))
	for i, str := range values {
		arr[i] = protocol.NewBulkString(str)
	}

	return protocol.NewArray(arr)
}