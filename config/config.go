// Package config provides functionality to manage and load configuration settings for the SDK.
// It allows for both default configuration settings and environment-based configuration loading.
//
// The package supports the following features:
//	- Creates a default configuration via the `New` function.
//	- Loads configuration from environment variables via the `NewFromEnv` function.
//	- Provides functions to read environment variables safely with fallback values for various types (string, int).
//	- Validates important configuration fields like the consumer key and secret.
//
// Example usage:
//
//	// Loading configuration using default values
//	config := config.New("your-consumer-secret", "your-consumer-key", "DEBUG")
//	fmt.Println(config.MaxRetries)
//
//	// Loading configuration from environment variables
//	config, err := config.NewFromEnv()
//	if err != nil {
//	    log.Fatalf("Error loading config: %v", err)
//	}
//	fmt.Println(config.Timeout)
//
// Environment Variables:
// The `NewFromEnv` function will automatically load the following environment variables:
//	- `MAX_CONCURRENT_CONN`: The maximum number of concurrent connections (default: 1000).
//	- `MAX_RETRIES`: The maximum number of retries for failed requests (default: 3).
//	- `TIMEOUT`: The timeout duration for requests in seconds (default: 5).
//	- `CONSUMER_SECRET`: The consumer secret for authentication (must be set).
//	- `CONSUMER_KEY`: The consumer key for authentication (must be set).
//	- `LOG_LEVEL`: The logging level (default: "DEBUG").
//	- `ENVIROMENT`: The environment ("SANDBOX" or "PRODUCTION", default: "SANDBOX").
package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the configuration values for the SDK.
// It includes fields for authentication, retries, timeouts,
// logging, and environment settings.
type Config struct {
	// Maximum number of concurrent connections
	MaxConcurrentConn int
	// Maximum number of retry attempts for failed requests
	MaxRetries int
	// Timeout duration for requests in seconds
	Timeout int
	// Consumer secret for authentication
	ConsumerSecret string
	// Consumer key for authentication
	ConsumerKey string
	// Logging level (e.g., "DEBUG", "INFO", "ERROR")
	LogLevel string
	// Environment ("PRODUCTION" or "SANDBOX")
	Enviroment string
}

// New creates a new configuration instance with the provided consumer key, secret, and log level.
// It uses default values for the other configuration parameters:
//	- Timeout: 5 seconds
//	- MaxRetries: 3
//	- MaxConcurrentConn: 1000
//
// Parameters:
//	- consumerSecret: The consumer secret for API authentication.
//	- consumerKey: The consumer key for API authentication.
//	- logLevel: The logging level (e.g., "DEBUG", "INFO").
//
// Returns:
//	- A pointer to the Config instance with the provided and default values.
func New(consumerSecret, consumerKey string, logLevel string) *Config {
	return &Config{
		ConsumerSecret:    consumerSecret,
		ConsumerKey:       consumerKey,
		LogLevel:          logLevel,
		Timeout:           5,
		MaxRetries:        3,
		MaxConcurrentConn: 1000,
	}
}

// getEnv is a helper function that retrieves an environment variable's value
// or returns the provided fallback value if the environment variable is not set.
//
// Parameters:
//	- key: The environment variable key.
//	- fallback: The fallback value to return if the environment variable is not found.
//
// Returns:
//	- The value of the environment variable or the fallback value if not found.
func getEnv(key string, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

// getEnvInt is a helper function that retrieves an environment variable's value,
// converts it to an integer, and returns it. If the conversion fails or the variable
// is not set, it returns the provided fallback value.
//
// Parameters:
//	- key: The environment variable key.
//	- fallback: The fallback integer value to return if the environment variable is not found or invalid.
//
// Returns:
//	- The integer value of the environment variable or the fallback value if not found or invalid.
func getEnvInt(key string, fallback int) int {
	if v, ok := os.LookupEnv(key); ok {
		res, err := strconv.Atoi(v)
		if err == nil { // is it valid other wise return fallback
			return res
		}
	}
	return fallback
}

// NewFromEnv creates a new configuration instance by loading values from environment variables.
// It supports configuration of concurrent connections, retries, timeouts, authentication keys,
// log level, and environment. It validates the required values (consumer key and secret) and
// returns an error if any required values are missing or invalid. This function will use the
// this (autoload)[https://github.com/joho/godotenv/tree/main/autoload] package that is provided by (joho)[https://github.com/joho].
//
// Returns:
//	- A pointer to the Config instance with values loaded from environment variables.
//	- An error if any required configuration values are missing or invalid.
//
// Example usage:
//
//	config, err := config.NewFromEnv()
//	if err != nil {
//	    log.Fatalf("Error loading config: %v", err)
//	}
//	fmt.Println(config.ConsumerKey)
func NewFromEnv() (*Config, error) {
	config := &Config{
		MaxConcurrentConn: getEnvInt("MAX_CONCURRENT_CONN", 1000),
		MaxRetries:        getEnvInt("MAX_RETRIES", 3),
		Timeout:           getEnvInt("TIMEOUT", 5),
		ConsumerSecret:    getEnv("CONSUMER_SECRET", ""),
		ConsumerKey:       getEnv("CONSUMER_KEY", ""),
		LogLevel:          getEnv("LOG_LEVEL", "DEBUG"),
		Enviroment:        getEnv("ENVIROMENT", "SANDBOX"),
	}

	if config.ConsumerKey == "" || config.ConsumerSecret == "" {
		return nil, fmt.Errorf("consumer Secret or consumer key is requried")
	}

	if config.Timeout == 0 {
		return nil, fmt.Errorf("timout has to be greater than 0")
	}

	return config, nil
}
