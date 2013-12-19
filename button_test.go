package coinbase

import (
	"testing"
)

func TestGetButton(t *testing.T) {
	button := GetButton(&ButtonRequest{
		Name: "Abhi Test1",
	})

	if button == nil {
		t.Fatal("Button is nil")
	}
}
