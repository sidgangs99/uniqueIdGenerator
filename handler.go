package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HealthCheck
func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{ping:pong}")
}
