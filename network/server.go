package network

import (
	"container/list"
	"fmt"
	"log"
	"net"
	
	"../config"
	"../types"
	"../world"
)

type Server struct {
	World *types.World
}

func (s *Server) Start() {
	playerList := list.New()
	s.World = world.NewWorld()
	
	addr := fmt.Sprintf(":%d", config.Port)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Server started...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection error: ", err.Error())
		}
		go handleConnection(conn, playerList, s.World)
	}
}