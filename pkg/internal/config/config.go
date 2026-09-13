package config

import (
	"fmt"
	"os"
)

type Config struct {
	Token string
	URL   string
}

func Load() (*Config, error) {
	token := os.Getenv("OPENGIST_CLI_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("OPENGIST_CLI_TOKEN is not set")
	}

	url := os.Getenv("OPENGIST_CLI_URL")
	if url == "" {
		return nil, fmt.Errorf("OPENGIST_CLI_URL is not set")
	}

	return &Config{
		Token: token,
		URL:   url,
	}, nil
}
