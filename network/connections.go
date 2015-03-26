package network

import (
	"container/list"
	"log"
	"net"

	"../player"
	"../tools"
	"../types"
)

func handleMessages(buffer []byte, connection *types.Connection) {
	pos := 0
	for pos < len(buffer) {
		len := tools.GetMessageLen(buffer[pos:pos + 2])
		msg := tools.DecodeMessage(buffer[pos:pos + len])
		log.Printf("%s > Length: %d; Type: 0x%x; Payload: 0x%x.",
			connection.IP, msg.Length, msg.Type, msg.Payload)
		connection.Incoming<- msg
		pos += len
	}
	
}

func handleRead(connection *types.Connection) {	
	buffer := make([]byte, 1024)
	for {
		n, err := connection.Conn.Read(buffer)
		if err != nil {
			// если у нас ошибка чтения, значит соединение разорвано
			connection.Close()		
			break
		}
		handleMessages(buffer[0:n], connection)
	}
}

func handleWrite(connection *types.Connection) {
	for {
		select {
		case msg := <-connection.Outgoing:
			log.Printf("%s < Length: %d; Type: 0x%x; Payload: 0x%x.",
				connection.IP, msg.Length, msg.Type, msg.Payload)
			buf := tools.EncodeMessage(msg)
			connection.Conn.Write(buf)
			if msg.Type == 0x02 {
				connection.Conn.Close()
			}
		case <-connection.Quit:
			log.Printf("%s disconnected...", connection.IP)
			connection.Conn.Close()			
		}
	}	
}

func handleConnection(conn net.Conn, connectionList *list.List, world *types.World) {
	ip := conn.RemoteAddr().(*net.TCPAddr).IP.String();
	log.Printf("%s connected...", ip)
	connection := &types.Connection{conn, ip, make(chan types.Message), make(chan types.Message), 
		make(chan bool), connectionList, world, player.NewPlayer()}
	go handleRead(connection)
	go handleWrite(connection)
	go Process(connection)
	connectionList.PushBack(connection)
}