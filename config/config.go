package config

import "os"

var (
	PORT       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
)

// Initialize reads env vars into package vars
func Initialize() {
	DBHost = os.Getenv("DB_HOST")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBPort = os.Getenv("DB_PORT")
	DBSSLMode = os.Getenv("DB_SSL_MODE")
	PORT = os.Getenv("APP_SERVER_PORT")
}
