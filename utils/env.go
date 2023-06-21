package utils

import (
	"os"
)

func GetValueFromEnv(key string) string {

	// load .env file
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	return os.Getenv(key)
}
