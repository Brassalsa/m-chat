package main

import (
	"context"
	"log"
	"os"

	"github.com/Brassalsa/m-chat/internal/api"
	"github.com/Brassalsa/m-chat/internal/db"
	"github.com/Brassalsa/m-chat/pkg"
	"github.com/joho/godotenv"
)

type Page struct {
	Payload string
}

func main() {
	godotenv.Load()

	dbUrl := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	if dbUrl == "" || dbName == "" {
		log.Fatal("DB_URI or DB_NAME not found in .env")
	}

	s := pkg.NewServer(":3000", "web/static")

	// connect to db
	ctx := context.Background()
	dbC := db.NewMongoDb(dbUrl, dbName)
	dbC.Connect(ctx)

	s.RegisterHandler("/", api.HandleV1Route(ctx, dbC))
	log.Printf("serving @ http://localhost:%s\n", s.ListenAddr)
	log.Fatal(s.Listen())
}
