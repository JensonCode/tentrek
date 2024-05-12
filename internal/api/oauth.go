package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/JensonCode/tentrek/jwt"
	"github.com/JensonCode/tentrek/model"
	"github.com/gorilla/mux"

	"github.com/markbates/goth/gothic"
)

func (s *Server) OAuthLogin(w http.ResponseWriter, r *http.Request) error {
	_, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println(err)
		gothic.BeginAuthHandler(w, r)
	}

	return nil
}

func (s *Server) OAuthLogout(w http.ResponseWriter, r *http.Request) error {
	err := gothic.Logout(w, r)
	if err != nil {
		return err
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)

	return nil
}

func (s *Server) OAuthCallback(w http.ResponseWriter, r *http.Request) error {

	provider := mux.Vars(r)["provider"]

	gothicUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}

	user, _ := s.store.UserStore.FindByField("email", gothicUser.Email)
	if user == nil {
		newUser := &model.CreateUserRequest{
			Email:    gothicUser.Email,
			Provider: provider,
			Avatar:   gothicUser.AvatarURL,
		}

		u, err := s.store.UserStore.InsertUser(newUser)
		if err != nil {
			return err
		}

		user = u
	}

	token, err := jwt.GenerateToken(user.UID)
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
