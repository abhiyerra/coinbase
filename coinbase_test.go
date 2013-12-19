package coinbase

import (
	"testing"
)

func TestApiKey(t *testing.T) {
	if apiKey() == "" {
		t.Fatal("Don't have an API key")
	}
}
