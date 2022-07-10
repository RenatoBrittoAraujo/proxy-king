package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/renatobrittoaraujo/proxy-king/internal/config"
	"github.com/renatobrittoaraujo/proxy-king/internal/proxies"
)

var (
	fs *flag.FlagSet

	proxyType string
	port      int
	help      bool

	err error
)

func init() {
	fs = flag.NewFlagSet("proxy-king", flag.ExitOnError)

	fs.StringVar(&proxyType, "t", "http", "proxy type - may be 'http' or 'socks5', defaults to http")
	fs.IntVar(&port, "p", 80, "proxy port - defaults to '80'")
	fs.BoolVar(&help, "h", false, "shows all commands")
	fs.Parse(os.Args[1:])
}

func main() {
	if help {
		fs.Usage()
		return
	}

	cfg := &config.Config{
		ProxyType: proxyType,
		Port:      strconv.Itoa(port),
	}

	err = config.ValidateConfig(cfg)
	if err != nil {
		fmt.Println("Failed to validate input: ", err)
	}

	fmt.Println("Proxing port", cfg.Port+"...")

	switch proxyType {
	case config.HTTP_PROXY:
		err = proxies.HTTPListen(cfg)
	case config.SOCKS5_PROXY:
		fmt.Println("Socks5 is not supported as of now")
	}

	if err != nil {
		fmt.Println("Proxy server crashed: ", err)
	}
}
