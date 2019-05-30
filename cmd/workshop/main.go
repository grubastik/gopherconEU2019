package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/grubastik/gopherconEU2019/internal/app"
)

func main() {
	fmt.Println("app started")
	// Do not use for prod!!! Always customize server settings
	//http.ListenAndServe("localhost:8080", nil)
	port := os.Getenv("WORKSHOP_PORT")
	if port == "" {
		fmt.Println("WORKSHOP_PORT env var is not set")
		os.Exit(1)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", app.HomeHandler)

	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: r,
	}
	server.ListenAndServe()

	defer fmt.Println("app finished")
}
