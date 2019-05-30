package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/grubastik/gopherconEU2019/internal/app"
	"github.com/grubastik/gopherconEU2019/internal/diagnostics"
	logger "github.com/sirupsen/logrus"
)

func main() {
	//logger
	logger := logger.New()
	logger.SetOutput(os.Stdout)
	logger.Infoln("app started. Version:", diagnostics.Version, " Commit: ", diagnostics.Commit, " BuildTime: ", diagnostics.BuildTime)
	defer logger.Infoln("app finished")
	// Do not use for prod!!! Always customize server settings
	//http.ListenAndServe("localhost:8080", nil)

	//env vars
	port := os.Getenv("WORKSHOP_PORT")
	if port == "" {
		logger.Fatalln("WORKSHOP_PORT env var is not set")
	}

	//signals and shutdown
	interrupt := make(chan os.Signal, 1) //need bufferred because can have multiple interruptions in a row
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	shutdown := make(chan error, 1)

	//handlers
	r := mux.NewRouter()
	r.HandleFunc("/", app.HomeHandler1(logger))
	r.HandleFunc("/healthz", diagnostics.HealthHandler(logger))
	r.HandleFunc("/readyz", diagnostics.ReadyHandler(logger))

	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: r,
	}

	//start server
	go func() {
		err := server.ListenAndServe()
		shutdown <- err
	}()

	logger.Infoln("Serving requests")

	select {
	case signal := <-interrupt:
		switch signal {
		case os.Interrupt:
			logger.Infoln("Interrupted")
		case syscall.SIGTERM:
			logger.Infoln("Got sygterm")
		}
	case <-shutdown:
		logger.Infoln("got shutdown")
	}
	err := server.Shutdown(context.Background())
	if err != nil {
		logger.Fatalln("Error stopping server: " + err.Error())
	}
}
