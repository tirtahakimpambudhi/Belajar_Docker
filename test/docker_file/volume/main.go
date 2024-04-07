package main

import (
	"fmt"
	"go_volume/internal/handler"
	log "go_volume/internal/logging"
	"go_volume/internal/middleware"
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
	mux.HandleFunc("POST /item",handl.CreateItems)
	logging := middleware.NewLogging(log.SetupLogger())
	loggingMiddleware := logging.LoggingMiddleware(mux)
	recoverMiddleware := middleware.RecoverMiddleware(loggingMiddleware)
	http.ListenAndServe(fmt.Sprintf(":%s",port), recoverMiddleware)
}