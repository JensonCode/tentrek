package api

import (
	"encoding/json"
	"net/http"
)

func ParseRequest[T any](r *http.Request, req *T) error {
	return json.NewDecoder(r.Body).Decode(req)
}

func WriteResponse(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, err error) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
