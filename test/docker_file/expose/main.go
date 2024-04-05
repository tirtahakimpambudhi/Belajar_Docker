package main

import (
	"fmt"
	"go_expose/internal/handler"
	"net/http"
	"os"
)


func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8081"
	}
	mux := http.NewServeMux()
	handl := handler.NewItemHandler()
	
	mux.HandleFunc("GET /item/{id}",handl.GetItemByID)
	mux.HandleFunc("GET /items",handl.GetAll)
	fmt.Println("Start")
	http.ListenAndServe(fmt.Sprintf(":%s",port), mux)
}
