package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init() {
	Register("ECHO", Echo)
}

func Echo(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) != 1 {
		return protocol.NewError("ERR wrong number of arguments for 'ECHO' command")
	}
	
	str, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid argument for 'ECHO' command")
	}

	return protocol.NewSimpleString(str)
}