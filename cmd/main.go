package main

import (
	"log"

	"github.com/giftalapp/authsrv/config"
	"github.com/giftalapp/authsrv/src"
)

func main() {
	_, err := src.StartDB(config.Env.DBAddr)
	log.Printf("[INFO] Connecting to database at %s", config.Env.DBAddr)

	if err != nil {
		log.Fatalf("[ERROR] while starting database: %s", err)
	}

	log.Println("[INFO] Connecting to Amazon SNS")
	_, err = src.InitSNS()

	if err != nil {
		log.Fatalf("[ERROR] while while Connecting to Amazon SNS: %s", err)
	}
	//address := fmt.Sprintf("%s:%s", config.Env.APIHost, config.Env.APIPort)
	//authsrv := src.NewAuthService(address, db, fb)
	// log.Printf("[INFO] starting server at address at %s", address)

	// if err = authsrv.Run(); err != nil {
	// 	log.Fatalf("[ERROR] while running server: %s", err)
	// }
}
