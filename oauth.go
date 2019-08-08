package main

import (
	"net/http"
	"strings"

	oa "google.golang.org/api/oauth2/v2"
)

const (
	oauthURL = "https://oauth2.googleapis.com/tokeninfo?id_token="
)

var (
	httpClient = &http.Client{}
)

func auth(r *http.Request) bool {

	var token string
	tokens, ok := r.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = strings.TrimPrefix(tokens[0], "Bearer ")
	}

	if token == "" {
		logger.Println("Token not set")
		return false
	}

	oaSrv, err := oa.New(httpClient)
	if err != nil {
		logger.Printf("Error creating OAuth client: %v", err)
		return false
	}

	info, err := oaSrv.Tokeninfo().IdToken(token).Do()
	if err != nil {
		logger.Printf("Error validating token: %v", err)
		return false
	}
	logger.Printf("Token: %+v", info)

	// TODO: Validate host portion of audience is equal to the current request host
	// hosts := r.Header["Authorization"]
	// if info.Audience != host {
	// 	logger.Printf("Token for invalid client ID: %s", info.Audience)
	// 	return false
	// }

	if !info.VerifiedEmail {
		logger.Printf("Token for unverified email: %s", info.Audience)
		return false
	}

	return true

}
