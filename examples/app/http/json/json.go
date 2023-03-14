package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type entity struct {
	Bool   bool           `json:"bool"`
	Int    int            `json:"int"`
	String string         `json:"string"`
	Slice  []int          `json:"slice"`
	Map    map[int]string `json:"map"`
	// Приватні поля не беруть участь перетворенні JSON
	hidden string `json:"hidden"`
}

// У випадку, якщо нульові значення еквівалент
// https://www.sohamkamani.com/golang/omitempty/
type entity2 struct {
	B      bool           `json:"bool,omitempty"`
	I      int            `json:"int,omitempty"`
	S      string         `json:"string,omitempty"`
	SL     []int          `json:"slice,omitempty"`
	M      map[int]string `json:"map,omitempty"`
	hidden string         `json:"hidden,omitempty"`
}

// Нестандартна конвертація json
// https://calvinfeng.gitbook.io/gonotebook/idioms/custom-json-marshaling
type entity3 = entity2

func (e *entity3) UnmarshalJSON(data []byte) error {
	type entity3Alias entity3
	aux := struct {
		*entity3Alias
		I string `json:"int"`
	}{
		entity3Alias: (*entity3Alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	i, err := strconv.Atoi(aux.I)
	if err != nil {
		fmt.Println(err)
	}
	e.I = i

	return nil
}

// func (e *entity3) MarshalJSON() ([]byte, error) {}

func main() {
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/json2", jsonHandler2)
	http.HandleFunc("/json3", jsonHandler3)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func jsonHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var e entity
		err := json.NewDecoder(req.Body).Decode(&e)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(e)
		err = json.NewEncoder(w).Encode(&e)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}
}

func jsonHandler2(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var e entity2
		err := json.NewDecoder(req.Body).Decode(&e)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(e)
		err = json.NewEncoder(w).Encode(&e)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}
}

func jsonHandler3(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var e entity3
		err := json.NewDecoder(req.Body).Decode(&e)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(e)
		err = json.NewEncoder(w).Encode(&e)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}
}
