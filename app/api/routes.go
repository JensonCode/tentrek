package api

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func serve(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			WriteError(w, err)
		}
	}
}

func (s *Server) RegisterRoutes() {
	s.router.HandleFunc("/auth/user/{service}", serve(s.AuthHandlers))

	//OAuth
	s.router.HandleFunc("/auth/{provider}", serve(s.OAuthLogin))
	s.router.HandleFunc("/auth/{provider}/callback", serve(s.OAuthCallback))
	s.router.HandleFunc("/logout/{provider}", serve(s.OAuthLogout))
}

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
