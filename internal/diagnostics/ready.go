package diagnostics

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func ReadyHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infoln("processing ready request")
		w.WriteHeader((http.StatusOK))
	}
}
