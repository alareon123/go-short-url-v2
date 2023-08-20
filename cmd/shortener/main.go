package main

import (
	"context"
	"encoding/json"
	"github.com/alareon123/go-short-url.git/internal/app"
	"github.com/alareon123/go-short-url.git/internal/config"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

type shortURL struct {
	URL string `json:"url"`
}

type resultURL struct {
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

func apiShortURL(w http.ResponseWriter, r *http.Request) {

	var shortURLJson shortURL
	var resultJSON resultURL

	if err := json.NewDecoder(r.Body).Decode(&shortURLJson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURLID := app.ShortURL(shortURLJson.URL)

	resultJSON.Result = shortURLID

	jsonData, _ := json.Marshal(resultJSON)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write(jsonData)
	if err != nil {
		log.Fatal("error while writing response")
	}
}

func dataBasePing(w http.ResponseWriter, r *http.Request) {
	var connection *app.DBConnection
	if config.DataBaseCredentials != "" {
		connection = app.ConnectToDataBase(config.DataBaseCredentials)
	}

	err := connection.Conn.Ping(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func initRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Method("POST", "/", middleware(urlShortHandler))
	r.Method("GET", "/{id}", middleware(getURLHandler))
	r.Method("POST", "/api/shorten", middleware(apiShortURL))
	r.Method("GET", "/ping", middleware(dataBasePing))

	return r
}

func main() {
	config.Init()
	app.Init()

	http.ListenAndServe(config.AppServerURL, initRouter())
}

func middleware(h http.HandlerFunc) http.Handler {
	loggedFunc := app.RequestLogger(h)
	compressedFunc := app.GzipMiddleware(loggedFunc)
	return compressedFunc
}
