package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func jsonProxy(w http.ResponseWriter, r *http.Request) {
	j := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(j); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := httprouter.New()
	router.HandlerFunc("GET", "/json", loggerMW2(jsonProxy))

	log.Fatal(http.ListenAndServe(":8080", router))
}
