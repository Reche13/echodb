package commands

import (
	"strconv"
	"strings"
	"time"

	"github.com/reche13/echodb/internal/persistence"
	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

type Executor struct {
	store *store.Store
	persister persistence.Persistence
}

func NewExecutor(store *store.Store, executor persistence.Persistence) *Executor {
	return &Executor{store: store, persister: executor}
}

func (e *Executor) Execute(command *protocol.RESPValue) *protocol.RESPValue {
	result := Execute(e.store, command)

	if e.persister != nil && e.shouldLogCommand(command) {
		logCmd := command
		cmdArr, ok := command.GetArray()
		if ok && len(cmdArr) > 0 {
			cmdName, ok := cmdArr[0].GetString()
			if ok && strings.ToUpper(cmdName) == "EXPIRE" && len(cmdArr) == 3 {
				ttlStr, ok := cmdArr[2].GetString()
				if ok {
					seconds, err := strconv.ParseInt(ttlStr, 10, 64)
					if err == nil {
						absoluteTime := time.Now().UnixMilli() + (seconds * 1000)

						logCmd = protocol.NewArray([]*protocol.RESPValue{
							cmdArr[0],
							cmdArr[1],
							protocol.NewInteger(absoluteTime),
						})
					}
				}
			}
		}



		e.persister.Log(logCmd)
	}

	return result
}

func (e *Executor) shouldLogCommand(command *protocol.RESPValue) bool {
	args, ok := command.GetArray()
	if !ok || len(args) == 0 {
		return false
	}

	cmdName, ok := args[0].GetString()
	if !ok {
		return false
	}

	writeCommands := map[string]bool{
		"SET": true,
		"DEL": true,
		"LPUSH": true,
		"RPUSH": true,
		"LPOP": true,
		"RPOP": true,
		"EXPIRE": true,
		"PERSIST": true,
	}

	return writeCommands[strings.ToUpper(cmdName)]
}