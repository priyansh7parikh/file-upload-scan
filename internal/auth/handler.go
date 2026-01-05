package auth

import (
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	url := OAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	// replacing the token with _ as not used , can be used later if needed
	_, err := ExchangeCode(r.Context(), code)
	if err != nil {
		http.Error(w, "oauth exchange failed", http.StatusUnauthorized)
		return
	}

	// 🔑 Normally: fetch user info from provider
	// For now, mock user
	userID := "oauth-user-123"
	role := "user"

	jwtToken, err := GenerateJWT(userID, role)
	if err != nil {
		http.Error(w, "jwt generation failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"access_token": jwtToken,
	})
}
