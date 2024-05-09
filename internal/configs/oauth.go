package configs

import (
	"fmt"
	"os"
)

type OAuthConf struct {
	SessionKey   string
	ClientID     string
	ClientSecret string
	CallbackUrl  string
}

func NewOAuthConf() *OAuthConf {
	base := os.Getenv("API_BASE_URL")
	url := fmt.Sprintf("%s/auth/google/callback", base)

	return &OAuthConf{
		SessionKey:   os.Getenv("SESSION_KEY"),
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		CallbackUrl:  url,
	}
}
