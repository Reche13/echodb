package commands

import (
	"strings"

	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

type CommandFunc func(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue

var registry = map[string]CommandFunc{}

func Register(name string, fn CommandFunc) {
	registry[strings.ToUpper(name)] = fn
}

func Execute(store *store.Store, command *protocol.RESPValue) *protocol.RESPValue {
	args, ok := command.GetArray()
	if !ok || len(args) == 0 {
		return protocol.NewError("ERR empty command")
	}

	cmdName, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid command name")
	}

	fn, exists := registry[strings.ToUpper(cmdName)]
	if !exists {
		return protocol.NewError("ERR unknown command '" + cmdName + "'")
	}

	return fn(store, args[1:])
}