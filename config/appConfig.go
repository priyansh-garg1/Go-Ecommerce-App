package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
}

func SetupEnv() (cfg AppConfig, err error) {

	godotenv.Load()

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	Dsn := os.Getenv("DSN")
	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: Dsn}, nil
}
