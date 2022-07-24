package health

import (
"io"
"net/http"
log "github.com/sirupsen/logrus"
"github.com/gorilla/mux"
)

func Router() *mux.Router {
    router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/", health).Methods("GET")
    return router
}


func health(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health check: successful")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"ping": pong}`)
}