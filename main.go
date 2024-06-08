package main

import (
	"log"
)

type Page struct {
	Payload string
}

func main() {
	s := newServer(":4000", "static")
	s.RegisterHandler("/", handleV1Route())
	log.Printf("serving @ http://loclahost:%s\n", s.ListenAddr)
	log.Fatal(s.Listen())
}
