package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		written, err := io.WriteString(w, "Привіт! ✋")
		// written, err := w.Write([]byte("Привіт! ✋"))
		fmt.Println(written, err)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
