package app

import (
	"github.com/alareon123/go-short-url.git/internal/config"
	"io"
	"strconv"
)

var urls = make(map[string]string)

var baseUUID = 0

func Init() {
	consumer, err := NewConsumer(config.FileStoragePath)
	if err != nil {
		Logger.Fatal(err.Error())
	}
	defer consumer.Close()

	for {
		readEvent, err := consumer.ReadEvent()
		if err == io.EOF || readEvent == nil {
			break
		}
		urls[readEvent.ShortURL] = readEvent.OriginalURL
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
	producer, err := NewProducer(config.FileStoragePath)
	if err != nil {
		Logger.Fatal(err.Error())
	}
	defer func(producer *Producer) {
		err := producer.CloseFile()
		if err != nil {
			Logger.Fatal("error happened while closing file")
		}
	}(producer)
	writeEvent := URLEvent{
		OriginalURL: urlBase,
		ShortURL:    urlShort,
		UUID:        strconv.Itoa(baseUUID),
	}
	err = producer.WriteEvent(&writeEvent)
	if err != nil {
		Logger.Fatal(err.Error())
	}
	if err != nil {
		Logger.Fatal(err.Error())
	}
}

func getURL(urlShort string) string {
	return urls[urlShort]
}
