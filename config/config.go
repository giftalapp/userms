package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIHost    string
	APIPort    string
	DBUser     string
	DBPassword string
	DBAddr     string
	DBName     string
}

var Env = initConfig()

func initConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
	}

	return &Config{
		APIHost:    getEnv("API_HOST", "127.0.0.1"),
		APIPort:    getEnv("API_PORT", "8080"),
		DBUser:     getEnv("DB_USER", "server"),
		DBPassword: getEnv("DB_PASSWORD", "secret"),
		DBAddr:     getEnv("DB_ADDR", "./test.db"),
		DBName:     getEnv("DB_NAME", "giftal"),
	}
}

func getEnv(name string, fallback string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}

	return fallback
}
