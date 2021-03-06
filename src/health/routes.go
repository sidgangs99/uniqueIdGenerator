package health

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
    router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/", health).Methods("GET")
    return router
}


