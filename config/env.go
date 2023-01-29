package config

import (
	"os"
)

func LoadPORT() (string, error) {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	return port, nil
}
