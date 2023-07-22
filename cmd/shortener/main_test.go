package main

import (
	"bytes"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_urlShortHandler(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		emptyBody   bool
	}

	tests := []struct {
		name string
		body []byte
		want want
	}{
		{
			name: "positive test",
			body: []byte("ya.ru"),
			want: struct {
				contentType string
				statusCode  int
				emptyBody   bool
			}{
				contentType: "text/plain",
				statusCode:  201,
				emptyBody:   false,
			},
		},
		{
			name: "empty body test",
			want: struct {
				contentType string
				statusCode  int
				emptyBody   bool
			}{
				contentType: "text/plain",
				statusCode:  400,
				emptyBody:   true,
			},
		},
		{
			name: "positive api test",
			want: struct {
				contentType string
				statusCode  int
				emptyBody   bool
			}{
				contentType: "application/json",
				statusCode:  200,
				emptyBody:   false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewReader(tt.body))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(urlShortHandler)
			h(w, request)

			result := w.Result()
			resultBodyBytes, _ := io.ReadAll(result.Body)
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.emptyBody, len(resultBodyBytes) == 0)
		})
	}
}

func Test_urlShortHandlerApi(t *testing.T) {
	handler := http.HandlerFunc(apiShortUrl)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	type want struct {
		contentType string
		statusCode  int
		respBody    string
	}

	tests := []struct {
		name string
		body string
		want want
	}{
		{
			name: "positive api test",
			want: struct {
				contentType string
				statusCode  int
				respBody    string
			}{
				contentType: "application/json",
				statusCode:  200,
				respBody:    "false",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := resty.New().R()

			req.Method = "POST"
			req.Body = "{\n  \"url\": \"https://practicum.yandex.ru\"\n}"
			req.URL = srv.URL
			req.SetHeader("Content-Type", "application/json")

			resp, err := req.Send()
			assert.Empty(t, err)
			assert.NotEmpty(t, resp.Body())
		})
	}
}
