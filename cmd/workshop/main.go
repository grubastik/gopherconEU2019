package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/grubastik/gopherconEU2019/internal/app"
	logger "github.com/sirupsen/logrus"
)

func main() {
	logger := logger.New()
	logger.SetOutput(os.Stdout)
	logger.Infoln("app started")
	defer logger.Infoln("app finished")
	// Do not use for prod!!! Always customize server settings
	//http.ListenAndServe("localhost:8080", nil)
	port := os.Getenv("WORKSHOP_PORT")
	if port == "" {
		logger.Fatalln("WORKSHOP_PORT env var is not set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", app.HomeHandler1(logger))

	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		logger.Fatalln("error starting server: " + err.Error())
	}
}
