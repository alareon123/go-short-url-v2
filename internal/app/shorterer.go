package app

import (
	"fmt"
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
	fmt.Printf("url %s stored with id %s", urlBase, urlShort)
	urls[urlShort] = urlBase
	fmt.Println(urls)
}

func getURL(urlShort string) string {
	return urls[urlShort]
}
