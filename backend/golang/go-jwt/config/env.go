package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBName     string
	DBUser     string
	DBPassword string
	DBAddress  string
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
	}
}

// * UTIL to lookup ENV
func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
