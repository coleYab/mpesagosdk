package auth

import (
	"testing"
)

func TestGetToken_InvalidCredentialsSandbox(t *testing.T) {
	token := New("invalidConsumerKey", "invalidConsumerSecret")

	env := "SANDBOX"

	_, err := token.GetToken(env)
	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
}

func TestGetToken_InvalidCredentials(t *testing.T) {
	token := New("invalidConsumerKey", "invalidConsumerSecret")
	env := "PRODUCTION"
	_, err := token.GetToken(env)
	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
}
