package config

import (
	"flag"
	"os"
)

var AppServerURL string
var BaseAddressURL string

func Init() {
	flag.StringVar(&AppServerURL, "a", ":8080", "address and port to run server")
	flag.StringVar(&BaseAddressURL, "b", "http://localhost:8080", "base address for short urls")
	flag.Parse()

	if envAppURL := os.Getenv("SERVER_ADDRESS"); envAppURL != "" {
		AppServerURL = envAppURL
	}
	if envBaseAddressURL := os.Getenv("BASE_URL"); envBaseAddressURL != "" {
		BaseAddressURL = envBaseAddressURL
	}
	BaseAddressURL += "/"
}
