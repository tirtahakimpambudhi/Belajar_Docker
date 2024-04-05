package main

import (
	"fmt"
	"go_expose/internal/handler"
	"go_expose/internal/log"
	"go_expose/internal/middleware"
	"log/slog"
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
	logging := middleware.NewLogging(slog.New(log.NewHandler(nil)))
	loggingMiddleware := logging.LoggingMiddleware(mux)
	recoverMiddleware := middleware.RecoverMiddleware(loggingMiddleware)
	fmt.Println("Start")
	http.ListenAndServe(fmt.Sprintf(":%s",port), recoverMiddleware)
}
