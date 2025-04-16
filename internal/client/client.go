// Package client provides a wrapper around the HTTP client to facilitate making
// API requests for the mpesagosdk.
// 
// This package provides an `HttpClient` struct that manages HTTP requests, including
// handling retries, authentication (using Bearer or Basic Auth), and configuring
// timeouts and maximum idle concurrent connections that your application has to keep.
// 
// Key Features:
//	- Handles HTTP requests with retries in case of timeouts.
//	- Supports multiple authentication schemes: Bearer and Basic.
//	- Provides a configurable client with timeout and maximum concurrent connections.
//	- Handles exponential backoff for retries to avoid server overload.
// 
// This package makes it easier to interact with an external API while managing
// important aspects of HTTP communication, such as retries, authentication, and connection limits.
// 
// Example usage:
//	cfg := &config.Config{
//	    ConsumerKey:      "consumer-key",
//	    ConsumerSecret:   "consumer-secret",
//	    MaxRetries:       3,
//	    MaxConcurrentConn: 10,
//	    Timeout:           30,
//	}
// 
//	client := client.New(cfg)
//	response, err := client.ApiRequest("PRODUCTION", "/v1/resource", "GET", nil, auth.AuthTypeBearer)
//
//	if err != nil {
//	    log.Fatalf("API request failed: %v", err)
//	}
//
//	defer response.Body.Close()
//	fmt.Println("Response:", response.Status)
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

	"github.com/coleYab/mpesagosdk/config"
	"github.com/coleYab/mpesagosdk/internal/auth"
	"github.com/coleYab/mpesagosdk/internal/utils"
)

// HttpClient is a wrapper over the standard HTTP client that manages retries, timeouts,
// and authentication when making API requests. It provides functionalities for
// sending requests with either Bearer or Basic authentication and supports retries
// with exponential backoff in case of timeouts.
type HttpClient struct {
	maxRetries int
	maxConn    int
	timeout    int
	client     *http.Client
	token      *auth.AuthToken
}

// New constructs a new HttpClient based on the provided configuration settings.
// It sets up the underlying HTTP client, including transport settings and token management.
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

// ApiRequest sends an HTTP request to the given endpoint, with the specified HTTP method
// (e.g., GET, POST) and payload. It automatically handles retries in case of timeouts
// and returns the HTTP response or an error. The function uses the provided `authType`
// to determine the authorization method (Bearer or Basic).
//
// 	- `env` specifies the environment (e.g., "PRODUTION" or "SANDBOX"), and `authType`
// specifies the authentication scheme to be used (either `AuthTypeBearer` or `AuthTypeBasic`).
//
//	-	Retries: retries are only done if the issue is timeout error or context DeadlineExceeded
//	 errors. We don't want to retiry other errors because it is useless in to retry in most
//	 of the other cases. We are using (Exponential backoff)[https://en.wikipedia.org/wiki/Exponential_backoff]
//
// Returns the HTTP response and an error, if any.
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

	// Retries
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		res, err = c.makeRequest(url, method, body, authType, env)
		if err == nil || !isTimeoutError(err) {
			break
		}

		// Exponential backoff: retry after increasing delay (not to pass the rate limit)
		time.Sleep(time.Duration(attempt+1) * time.Second)
	}

	return res, err
}

// makeRequest: sends the HTTP request with the given method, URL, body, and authentication.
// It returns the HTTP response or an error if something goes wrong.
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

// isTimeoutError checks if an error is due to a network timeout
// or a context deadline exceeded error. These errors trigger a retry.
func isTimeoutError(err error) bool {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}
	return errors.Is(err, context.DeadlineExceeded)
}
