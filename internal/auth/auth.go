package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/coleYab/mpesasdk/internal/utils"
)

const (
	AuthTypeBearer = "Bearer"
	AuthTypeNone   = ""
	AuthTypeBasic  = "Basic"
)

// AuthToken: a simple struct that will store the auth related data
// of the user for this application
type AuthToken struct {
	consumerKey    string
	consumerSecret string
	createdAt      time.Time
	expiresAt      time.Time
	token          string
}

func New(consumerKey, consumerSecret string) *AuthToken {
	token := &AuthToken{
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
	}
	return token
}

// GetToken: function that will get the authorization token.
// based on the enviroment that the application is currently running
func (a *AuthToken) GetToken(env string) (string, error) {
	if a.token != "" && time.Now().Before(a.expiresAt) {
		return a.token, nil
	}

	// otherwise we need to fetch the new token
	if err := a.fetchAuthToken(env); err != nil {
		return "", err
	}

	return a.token, nil
}

// TODO: do we really need this?
func (a *AuthToken) GetUserCredentials() (string, string) {
	return a.consumerKey, a.consumerSecret
}

func (a *AuthToken) setAuthToken(tokenType, token string, expiresIn int) {
	a.token = fmt.Sprintf("%v %v", tokenType, token)
	a.createdAt = time.Now()
	a.expiresAt = a.createdAt.Add(time.Duration(expiresIn-10) * time.Second)
	// I want to expire always before 10 seconds of the actual expiration time
	// that will ensure we will never have the authorization token expired error
}

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

	// TODO: use the sdk error handling
	if authResponse.ResultCode != "" {
		return fmt.Errorf("error occured due to: %v", authResponse.ResultDesc)
	}

	expiresIn, _ := strconv.Atoi(authResponse.ExpiresIn)
	a.setAuthToken(authResponse.TokenType, authResponse.AccessToken, expiresIn)

	return nil
}
