package app

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func HomeHandler1(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infoln("processing home request")
		w.WriteHeader((http.StatusOK))
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader((http.StatusOK))
}
