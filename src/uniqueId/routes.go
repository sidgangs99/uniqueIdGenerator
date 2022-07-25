package uniqueId

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
    router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/short", generateShortUniqueIdController).Methods("GET")
    return router
}


