package uniqueId

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func generateShortUniqueIdController(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health check: successful")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	const id = "493830"

	io.WriteString(w, `{"uniqueId": id}`)
}