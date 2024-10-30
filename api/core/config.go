package core

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB struct {
		Username string `envconfig:"DB_USERNAME" required:"true"`
		Host     string `envconfig:"DB_HOST" required:"true"`
		Password string `envconfig:"DB_PASSWORD" required:"true"`
		Name     string `envconfig:"DB_NAME" required:"true"`
		Port     string `envconfig:"DB_PORT" required:"true"`
	}
	Environment string `envconfig:"ENVIRONMENT" required:"true"` // prod or dev
}

func LoadConfig() Config {
	var config Config

	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to read .env file")
	}

	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err)
	}

	return config
}
