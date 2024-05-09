package api

import (
	"errors"
	"net/http"

	"github.com/JensonCode/tentrek/internal/auth"
	"github.com/JensonCode/tentrek/internal/email"
	"github.com/JensonCode/tentrek/internal/model"
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

	req := new(model.LoginReqeust)
	err := ParseRequest(r, req)
	if err != nil {
		return err
	}

	user, err := h.store.UserStore.FindByField("email", req.Email)
	if err != nil || user == nil {
		return errors.New("user email not found")
	}

	ok := auth.CompareHashAndPassword(user.Password, req.Password)
	if !ok {
		return errors.New("password is incorrect")
	}

	token, err := auth.GenerateJWTToken(user.UID)
	if err != nil {
		return err
	}

	return WriteResponse(w, http.StatusOK, map[string]string{"access_token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) error {

	req := new(model.EmailVerificationRequest)
	err := ParseRequest(r, req)
	if err != nil {
		return err
	}

	newUserReq, err := h.store.AuthStore.DeleteOTP(req)
	if err != nil {
		return err
	}

	user, err := h.store.UserStore.InsertUser(newUserReq)
	if err != nil {
		return err
	}

	token, err := auth.GenerateJWTToken(user.UID)
	if err != nil {
		return err
	}

	return WriteResponse(w, http.StatusOK, map[string]string{"access_token": token})

}

func (h *Handler) handleEmailVerification(w http.ResponseWriter, r *http.Request) error {
	req := new(model.CreateUserRequest)
	err := ParseRequest(r, req)
	if err != nil {
		return err
	}

	exist, err := h.store.UserStore.IsExist(req.Email)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("user email has been used")
	}

	otp, err := email.SendOTP(req.Email)
	if err != nil {
		return errors.New("cannot send verification to your email")
	}

	uuid := h.store.AuthStore.StoreOTP(*req, otp)

	return WriteResponse(w, http.StatusOK, map[string]string{"register_id": uuid})

}
