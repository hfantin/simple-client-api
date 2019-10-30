package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	ServerPort string
	DBHost     string
	DBUser     string
	DBPass     string
	DBName     string
}

func Config() Environment {
	env := Environment{"5000", "localhost", "guest", "guest", "test"}
	// Load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using default values....")
	}
	env.ServerPort = os.Getenv("SERVER_PORT")
	env.DBHost = os.Getenv("DB_HOST")
	env.DBUser = os.Getenv("DB_USER")
	env.DBPass = os.Getenv("DB_PASS")
	env.DBName = os.Getenv("DB_NAME")
	return env
}
