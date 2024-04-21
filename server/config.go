package main

import (
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Server struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type Config struct {
	Server Server `toml:"server"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open(".config.toml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	configuration, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(configuration, &config)
	if err != nil {
		return nil, nil
	}
	return &config, nil
}
