package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/coleYab/mpesasdk/config"
	"github.com/coleYab/mpesasdk/internal/auth"
	"github.com/coleYab/mpesasdk/internal/utils"
)

// HttpClient: this is a wrapper over the http client that will
// provide simple http client functionalities for this sdk
type HttpClient struct {
	maxRetries uint
	maxConn    uint
	timeout    uint
	client     *http.Client
	token      *auth.AuthToken
}

// New: constructs the HttpClient
func New(cfg *config.Config) *HttpClient {
	transport := &http.Transport{
		MaxIdleConns:        int(cfg.MaxConcurrentConn),
		MaxIdleConnsPerHost: int(cfg.MaxConcurrentConn),
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(cfg.Timeout) * time.Second,
	}

    // Authorization token that will be used by the application
	token := auth.New(cfg.ConsumerKey, cfg.ConsumerSecret)

	return &HttpClient{
		maxRetries: cfg.MaxRetries,
		maxConn:    cfg.MaxConcurrentConn,
		timeout:    cfg.Timeout,
		client:     client,
		token:      token,
	}
}

// TODO: what if i return []bytes insted of http.Response
func (c *HttpClient) ApiRequest(env string, endpoint, method string, payload interface{}, authType string) (*http.Response, error) {
    url := utils.ConstructURL(env, endpoint)

	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonData)
	}

	var res *http.Response
	var err error

	for attempt := uint(0); attempt <= c.maxRetries; attempt++ {
		res, err = c.makeRequest(url, method, body, authType, env)
		if err == nil || !isTimeoutError(err) {
			break
		}

		// exponential backoff to avoid server overloads
		time.Sleep(time.Duration(attempt+1) * time.Second)
	}

	return res, err
}

// makeRequest: a simple helper function to make request and return the response
func (c *HttpClient) makeRequest(url, method string, body io.Reader, authType string, env string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	switch authType {
	case auth.AuthTypeBearer:
		authToken, err := c.token.GetToken(env)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", authToken)
	case auth.AuthTypeBasic:
		req.SetBasicAuth(c.token.GetUserCredentials())
	}

	return c.client.Do(req)
}

// A simple helper function that will check if the error is either
// Timeout error or context DeadlineExceeded error so that the client
// will retry on those occasions.
func isTimeoutError(err error) bool {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}
	return errors.Is(err, context.DeadlineExceeded)
}
