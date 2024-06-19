package main

import (
	"fmt"
	"log"

	"github.com/giftalapp/authsrv/config"
	"github.com/giftalapp/authsrv/src"
)

func main() {
	db, err := src.StartDB(config.Env.DBAddr)

	if err != nil {
		log.Fatalf("[ERROR] while starting database: %s", err)
	}

	address := fmt.Sprintf("%s:%s", config.Env.APIHost, config.Env.APIPort)
	authsrv := src.NewAuthService(address, db)

	log.Printf("[INFO] starting server at address: %s", address)

	if err = authsrv.Run(); err != nil {
		log.Fatalf("[ERROR] while running server: %s", err)
	}
}
