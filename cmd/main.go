package main

import (
	"context"
	"fmt"
	"log"

	"github.com/giftalapp/userms/config"
	"github.com/giftalapp/userms/src"
	"github.com/giftalapp/userms/utilities/fb"
	"github.com/giftalapp/userms/utilities/pub"
	"github.com/jackc/pgx/v5"
)

func main() {
	db, err := pgx.Connect(context.Background(), config.Env.DBAddr)
	log.Printf("[INFO] Connecting to database")

	if err != nil {
		log.Fatalf("[ERROR] while connecting to database: %s\n", err)
	}
	defer db.Close(context.Background())

	fb, err := fb.Start()
	log.Println("[INFO] Connecting to firebase")

	if err != nil {
		log.Fatalf("[ERROR] while connecting to firebase: %s\n", err)
	}

	log.Println("[INFO] Connecting to Pub Client")
	pubc, err := pub.NewPubClient(config.Env.RedisURL)

	if err != nil {
		log.Fatalf("[ERROR] while Connecting to Pub Client: %s\n", err)
	}

	address := fmt.Sprintf("%s:%s", config.Env.APIHost, config.Env.APIPort)
	userms := src.NewUserService(address, db, fb, pubc)
	log.Printf("[INFO] starting server at address at %s\n", address)

	if err = userms.Run(); err != nil {
		log.Fatalf("[ERROR] while running server: %s\n", err)
	}
}
