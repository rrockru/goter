package tools

import (
	"bytes"
	"encoding/binary"
	"log"
	
	"../types"
)

func GetMessageLen(buffer []byte) int {
	var len int16
	buf := bytes.NewBuffer(buffer)
	err := binary.Read(buf, binary.LittleEndian, &len)
	if err != nil {
		log.Fatal("Can't decode length: ", err.Error())
	}
	return int(len);	
}

func GetStringLen(buffer []byte) int {
	var len int8
	buf := bytes.NewBuffer(buffer)
	err := binary.Read(buf, binary.LittleEndian, &len)
	if err != nil {
		log.Fatal("Can't decode length: ", err.Error())
	}
	return int(len);	
}

func Int8ToBytes(x int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int8(x))
	if err != nil {
		log.Fatal("Can't encode length: ", err.Error())
	}
	return buf.Bytes()
}

func Int16ToBytes(x int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int16(x))
	if err != nil {
		log.Fatal("Can't encode length: ", err.Error())
	}
	return buf.Bytes()
}

func Int32ToBytes(x int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int32(x))
	if err != nil {
		log.Fatal("Can't encode length: ", err.Error())
	}
	return buf.Bytes()
}

func Float32ToBytes(x float32) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, float32(x))
	if err != nil {
		log.Fatal("Can't encode length: ", err.Error())
	}
	return buf.Bytes()
}

func GetEncodedString(str string) []byte {
	len := len(str)
	ret := append(Int8ToBytes(len))
	ret = append(ret, []byte(str)...)
	return ret
}

func EncodeMessage(msg types.Message) []byte {
	tmpBuf := new(bytes.Buffer)
	binary.Write(tmpBuf, binary.LittleEndian, &msg.Length)
	buf := append([]byte(tmpBuf.Bytes()))
	buf = append(buf, msg.Type)
	buf = append(buf, msg.Payload...)
	return buf
}

func DecodeMessage(buffer []byte) types.Message {
	var msg types.Message
	buf := bytes.NewBuffer(buffer[0:2])
	err := binary.Read(buf, binary.LittleEndian, &msg.Length)
	if err != nil {
		log.Fatal("Can't decode length: ", err.Error())
	}
	msg.Type = buffer[2]
	msg.Payload = make([]byte, msg.Length)
	msg.Payload = buffer[3:len(buffer)]
	return msg
}