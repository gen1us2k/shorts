package config

type ShortsConfig struct {
	BindAddr  string `envconfig:"BIND_ADDR"`
	DSN       string `envconfig:"DSN"`
	URLLength int    `envconfig:"URL_LENGTH" default:"8"`
}
