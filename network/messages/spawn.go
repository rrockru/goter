package messages

import (
	"../../types"
)

func Spawn() types.Message {
	var outMsg types.Message
	outMsg.Type = 0x31
	outMsg.Length = int16(len(outMsg.Payload) + 3)
	return outMsg
}