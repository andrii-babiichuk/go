package main

import (
	"log"
	"net/http"
)

// writerWrapper wraps response writer to add custom logic
type writerWrapper struct {
	http.ResponseWriter
	status int
}

func (w *writerWrapper) WriteHeader(code int) {
	w.status = code
}

func (w *writerWrapper) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}

// loggerMW логує URL та статус відповіді запитів
func loggerMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wr := &writerWrapper{ResponseWriter: w}

		next(wr, r)

		log.Printf("Request URL: %s, Response status: %d\n", r.URL.String(), wr.status)
	}
}
