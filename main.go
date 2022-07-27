package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/sidgangs99/uniqueIdGenerator.git/src/health"
	"github.com/sidgangs99/uniqueIdGenerator.git/src/middleware"
	"github.com/sidgangs99/uniqueIdGenerator.git/src/uniqueId"
)

// TODO: import this from env variables
var port = ":8000"

// TODO: Add log levels to the environment	
// TODO: Error handlers should be created here instead
// TODO: API path should be relative
func main() {

	log.Info("Starting server, ", port)
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.LoggingMiddleware)
	
	mount(router, "/health", health.Router())
	mount(router, "/uniqueId", uniqueId.Router())

	err := http.ListenAndServe(port, router)

	if err != nil {
        fmt.Println(err)
    }
}



func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
    )
}
