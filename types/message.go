package types

type Message struct {
	Length 	int16
	Type	byte
	Payload []byte
}