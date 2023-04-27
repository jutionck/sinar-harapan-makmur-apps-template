package common

import (
	"errors"
	"github.com/joho/godotenv"
)

func LoadFileEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return errors.New("failed to load .env File")
	}
	return nil
}
