package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	//router.HandlerFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "build/index.html")
	//})
	//router.ServeFiles("/static/*filepath", http.Dir("build/static"))

	router.NotFound = http.FileServer(http.Dir("build"))
	router.HandlerFunc("GET", "/api/status", status)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func status(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
	})
}
