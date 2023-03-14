package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// status обробник що повертає статус OK (підходить для різних роутерів)
func status(w http.ResponseWriter, req *http.Request) {
	// Стандартний роутер не підтримує метод як частину навігації
	if req.Method == "GET" {
		w.Write([]byte("OK!"))
	}
}

// standard повертає HTTP обробник що використовує вбудований роутер
func standard() http.Handler {
	// Стандартний роутер не підтримує параметри в URL
	// Шлях описаний в роутері працює як префікс, всі шляхи що починаються зі /status/ буду викликати statusById
	http.HandleFunc("/status/", statusById)
	http.HandleFunc("/status", status)

	return http.DefaultServeMux
}

// vendor повертає HTTP обробник що використовує httprouter
func vendor() http.Handler {
	router := httprouter.New()
	// За замовчуванням всі методи які не описані повертатимуть 405 помилку
	// За замовчуванням всі шляхи які не описані повертатимуть 404 помилку
	// httprouter підтримує змінні, але не regex вирази в URL
	router.HandlerFunc("GET", "/status/:id", statusById2)
	router.HandlerFunc("GET", "/status", status)

	return router
}

func main() {
	// router := standard()
	router := vendor()
	log.Fatal(http.ListenAndServe(":8080", router))
}

var elements = map[string]string{
	"0": "Справний",
	"1": "Обробляється",
	"2": "Помилка",
}

// status обробник що повертає статус в залежності від ID (стандартний роутер)
func statusById(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		parts := strings.Split(req.URL.String(), "/")
		id := parts[len(parts)-1]
		status, exists := elements[id]
		if exists {
			w.Write([]byte(status))
			return
		}
		w.Write([]byte("Невідомо"))
	}
}

// status обробник що повертає статус в залежності від ID (httprouter)
func statusById2(w http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())
	status, exists := elements[params.ByName("id")]
	if exists {
		w.Write([]byte(status))
		return
	}
	w.Write([]byte("Невідомо"))
}
