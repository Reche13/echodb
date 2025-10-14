package commands

import (
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init() {
	Register("PING", Ping)
}

func Ping(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) == 0 {
		return protocol.NewSimpleString("PONG")
	}

	str, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid argument for PING")
	}

	return protocol.NewSimpleString(str)
}