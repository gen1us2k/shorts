package config

import "github.com/kelseyhightower/envconfig"

const (
	KratosSessionKey = "ory_kratos_session"
	OwnerKey         = "owner_id"
)

// ShortsConfig stores configuration for the application
type ShortsConfig struct {
	BindAddr  string `envconfig:"BIND_ADDR" default:":8090"`
	DSN       string `envconfig:"DSN"`
	URLLength int    `envconfig:"URL_LENGTH" default:"8"`
	APIURL    string `envconfig:"API_URL"`
	UIURL     string `envconfig:"UI_URL"`
}

// Parse parses environment variables and returns
// filled struct
func Parse() (*ShortsConfig, error) {
	var c ShortsConfig
	err := envconfig.Process("", &c)
	return &c, err
}
