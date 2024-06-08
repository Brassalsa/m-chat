package main

import (
	"log"
)

type Page struct {
	Payload string
}

func main() {
	s := newServer(":3000", "static")
	s.RegisterHandler("/", handleV1Route())
	log.Printf("serving @ http://localhost:%s\n", s.ListenAddr)
	log.Fatal(s.Listen())
}
