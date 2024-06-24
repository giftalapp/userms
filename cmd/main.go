package main

import (
	"fmt"
	"log"

	"github.com/giftalapp/authsrv/config"
	"github.com/giftalapp/authsrv/src"
	"github.com/giftalapp/authsrv/utilities/db"
	"github.com/giftalapp/authsrv/utilities/fb"
	"github.com/giftalapp/authsrv/utilities/pub"
)

func main() {
	db, err := db.Start(config.Env.DBAddr)
	log.Printf("[INFO] Connecting to database at %s\n", config.Env.DBAddr)

	if err != nil {
		log.Fatalf("[ERROR] while connecting to database: %s\n", err)
	}

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
	authsrv := src.NewAuthService(address, db, fb, pubc)
	log.Printf("[INFO] starting server at address at %s\n", address)

	if err = authsrv.Run(); err != nil {
		log.Fatalf("[ERROR] while running server: %s\n", err)
	}
}
