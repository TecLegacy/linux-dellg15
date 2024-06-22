package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBName     string
	DBUser     string
	DBPassword string
	DBAddress  string

	JWTExpiration int64
	JWTSecret     string
}

// Global singleton to hold initConfig
var Envs = InitConfig()

func InitConfig() Config {
	//lOAD ENV
	godotenv.Load()

	return Config{
		PublicHost: getenv("PUBLIC_HOST", "http://localhost"),
		Port:       getenv("PORT", "8080"),
		DBUser:     getenv("DB_USER", "root"),
		DBName:     getenv("DB_NAME", "ecom"),
		DBPassword: getenv("DB_PASS", "mypass"),
		DBAddress:  fmt.Sprintf("%s:%s", getenv("DB_HOST", "127.0.0.1"), getenv("DB_PORT", "3306")),

		JWTExpiration: getenvInt("JWT_EXPIRATION", 3600*24*7), // Day validation

		JWTSecret: getenv("JWT_SECRET", "ecom"),
	}
}

// * UTIL to lookup ENV
func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getenvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64) // int64
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
