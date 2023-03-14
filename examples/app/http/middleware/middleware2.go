package main

import (
	"io"
	"log"
	"net/http"
)

// writerWrapper wraps response writer to add custom logic
type writerWrapper2 struct {
	http.ResponseWriter
	status int
}

func (w *writerWrapper2) WriteHeader(code int) {
	w.status = code
}

func (w *writerWrapper2) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}

// loggerMW2 логує URL body та статус відповіді
func loggerMW2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// r.Body = io.NopCloser(bytes.NewReader(body))

		wr := &writerWrapper2{ResponseWriter: w}

		next(wr, r)

		log.Printf("Request URL: %s; Body: %s; Response status: %d\n", r.URL.String(), string(body), wr.status)
	}
}
