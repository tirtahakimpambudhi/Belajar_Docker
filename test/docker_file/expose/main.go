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

type generalResponse struct {
	Message string 		`json:"message"`
	Data any			`json:"data"`
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

func getRequestJSON(w http.ResponseWriter,r *http.Request, field any) {
	err := json.NewDecoder(r.Body).Decode(&field)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
}

func main() {
	stores := make(map[string]*item)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /item",func(w http.ResponseWriter, r *http.Request) {
		var requestBody item
		getRequestJSON(w,r,&requestBody)
	
		fmt.Println("Received JSON data:", requestBody)
		fmt.Fprintln(w,"hello")
	})
	mux.HandleFunc("GET /items", func(w http.ResponseWriter, r *http.Request) {
		var items []*item
		if len(stores) == 0 {
			writeJSON(w,http.StatusNotFound,&generalResponse{Message: "items is empty", Data: nil})
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
			writeJSON(w,http.StatusNotFound,&generalResponse{Message: fmt.Sprintf("item with id %s not found", id),Data: nil})
			return
		}
		writeJSON(w,http.StatusOK,item)
		return
	})
	fmt.Println("Start")
	http.ListenAndServe(":8081", mux)
}
