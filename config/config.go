package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OdooUsername string
	OdooPassword string
	OdooName     string
	OdooEndpoint string
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

	odooName := os.Getenv("ODOO_NAME")
	if odooName == "" {
		return nil, errors.New("ODOO_NAME is empty")
	}

	odooEndpoint := os.Getenv("ODOO_ENDPOINT")
	if odooEndpoint == "" {
		return nil, errors.New("ODOO_ENDPOINT is empty")
	}

	return &Config{
		OdooUsername: odooUsername,
		OdooPassword: odooPassword,
		OdooName:     odooName,
		OdooEndpoint: odooEndpoint,
	}, nil
}
