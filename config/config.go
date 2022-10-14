package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OdooUsername string
	OdooPassword string
}

func NewConfigFromEnv() (*Config, error) {
	godotenv.Load()

	odooUsername := os.Getenv("ODOO_USERNAME")
	if odooUsername == "" {
		return nil, errors.New("ODOO_USERNAME is empty")
	}

	odooPassword := os.Getenv("ODOO_PASSWORD")
	if odooPassword == "" {
		return nil, errors.New("ODOO_PASSWORD is empty")
	}

	return &Config{
		OdooUsername: odooUsername,
		OdooPassword: odooPassword,
	}, nil
}
