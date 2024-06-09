package main

import (
	"log"

	"github.com/Brassalsa/m-chat/internal/api"
	"github.com/Brassalsa/m-chat/pkg"
)

type Page struct {
	Payload string
}

func main() {
	s := pkg.NewServer(":3000", "web/static")
	s.RegisterHandler("/", api.HandleV1Route())
	log.Printf("serving @ http://localhost:%s\n", s.ListenAddr)
	log.Fatal(s.Listen())
}
