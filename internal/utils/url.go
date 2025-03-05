package utils

import "regexp"

const ProductionURL = "https://api.safaricom.et"
const SandboxURL = "https://apisandbox.safaricom.et"

func ConstructURL(env string, endpoint string) string {
	finUrl := baseUrl(env) + endpoint
	return finUrl
}

func baseUrl(env string) string {
	if env == "PRODUCTION" {
		return ProductionURL
	}

	return SandboxURL
}

func MaskEndpoint(endpoint string) string {
    re := regexp.MustCompile(`apikey=[A-Za-z0-9]+`)
	return re.ReplaceAllString(endpoint, "apikey=*****************")
}
