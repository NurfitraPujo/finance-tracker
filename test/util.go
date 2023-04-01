package test

import (
	"os"
	"testing"
)

func SkipCI(t *testing.T) {
	if os.Getenv("IS_CI") == "true" {
		t.Skip("Skipping testing in CI environment")
	}
}
