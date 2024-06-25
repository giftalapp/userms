package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName                  string
	APIHost                  string
	APIPort                  string
	DBAddr                   string
	RedisURL                 string
	OTPExpire                time.Duration
	OTPRefresh               time.Duration
	JWTSecret                string
	RateLimitInversedSeconds float64
	RateLimitRequests        int
}

var Env = initConfig()

func initConfig() *Config {
	err := godotenv.Load(".config")

	if err != nil {
		log.Fatalln(err)
	}

	return &Config{
		AppName:                  "Giftal [User]",
		APIHost:                  getEnv("API_HOST", "127.0.0.1"),
		APIPort:                  getEnv("API_PORT", "8080"),
		DBAddr:                   getEnv("DB_ADDR", "./test.db"),
		RedisURL:                 getEnv("REDIS_URL", "redis://:@localhost:6379/0"),
		OTPExpire:                time.Minute * 6,
		OTPRefresh:               time.Minute * 2,
		JWTSecret:                getEnv("JWT_SECRET", "secret"),
		RateLimitInversedSeconds: 1,
		RateLimitRequests:        5,
	}
}

func getEnv(name string, fallback string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}

	return fallback
}
