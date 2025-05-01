package resp

import (
	"encoding/json"
	"net/http"
)

func SetJson(w http.ResponseWriter, res any, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
