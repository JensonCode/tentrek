package auth

import (
	"os"

	"github.com/JensonCode/tentrek/internal/configs"
	"github.com/gorilla/sessions"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	MaxAge = 86400 * 30
)

func NewOAuth() {

	conf := configs.NewOAuthConf()

	store := sessions.NewCookieStore([]byte(conf.SessionKey))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd()

	gothic.Store = store

	goth.UseProviders(
		google.New(conf.ClientID, conf.ClientSecret, conf.CallbackUrl),
	)
}

func isProd() bool {
	mode := os.Getenv("APP_MODE")
	return mode == "prod"
}
