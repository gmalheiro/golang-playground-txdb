package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	config cfg
)

type cfg struct {
	Port       string
	DbUser     string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
}

func Load() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("an error occurred while trying to parse config file: %w", err)
	}

	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "8080"
	}

	config.DbUser = os.Getenv("DB_USER")
	if config.DbUser == "" {
		config.DbUser = "root"
	}

	config.DbPassword = os.Getenv("DB_PASSWORD")
	if config.DbPassword == "" {
		config.DbPassword = "root"
	}

	config.DbHost = os.Getenv("DB_HOST")
	if config.DbHost == "" {
		config.DbHost = "localhost"
	}

	config.DbPort = os.Getenv("DB_PORT")
	if config.DbPort == "" {
		config.DbPort = "3306"
	}

	config.DbName = os.Getenv("DB_NAME")
	if config.DbName == "" {
		config.DbName = "meli"
	}

	return nil
}

func Get() cfg {
	return config
}
