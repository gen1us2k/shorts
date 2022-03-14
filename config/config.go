package config

import "github.com/kelseyhightower/envconfig"

const (
	KratosSessionKey = "ory_kratos_session"
	OwnerKey         = "owner_id"
	ProviderSupabase = "supabase"
	ProviderPostgres = "postgres"
)

// ShortsConfig stores configuration for the application
type ShortsConfig struct {
	BindAddr         string `envconfig:"BIND_ADDR" default:":8090"`
	DSN              string `envconfig:"DSN"`
	DatabaseProvider string `envconfig:"DATABASE_PROVIDER" default:"supabase"`
	URLLength        int    `envconfig:"URL_LENGTH" default:"8"`
	KratosAPIURL     string `envconfig:"KRATOS_API_URL" default:"http://kratos:4433"`
	UIURL            string `envconfig:"UI_URL" default:"http://127.0.0.1:4455/login"`
	SupabaseKey      string `envconfig:"SUPABASE_KEY"`
	SupabaseURL      string `envconfig:"SUPABASE_URL"`
}

// Parse parses environment variables and returns
// filled struct
func Parse() (*ShortsConfig, error) {
	var c ShortsConfig
	err := envconfig.Process("", &c)
	return &c, err
}
