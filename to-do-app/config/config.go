package config

import (
	"log"
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host string `yaml:"host"`
		Port int `yaml:"port"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"dbname"`
	} `yaml:"database"`
	Redis struct {
		Host string `yaml:"host"`
		Port int `yaml:"port"`
	} `yaml:"redis"`
}

func loadConfig() Config {
	file, err := os.ReadFile("config/config.yaml")

	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	return config
}