package config

import "errors"

const (
	HTTP_PROXY   = "http"
	SOCKS5_PROXY = "socks5"
)

var (
	ErrProxyTypeNotFound = errors.New("proxy type not found")
)

type Config struct {
	ProxyType string
	Port      string
}

func ValidateConfig(cfg *Config) error {
	if cfg.ProxyType != HTTP_PROXY && cfg.ProxyType != SOCKS5_PROXY {
		return ErrProxyTypeNotFound
	}

	return nil
}
