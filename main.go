package main

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/sidgangs99/uniqueIdGenerator.git/src/health"
)

// TODO: import this from env variables
var port = 8000
var serverApi = "127.0.0.1:" + strconv.Itoa(port)

func main() {

	log.Info("Starting server, ", serverApi)
	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	// http.Handle("/", middleware.HttpInterceptor(router))
	mount(router, "/health", health.Router())
	
	err := http.ListenAndServe(serverApi, router)

	if err != nil {
        fmt.Println(err)
    }
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
    )
}
