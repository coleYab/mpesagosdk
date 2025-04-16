// Package utils provides utility functions that assist with constructing URLs,
// selecting the appropriate API base URL, and masking sensitive information
// within endpoints.
// 
// It supports determining the correct base URL based on the environment (production or sandbox),
// and masking API keys or other sensitive information from endpoint strings to ensure security.
// 
// Functions:
//	- ConstructURL: Builds a complete URL for an API request by combining the base URL
//	  with the specified endpoint.
//	- MaskEndpoint: Masks the API key in the endpoint string to prevent exposing sensitive
//	  information.
//	- baseUrl: Returns the correct base URL (either production or sandbox) based on the environment.
// 
// Example usage:
//	url := utils.ConstructURL("PRODUCTION", "/v1/someendpoint")
//	fmt.Println(url)
// Output: https://api.safaricom.et/v1/someendpoint
// 
//	maskedEndpoint := utils.MaskEndpoint("/v1/someendpoint?apikey=12345abcd")
//	fmt.Println(maskedEndpoint)
//
// Output: /v1/someendpoint?apikey=*****************
package utils

import "regexp"

const ProductionURL = "https://api.safaricom.et"
const SandboxURL = "https://apisandbox.safaricom.et"

// ConstructURL builds and returns the complete URL by combining the base URL
// (based on the environment) and the provided endpoint.
//
// Parameters:
//	- env: The environment, either "PRODUCTION" or any other value for sandbox.
//	- endpoint: The specific endpoint to be appended to the base URL.
//
// Returns:
//	- The full constructed URL as a string.
// 	
// Example usage:
//
//	url := ConstructURL("PRODUCTION", "/v1/endpoint")
//	fmt.Println(url)
//
// Output: https://api.safaricom.et/v1/endpoint
func ConstructURL(env string, endpoint string) string {
	finUrl := baseUrl(env) + endpoint
	return finUrl
}

// baseUrl: determines and returns the correct base URL based on the environment.
//
// Parameters:
//	- env: The environment, typically "PRODUCTION" for production or anything else for sandbox.
//
// Returns:
//	- A string representing the base URL for the given environment.
//
// Example usage:
//
//	url := baseUrl("PRODUCTION")
//	fmt.Println(url)
//
// Output: https://api.safaricom.et
func baseUrl(env string) string {
	if env == "PRODUCTION" {
		return ProductionURL
	}

	return SandboxURL
}

// MaskEndpoint takes an endpoint string and replaces the API key with a masked version.
//
// This function is useful for preventing the exposure of sensitive data like API keys
// in logs or error messages.
//
// Parameters:
//	- endpoint: The endpoint string potentially containing an API key.
//
// Returns:
//	- A string with the API key masked with asterisks.
//
// Example usage:
//
//	masked := MaskEndpoint("/v1/endpoint?apikey=12345abcd")
//	fmt.Println(masked)
//
// Output: /v1/endpoint?apikey=*****************
func MaskEndpoint(endpoint string) string {
	re := regexp.MustCompile(`apikey=[A-Za-z0-9]+`)
	return re.ReplaceAllString(endpoint, "apikey=*****************")
}
