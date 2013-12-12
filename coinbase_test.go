package coinbase

import (
	"testing"
)

func TestapiKey(t *testing.T) {
	if apiKey() == "" {
		t.Fatal("Don't have an API key")
	}
}

func TestRequest_good(t *testing.T) {
	button := GetButton(&ButtonRequest{
		Name: "Abhi Test1",
	})

	if button == nil {
		t.Fatal("Button is nil")
	}
}
