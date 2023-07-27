package app

import (
	"github.com/alareon123/go-short-url.git/internal/config"
	"io"
	"strconv"
)

var urls = make(map[string]string)

const fileName = "tmp/short-url-db.json"

var baseUUID = 0

func init() {
	consumer, err := NewConsumer(fileName)
	if err != nil {
		Logger.Fatal(err.Error())
	}
	defer consumer.Close()

	for {
		readEvent, err := consumer.ReadEvent()
		if err == io.EOF || readEvent == nil {
			break
		}
		urls[readEvent.ShortUrl] = readEvent.OriginalUrl
		baseUUID++
	}
}

func ShortURL(url string) string {
	randString := RandStringBytes(8)
	storeURL(url, randString)
	return config.BaseAddressURL + randString
}

func GetURLByID(shortURL string) string {
	return getURL(shortURL)
}

func storeURL(urlBase string, urlShort string) {
	baseUUID++
	urls[urlShort] = urlBase
	producer, err := NewProducer(fileName)
	if err != nil {
		Logger.Fatal(err.Error())
	}
	defer producer.CloseFile()
	writeEvent := UrlEvent{
		OriginalUrl: urlBase,
		ShortUrl:    urlShort,
		Uuid:        strconv.Itoa(baseUUID),
	}
	err = producer.WriteEvent(&writeEvent)
	if err != nil {
		Logger.Fatal(err.Error())
	}
	err = producer.CloseFile()
	if err != nil {
		Logger.Fatal(err.Error())
	}
}

func getURL(urlShort string) string {
	return urls[urlShort]
}
