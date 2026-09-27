package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func MustLoad() *Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is not configured in env")
	}
	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is not configured in env")
	}

	cfg := Config{
		Port: port,
		Env:  env,
	}
	return &cfg
}
