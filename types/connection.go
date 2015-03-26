package types

import (
	"container/list"
	"net"
)

type Connection struct {	
	Conn net.Conn
	IP string
	Incoming chan Message
	Outgoing chan Message
	Quit chan bool
	ConnectionList *list.List
	World *World
	Player *Player
}

func (c *Connection) RemoveMe() {
	for e := c.ConnectionList.Front(); e != nil; e = e.Next() {
		connection := e.Value.(*Connection)
		if c.Equal(connection) {
			c.ConnectionList.Remove(e)
		}
	}
}

func (c *Connection) Equal(other *Connection) bool{
	if c.Conn == other.Conn {
		return true
	}
	return false
}

func (c *Connection) Close() {
	c.Quit<- true
	c.Conn.Close()
	c.RemoveMe()
}