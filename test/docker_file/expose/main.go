package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type item struct {
	Name  string    `json:"name"`
	Qtt   int       `json:"quantity"`
	Price *currency `json:"price"`
}

type currency struct {
	Code   string  `json:"code"`
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
}

func writeJSON(w http.ResponseWriter, code int, data any) {
	dataJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(w, "error converting data to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(dataJSON)
}

func main() {
	stores := make(map[string]*item)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /items", func(w http.ResponseWriter, r *http.Request) {
		var items []*item
		if len(stores) == 0 {
			fmt.Fprint(w,"items is empty")
			return
		}
		for _, item := range stores {
			items = append(items, item)	
		}
		writeJSON(w,http.StatusOK,items)
		return
	})
	mux.HandleFunc("GET /item/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		item, ok := stores[id]
		if !ok {
			http.Error(w, fmt.Sprintf("item with id %s not found", id), http.StatusNotFound)
			return
		}
		writeJSON(w,http.StatusOK,item)
		return
	})
	http.ListenAndServe(":8081", mux)
	fmt.Println("Start")
}
