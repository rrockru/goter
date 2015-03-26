package main

import (
	"./network"
)

func main() {
	s := new(network.Server)
	s.Start()
}