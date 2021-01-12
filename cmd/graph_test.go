package cmd

import (
	"testing"
)

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
