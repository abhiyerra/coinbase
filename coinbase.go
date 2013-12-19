package coinbase

import (
	"os"
)

var (
	CoinbaseApiKey         string
	CoinbaseCallbackSecret string
)

func apiKey() string {
	if CoinbaseApiKey == "" {
		CoinbaseApiKey = os.Getenv("COINBASE_API_KEY")
		if CoinbaseApiKey == "" {
			return ""
		}
	}

	return CoinbaseApiKey
}
