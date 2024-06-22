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
	RedisURL   string
}

var Env = initConfig()

func initConfig() *Config {
	err := godotenv.Load(".config")

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
		RedisURL:   getEnv("REDIS_URL", "redis://:@localhost:6379/0"),
	}
}

func getEnv(name string, fallback string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}

	return fallback
}
