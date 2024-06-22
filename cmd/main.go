package main

import (
	"fmt"
	"log"

	"github.com/giftalapp/authsrv/config"
	"github.com/giftalapp/authsrv/src"
	"github.com/giftalapp/authsrv/utilities/db"
	"github.com/giftalapp/authsrv/utilities/pub"
)

func main() {
	db, err := db.Start(config.Env.DBAddr)
	log.Printf("[INFO] Connecting to database at %s", config.Env.DBAddr)

	if err != nil {
		log.Fatalf("[ERROR] while starting database: %s", err)
	}

	log.Println("[INFO] Connecting to Pub Client")
	pubc, err := pub.NewPubClient()

	if err != nil {
		log.Fatalf("[ERROR] while Connecting to Pub Client: %s", err)
	}

	address := fmt.Sprintf("%s:%s", config.Env.APIHost, config.Env.APIPort)
	authsrv := src.NewAuthService(address, db, pubc)
	log.Printf("[INFO] starting server at address at %s", address)

	if err = authsrv.Run(); err != nil {
		log.Fatalf("[ERROR] while running server: %s", err)
	}
}
