package service

import (
	"errors"
	"strings"

	constants "github.com/sureshsolanki17/ant-golang/const"
)

// Config contains the configuration for initializing the AntApp client
type Config struct {
	Exchange string // Name of the exchange (e.g., NSE, BSE, etc.)
	BaseURL  string // Base URL for the API
}

// Validate checks the configuration for missing or invalid values
func (c *Config) Validate() error {
	if strings.TrimSpace(c.Exchange) == "" {
		return errors.New("exchange cannot be empty")
	}
	if strings.TrimSpace(c.BaseURL) == "" {
		return errors.New("baseURL cannot be empty")
	}
	return nil
}

// DefaultConfig provides a default configuration for initializing the AntApp
func DefaultConfig() *Config {
	return &Config{
		Exchange: "NSE",
		BaseURL:  constants.BaseURL,
	}
}
