package messages

import (
	"../../types"
)

func RequestPassword() types.Message {
	var outMsg types.Message
	outMsg.Type = 0x25
	outMsg.Length = int16(len(outMsg.Payload) + 3)
	return outMsg
}