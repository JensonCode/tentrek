package api

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) AuthHandlers(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	service := vars["service"]

	if r.Method != "POST" {
		return errors.New("auth route: request service is not found")
	}

	if service == "login" {
		return h.handleLogin(w, r)
	}
	if service == "otp" {
		return h.handleEmailVerification(w, r)
	}
	if service == "register" {
		return h.handleRegister(w, r)
	}

	return nil

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) error {

	return WriteResponse(w, http.StatusOK, "login")
}

func (h *Handler) handleEmailVerification(w http.ResponseWriter, r *http.Request) error {

	return WriteResponse(w, http.StatusOK, "register")

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) error {

	return WriteResponse(w, http.StatusOK, "register")

}
