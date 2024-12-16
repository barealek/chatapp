package middleware

import (
	"encoding/json"
	"net/http"
)

func unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Unauthorized",
	})
}

func internalservererror(w http.ResponseWriter, err, errid string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err,
		"code":  errid,
	})
}
