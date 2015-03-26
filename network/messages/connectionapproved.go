package messages

import (
	"../../types"
)

func ConnectionApproved(playerSlot int) types.Message {
	var outMsg types.Message
	outMsg.Type = 0x03
	outMsg.Payload = make([]byte, 1)
	outMsg.Payload = []byte{byte(playerSlot)}
	outMsg.Length = int16(len(outMsg.Payload) + 3)
	return outMsg
}