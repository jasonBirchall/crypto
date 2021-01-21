package cmd

import (
	"testing"
)

// TestTestGetData tests the getData function and confirms it returns a string
// greater than 0 chars. If not, the test will fail.
func TestGetData(t *testing.T) {
	coin := "btc"

	check, err := getData(coin)
	if err != nil {
		t.Error("Unable to get data from graph command", err)
	}

	if len(check) < 1 {
		t.Error("Expected value doesn't contain values")
	}
}
