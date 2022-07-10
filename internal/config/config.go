package config

import (
	"errors"
	"os"
)

const (
	HTTP_PROXY   = "http"
	SOCKS5_PROXY = "socks5"

	DB_USER_ENV = "DB_USER"
	DB_PASS_ENV = "DB_PASS"
)

var (
	ErrProxyTypeNotFound = errors.New("proxy type not found")
)

type Config struct {
	ProxyType string
	Port      string

	DbUser     string
	DbPassword string
}

func InitializeConfig(cfg *Config) (*Config, error) {
	if cfg.ProxyType != HTTP_PROXY && cfg.ProxyType != SOCKS5_PROXY {
		return nil, ErrProxyTypeNotFound
	}

	cfg.DbUser = os.Getenv(DB_USER_ENV)
	cfg.DbPassword = os.Getenv(DB_PASS_ENV)

	return cfg, nil
}
