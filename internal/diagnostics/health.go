package diagnostics

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func HealthHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infoln("processing health request")
		w.WriteHeader((http.StatusOK))
		w.Write([]byte("OK"))
	}
}
