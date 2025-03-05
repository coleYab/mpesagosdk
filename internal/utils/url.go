package utils

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
