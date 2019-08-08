package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserClient(t *testing.T) {

	token := "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IjM0OTRiMWU3ODZjZGFkMDkyZTQyMzc2NmJiZTM3ZjU0ZWQ4N2IyMmQiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJodHRwczovL3ByZXByb2Nlc3Nvci0yZ3RvdW9zMnBxLXVjLmEucnVuLmFwcC92MS9wdXNoIiwiYXpwIjoiMTA5OTY4ODAxMzA4MzYzNjE2ODMxIiwiZW1haWwiOiJwcmVwcm9jZXNzb3Itc2FAY2xvdWR5bGFicy5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJleHAiOjE1NjUyMzI2MDksImlhdCI6MTU2NTIyOTAwOSwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tIiwic3ViIjoiMTA5OTY4ODAxMzA4MzYzNjE2ODMxIn0.pHKaKOO5myk0Rwvm-zCSY_W7QUpcPmBDoOCW-8_1KhtPjwG8n0gsoN-ZzaqJwvRrjT4JrnwHCHuTVkaoIDVuzToDWwxCR9QM9Y2mNpLsL8Y6RR9rD84DVot03ovn3NASy7Kjs7EQWrA_iGhgoBjIUsIguoNqgAV6Uv7xgyVSZC9Aw6Y0Mhwu1j4Qh4eN5JUSNgPZNBwCfLMIDhF7j8YtlXu0Kqfc_M6QFATbMi_R6Ej9kz9LMljB7lu0RX01U1L06uKveWzsudeGze0z1tccHNtOHDwZcDi3C0Ug4rt1lpk1jBYIxU51p4yeesWGyfY28uVzzgimatgEeKJ7lUa42A"

	req, err := http.NewRequest("GET", "/v1/api", nil)
	assert.Nil(t, err)
	assert.NotNil(t, req)

	req.Header.Set("Authorization", token)

	rr := httptest.NewRecorder()
	assert.NotNil(t, rr)
	handler := http.HandlerFunc(apiHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Profile:
	// 	alg:RS256
	// 	aud:https://preprocessor-2gtouos2pq-uc.a.run.app/v1/push
	// 	azp:109968801308363616831
	// 	email:preprocessor-sa@cloudylabs.iam.gserviceaccount.com
	// 	email_verified:true
	// 	exp:1565232609
	// 	iat:1565229009
	// 	iss:https://accounts.google.com
	// 	kid:3494b1e786cdad092e423766bbe37f54ed87b22d
	// 	sub:109968801308363616831
	// 	typ:JWT

}
