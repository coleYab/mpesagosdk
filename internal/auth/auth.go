// Package auth provides a simple authentication mechanism for managing
// and refreshing API authorization tokens.
//
// The package allows the creation of an `AuthToken` object, which stores
// the credentials (consumer key and secret) and manages the fetching,
// caching, and refreshing of an access token. It ensures that the token
// is always valid by checking its expiration and automatically fetching
// a new one when needed.
//
// It supports the following authentication types:
//	- Bearer: Used for API token-based authentication.
//	- Basic: Used for Basic HTTP Authentication.
//
// Key Features:
//	- Fetches and caches the authorization token to minimize unnecessary
//	  network requests.
//	- Automatically refreshes the token before it expires, with a buffer
//	  to ensure that the token remains valid during usage.
//	- Provides methods to retrieve the current token and the user's
//	  authentication credentials.
//
// This package is typically used in scenarios where the application
// requires authentication with an API, ensuring the token remains valid
// during usage without manual intervention.
//
// Example usage:
//
//	authToken := auth.New("consumer_key", "consumer_secret")
//	token, err := authToken.GetToken("PRODUCTION")
//	if err != nil {
//	    log.Fatalf("Error fetching token: %v", err)
//	}
//
//	fmt.Println("Access Token:", token)
package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/coleYab/mpesagosdk/internal/utils"
)

const (
	// AuthTypeBearer is used when the authorization header needs a Bearer token.
	AuthTypeBearer = "Bearer"
	// AuthTypeNone is used when no authorization is required.
	AuthTypeNone = ""
	// AuthTypeBasic is used for HTTP Basic Authentication.
	AuthTypeBasic = "Basic"
)

// AuthToken stores authentication credentials and access token metadata.
// It manages fetching and refreshing the authorization token.
type AuthToken struct {
	consumerKey    string
	consumerSecret string
	createdAt      time.Time
	expiresAt      time.Time
	token          string
}

// New initializes and returns a new instance of AuthToken using the
// provided consumerKey and consumerSecret.
func New(consumerKey, consumerSecret string) *AuthToken {
	token := &AuthToken{
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
	}
	return token
}

// GetToken returns a valid authorization token. If the current token is expired
// or not yet fetched, it automatically fetches a new one from the API.
//
// It uses the environment string to determine which API endpoint to call.
func (a *AuthToken) GetToken(env string) (string, error) {
	if a.token != "" && time.Now().Before(a.expiresAt) {
		return a.token, nil
	}

	if err := a.fetchAuthToken(env); err != nil {
		return "", err
	}

	return a.token, nil
}

func (a *AuthToken) GetUserCredentials() (string, string) {
	return a.consumerKey, a.consumerSecret
}

// setAuthToken sets the token, token type, and expiration time.
// It subtracts 10 seconds from the actual expiry to ensure buffer time
// and avoid token expiration errors during API requests.
func (a *AuthToken) setAuthToken(tokenType, token string, expiresIn int) {
	expiryMargin := 10
	validFor := time.Duration(expiresIn-expiryMargin) * time.Second

	a.token = fmt.Sprintf("%v %v", tokenType, token)
	a.createdAt = time.Now()
	a.expiresAt = a.createdAt.Add(validFor)
}

// fetchAuthToken makes an HTTP request to the API to obtain a new token.
// It constructs the appropriate URL based on the environment, and uses Basic Auth
// for authentication. The token response is parsed and stored.
func (a *AuthToken) fetchAuthToken(env string) error {
	url := utils.ConstructURL(env, "/v1/token/generate?grant_type=client_credentials")
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("error: while creating auth request")
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(a.consumerKey, a.consumerSecret)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var authResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   string `json:"expires_in"`
		ResultCode  string `json:"resultCode"`
		ResultDesc  string `json:"resultDesc"`
	}

	if err := json.Unmarshal(body, &authResponse); err != nil {
		return err
	}

	if authResponse.ResultCode != "" {
		return fmt.Errorf("error occured due to: %v", authResponse.ResultDesc)
	}

	expiresIn, _ := strconv.Atoi(authResponse.ExpiresIn)
	a.setAuthToken(authResponse.TokenType, authResponse.AccessToken, expiresIn)

	return nil
}
