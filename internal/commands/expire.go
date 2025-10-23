package commands

import (
	"strconv"
	"time"

	"github.com/reche13/echodb/internal/protocol"
	"github.com/reche13/echodb/internal/store"
)

func init(){
	Register("EXPIRE", Expire)
}

func Expire(store *store.Store, args []*protocol.RESPValue) *protocol.RESPValue {
	if len(args) != 2 {
		return protocol.NewError("ERR wrong number of arguments for 'EXPIRE' command")
	}

	key, ok := args[0].GetString()
	if !ok {
		return protocol.NewError("ERR invalid key for 'EXPIRE' command")
	}

	secondsStr, ok := args[1].GetString()
	if !ok {
		return protocol.NewError("ERR invalid seconds argument for 'EXPIRE' command")
	}

	seconds, err := strconv.ParseInt(secondsStr, 10, 64)
	if err != nil {
		return protocol.NewError("ERR value is invalid or out of range")
	}

	expiresAt := time.Now().Unix() + seconds

	if store.Expire(key, expiresAt) {
		return protocol.NewInteger(1)
	}
	
	return protocol.NewInteger(0)
}