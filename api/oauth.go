package api

import (
	"net/http"
	"os"

	"github.com/JensonCode/tentrek/internal/auth"
	"github.com/JensonCode/tentrek/internal/model"
	"github.com/gorilla/mux"

	"github.com/markbates/goth/gothic"
)

func (h *Handler) OAuthLogin(w http.ResponseWriter, r *http.Request) error {
	// try to get the user without re-authenticating
	_, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		gothic.BeginAuthHandler(w, r)
	}

	return nil
}

func (h *Handler) OAuthLogout(w http.ResponseWriter, r *http.Request) error {
	err := gothic.Logout(w, r)
	if err != nil {
		return err
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)

	return nil
}

func (h *Handler) OAuthCallback(w http.ResponseWriter, r *http.Request) error {

	provider := mux.Vars(r)["provider"]

	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}

	user, _ := h.store.UserStore.FindByField("email", gothUser.Email)
	if user == nil {
		req := &model.CreateUserRequest{
			Email:    gothUser.Email,
			Provider: provider,
			Avatar:   gothUser.AvatarURL,
		}

		u, err := h.store.UserStore.InsertUser(req)
		if err != nil {
			return err
		}

		user = u
	}

	token, err := auth.GenerateJWTToken(user.UID)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})

	url := os.Getenv("APP_BASE_URL")

	http.Redirect(w, r, url, http.StatusFound)

	return nil
}
