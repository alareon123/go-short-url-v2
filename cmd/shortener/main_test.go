package main

import (
	"bytes"
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
