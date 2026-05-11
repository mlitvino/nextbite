package config

import (
	"errors"
	"os"
)

type Config struct {
	Port string
}

func Load() (Config, error) {
	port, err := requireEnv("BACKEND_PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{Port: port}, nil
}

func requireEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New(key + " is required")
	}
	return value, nil
}
