package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	oauthURL = "https://oauth2.googleapis.com/tokeninfo?id_token="
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

	req, err := http.NewRequest("GET", oauthURL+token, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Printf("Error calling backend: %v", err)
		return false
	}
	defer resp.Body.Close()

	logger.Printf("Response status: %s", resp.Status)
	if resp.StatusCode != http.StatusOK {
		logger.Printf("Invalid response code: %v", resp.StatusCode)
		return false
	}

	dataMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&dataMap)

	logger.Printf("Profile: %v", dataMap)

	//TODO: Check the validity of aud
	//      Must claim contains role or app ID?

	return true

}
