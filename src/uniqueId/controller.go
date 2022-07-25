package uniqueId

import (
	"net/http"
	"encoding/json"
)

func generateShortUniqueIdController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	shortId := Result{}
	shortId.UniqueId = generateShortUniqueIdService()
	json.NewEncoder(w).Encode(shortId)
}