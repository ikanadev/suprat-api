package main

import (
	"fmt"
	"os"
)

// Config here goes the configuration variables for this app
type Config struct {
	DB struct {
		User     string
		DBName   string
		Password string
		Host     string
		Port     string
	}
}

// GetConfig reads the env variables and return them
func GetConfig() (Config, error) {
	var conf Config
	envKeys := []string{
		"DB_USER",
		"DB_NAME",
		"DB_PASSWORD",
		"DB_HOST",
		"DB_PORT",
	}
	for _, key := range envKeys {
		if err := checkEnv(key); err != nil {
			return conf, err
		}
	}
	conf.DB.User = os.Getenv("DB_USER")
	conf.DB.DBName = os.Getenv("DB_NAME")
	conf.DB.Password = os.Getenv("DB_PASSWORD")
	conf.DB.Host = os.Getenv("DB_HOST")
	conf.DB.Port = os.Getenv("DB_PORT")
	return conf, nil
}

func checkEnv(key string) error {
	value := os.Getenv(key)
	if value == "" {
		return fmt.Errorf("env variable %s is empty, maybe is not providec", key)
	}
	return nil
}
