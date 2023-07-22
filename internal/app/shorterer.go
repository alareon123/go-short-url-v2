package app

import (
	"github.com/alareon123/go-short-url.git/internal/config"
)

var urls = make(map[string]string)

func ShortURL(url string) string {
	randString := RandStringBytes(8)
	storeURL(url, randString)
	return config.BaseAddressURL + randString
}

func GetURLByID(shortURL string) string {
	return getURL(shortURL)
}

func storeURL(urlBase string, urlShort string) {
	urls[urlShort] = urlBase
}

func getURL(urlShort string) string {
	return urls[urlShort]
}
