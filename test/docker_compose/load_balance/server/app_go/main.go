package main

import (
	"net/http"
	"go_load_balance/handler"
)

func main()  {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	handl := handler.NewHandler()
	mux.HandleFunc("GET /",handl.Get)
}