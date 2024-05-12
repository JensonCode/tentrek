package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/JensonCode/tentrek/email"
	"github.com/JensonCode/tentrek/helpers"
	"github.com/JensonCode/tentrek/jwt"
	"github.com/JensonCode/tentrek/model"
	"github.com/JensonCode/tentrek/store"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	store *store.Store
}

func (s *Server) AuthHandlers(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errors.New("auth route: request service is not found")
	}

	vars := mux.Vars(r)
	service := vars["service"]

	h := &AuthHandler{store: s.store}

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

func (h *AuthHandler) handleLogin(w http.ResponseWriter, r *http.Request) error {

	req := new(model.LoginReqeust)
	err := ParseRequest(r, req)
	if err != nil {
		return err
	}

	user, err := h.store.UserStore.FindByField("email", req.Email)
	if err != nil || user == nil {
		return errors.New("email not found")
	}

	ok := helpers.CompareHashAndPassword(user.Password, req.Password)
	if !ok {
		return errors.New("password is incorrect")
	}

	token, err := jwt.GenerateToken(user.UID)
	if err != nil {
		return err
	}

	return WriteResponse(w, http.StatusOK, map[string]string{"access_token": token})
}

func (h *AuthHandler) handleRegister(w http.ResponseWriter, r *http.Request) error {

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

	token, err := jwt.GenerateToken(user.UID)
	if err != nil {
		return err
	}

	return WriteResponse(w, http.StatusOK, map[string]string{"access_token": token})

}

func (h *AuthHandler) handleEmailVerification(w http.ResponseWriter, r *http.Request) error {
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
		return fmt.Errorf("cannot send verification code to email: %s", req.Email)
	}

	uuid := h.store.AuthStore.StoreOTP(*req, otp)

	return WriteResponse(w, http.StatusOK, map[string]string{"register_id": uuid})

}
