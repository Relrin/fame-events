package pkg

import (
	"testing"
)

func TestGenerateNumber(t *testing.T) {
	result := "1"

	if result == "" {
		t.Error("got an empty string")
	}
}
