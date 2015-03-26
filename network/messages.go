package network

import (
	"bytes"
	
	"./messages"
	
	"../config"	
	"../tools"
	"../types"
)

func Process(connection *types.Connection) {
	for {
		msg := <-connection.Incoming
		switch {
		case msg.Type == 0x01:
			if config.Password == "" {
				connection.Outgoing<- messages.ConnectionApproved(connection.ConnectionList.Len() - 1)
			} else {
				connection.Outgoing<- messages.RequestPassword()
			}			
		case msg.Type == 0x06:
			connection.Outgoing<- messages.WorldInfo(connection.World)
		case msg.Type == 0x08:
			connection.Outgoing<- messages.StatusText("Receiving tile data");
			connection.Outgoing<- messages.Spawn()
		case msg.Type == 0x26:
			len := tools.GetStringLen(msg.Payload[0:1])
			if bytes.Equal(msg.Payload[1:len + 1], []byte(config.Password)) {
				connection.Outgoing<- messages.ConnectionApproved(connection.ConnectionList.Len() - 1)
			} else {
				connection.Outgoing<- messages.Disconnect("Wrong password!")
			}
		}
	}
}