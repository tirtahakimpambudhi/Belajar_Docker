package main

import (
	"fmt"
	"go_load_balance/handler"
	"log"
	"net"
	"net/http"
	"os"
)

func main()  {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8081"
	}
	mux := http.NewServeMux()
	handl := handler.NewHandler()
	mux.HandleFunc("GET /",handl.Get)
	listen , err := net.Listen("tcp",fmt.Sprintf(":%s",port))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Server running on the port %s \n",port)
	http.Serve(listen,mux)
}