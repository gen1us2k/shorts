package config

import "github.com/kelseyhightower/envconfig"

// ShortsConfig stores configuration for the application
type ShortsConfig struct {
	BindAddr  string `envconfig:"BIND_ADDR" default:":8090"`
	DSN       string `envconfig:"DSN"`
	URLLength int    `envconfig:"URL_LENGTH" default:"8"`
}

// Parse parses environment variables and returns
// filled struct
func Parse() (*ShortsConfig, error) {
	var c ShortsConfig
	err := envconfig.Process("", &c)
	return &c, err
}
