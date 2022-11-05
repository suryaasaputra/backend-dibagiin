package config

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

var (
	SERVER_PORT, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	SECRET_KEY     = os.Getenv("SECRET_KEY")

	HOST        = os.Getenv("DB_HOST")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_PORT, _  = strconv.Atoi(os.Getenv("DB_PORT"))
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_CONFIG   = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)

	STORAGE_PATH = os.Getenv("STORAGE_PATH")
	BUCKET_NAME  = os.Getenv("BUCKET_NAME")
)
