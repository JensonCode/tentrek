package api

import (
	"errors"
	"net/http"

	"github.com/JensonCode/tentrek/store"
	"github.com/gorilla/mux"
)

type UserUIDHandler struct {
	store *store.Store
}

func getUID(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["uid"]
}

func (s *Server) UserUIDHandlers(w http.ResponseWriter, r *http.Request) error {

	h := &UserUIDHandler{store: s.store}

	if r.Method == "GET" {
		return h.handleGetUser(w, r)
	}

	return errors.New("user uid route: request service is not found")
}

func (h *UserUIDHandler) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	uid := getUID(r)

	user, err := h.store.UserStore.FindByField("uid", uid)
	if err != nil {
		return err
	}

	return WriteResponse(w, http.StatusOK, user)
}
