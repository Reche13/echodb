package commands

import (
	"strconv"

	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init(){
	Register("LRANGE", LRange)
}

func LRange(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) != 3 {
		return protocol.NewError("ERR wrong number of arguments for 'LRANGE' command")
	}

	key, ok1 := args[0].GetString()
	startStr, ok2 := args[1].GetString()
	stopStr, ok3 := args[2].GetString()
	if !ok1 || !ok2 || !ok3 {
		return protocol.NewError("ERR invalid arguments for 'LRANGE' command")
	}

	start, err1 := strconv.Atoi(startStr)
	stop, err2 := strconv.Atoi(stopStr)
	if err1 != nil || err2 != nil {
		return protocol.NewError("ERR value is not an integer or out of range")
	}

	values := store.LRange(key, start, stop)
	if values == nil {
		return protocol.NewArray(nil)
	}

	arr := make([]*protocol.RESPValue, len(values))
	for i, str := range values {
		arr[i] = protocol.NewBulkString(str)
	}

	return protocol.NewArray(arr)
}