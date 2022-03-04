package config

type ShortsConfig struct {
	BindAddr string `envconfig:"BIND_ADDR"`
	DSN      string `envconfig:"DSN"`
}
