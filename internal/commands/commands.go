package commands

import (
	"strings"

	"github.com/reche13/echodb/internal/protocol"
)

type CommandFunc func(args []*protocol.RESPValue) *protocol.RESPValue

var registry = map[string]CommandFunc{}

func Register(name string, fn CommandFunc) {
	registry[strings.ToUpper(name)] = fn
}

func Execute(command *protocol.RESPValue) *protocol.RESPValue {
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

	return fn(args[1:])
}