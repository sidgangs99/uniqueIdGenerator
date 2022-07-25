package health

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health check: successful")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"ping": "pong"}`)
}