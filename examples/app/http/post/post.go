package main

import (
	"io"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Write([]byte("Привіт! ✋"))
	case "POST":
		data, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(data)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
