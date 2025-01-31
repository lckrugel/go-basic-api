package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig      AppConfig
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

type AppConfig struct {
	Host string
	Port string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	db_host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		return nil, errors.New("DB_HOST is required")
	}

	db_port, exists := os.LookupEnv("DB_PORT")
	if !exists {
		return nil, errors.New("DB_PORT is required")
	}

	db_user, exists := os.LookupEnv("DB_USER")
	if !exists {
		return nil, errors.New("DB_USER is required")
	}

	db_password, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		return nil, errors.New("DB_PASSWORD is required")
	}

	app_host, exists := os.LookupEnv("APP_HOST")
	if !exists {
		return nil, errors.New("APP_HOST is required")
	}

	app_port, exists := os.LookupEnv("APP_PORT")
	if !exists {
		return nil, errors.New("APP_PORT is required")
	}

	dbConfig := DatabaseConfig{
		Host:     db_host,
		Port:     db_port,
		User:     db_user,
		Password: db_password,
	}

	app_config := AppConfig{
		Host: app_host,
		Port: app_port,
	}

	return &Config{
		AppConfig:      app_config,
		DatabaseConfig: dbConfig,
	}, nil
}
