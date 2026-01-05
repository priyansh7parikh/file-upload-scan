package auth

import (
	"context"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var OAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
	Scopes:       []string{"email", "profile"},
	Endpoint:     google.Endpoint,
}

func ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return OAuthConfig.Exchange(ctx, code)
}
