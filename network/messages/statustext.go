package messages

import (
	"../../tools"
	"../../types"
)

func StatusText(str string) types.Message {
	var outMsg types.Message
	outMsg.Type = 0x09
	payload := tools.Int32ToBytes(1)
	payload = append(payload, tools.GetEncodedString(str) ...)
	outMsg.Payload = payload
	outMsg.Length = int16(len(outMsg.Payload) + 3)
	return outMsg
}