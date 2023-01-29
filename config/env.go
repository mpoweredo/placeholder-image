package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadPORT() (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}

	port := os.Getenv("PORT")

	return port, nil
}
