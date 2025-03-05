package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config: a simple struct that will be used as a config
// setup over all the functions of this sdk.
type Config struct {
	MaxConcurrentConn int
	MaxRetries        int
	Timeout           int
	ConsumerSecret    string
	ConsumerKey       string
	LogLevel          string
	Enviroment        string
}

// getEnv: this function is usefull to get the
// key from the enviroment variables if it is set
// otherwise it will return the fallback value that
// is provided
func getEnv(key string, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

// getEnvUint: this function is usefull to get the
// key from the enviroment variables if it is set
// otherwise it will return the fallback value that
// is provided and it will convert the value of the
// variables to integer
func getEnvInt(key string, fallback int) int {
	if v, ok := os.LookupEnv(key); ok {
		res, err := strconv.Atoi(v)
		if err == nil { // is it valid other wise return fallback
			return res
		}
	}
	return fallback
}

// NewFromEnv: this function is the function that is used to load
// the config data from the user enviroment variables it uses the
// autoload package provided by () to autoload the enviroment variable
// this will let them to the users.
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

	// TODO: take this out as validation error
	if config.ConsumerKey == "" || config.ConsumerSecret == "" {
		return nil, fmt.Errorf("consumer Secret or consumer key is requried")
	}

	// TODO: take this as a validation error
	if config.Timeout == 0 {
		return nil, fmt.Errorf("timout has to be greater than 0")
	}

	return config, nil
}
