package messages

import (
	"../../tools"
	"../../types"
)

func Disconnect(str string) types.Message {
	var outMsg types.Message
	outMsg.Type = 0x02
	payload := tools.GetEncodedString(str)
	outMsg.Payload = make([]byte, len(payload))
	outMsg.Payload = payload
	outMsg.Length = int16(len(outMsg.Payload) + 3)
	return outMsg
}