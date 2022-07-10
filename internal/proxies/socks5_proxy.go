package proxies

import (
	"errors"

	"github.com/renatobrittoaraujo/proxy-king/internal/config"
)

var (
	ErrNotImplemented = errors.New("socks5 not implemented")
)

func SOCKS5Listen(cfg *config.Config) error {
	return ErrNotImplemented
}
