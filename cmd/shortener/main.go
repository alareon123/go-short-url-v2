package main

import (
	"encoding/json"
	"github.com/alareon123/go-short-url.git/internal/app"
	"github.com/alareon123/go-short-url.git/internal/config"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

type shortUrl struct {
	Url string `json:"url"`
}

type resultUrl struct {
	Result string `json:"result"`
}

func urlShortHandler(w http.ResponseWriter, r *http.Request) {
	reqBodyBytes, _ := io.ReadAll(r.Body)

	if len(reqBodyBytes) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shortURL := app.ShortURL(string(reqBodyBytes))

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(shortURL))
	if err != nil {
		log.Fatal("error while writing response")
	}
}

func getURLHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	longURL := app.GetURLByID(id)

	if longURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func apiShortUrl(w http.ResponseWriter, r *http.Request) {

	var shortUrlJson shortUrl
	var resultJson resultUrl

	if err := json.NewDecoder(r.Body).Decode(&shortUrlJson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortUrlID := app.ShortURL(shortUrlJson.Url)

	resultJson.Result = shortUrlID

	jsonData, _ := json.Marshal(resultJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(jsonData)
	if err != nil {
		log.Fatal("error while writing response")
	}
}

func main() {
	config.Init()

	r := chi.NewRouter()

	r.Method("POST", "/", app.RequestLogger(urlShortHandler))
	r.Method("GET", "/{id}", app.RequestLogger(getURLHandler))
	r.Method("POST", "/api/shorten", app.RequestLogger(apiShortUrl))

	http.ListenAndServe(config.AppServerURL, r)
}
