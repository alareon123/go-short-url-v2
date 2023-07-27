package config

import (
	"flag"
	"os"
)

var AppServerURL string
var BaseAddressURL string
var FileStoragePath string

func Init() {
	flag.StringVar(&AppServerURL, "a", ":8080", "address and port to run server")
	flag.StringVar(&BaseAddressURL, "b", "http://localhost:8080", "base address for short urls")
	flag.StringVar(&FileStoragePath, "f", "/tmp/short-url-db.json", "file storage path")
	flag.Parse()

	if envAppURL := os.Getenv("SERVER_ADDRESS"); envAppURL != "" {
		AppServerURL = envAppURL
	}
	if envBaseAddressURL := os.Getenv("BASE_URL"); envBaseAddressURL != "" {
		BaseAddressURL = envBaseAddressURL
	}
	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		FileStoragePath = envFileStoragePath
	}
	BaseAddressURL += "/"
}
