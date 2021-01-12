package cmd

import (
	"log"
	"os/exec"
	"regexp"
	"testing"
)

const mainGo = "../main.go"

func TestExecute(t *testing.T) {

	out, err := exec.Command("go", "run", mainGo, "track", "-c", "btc,eth").Output()
	if err != nil {
		log.Fatal(err)
	}

	actual := string(out)

	matched, err := regexp.MatchString(`£`, actual)
	if err != nil {
		log.Fatal(err)
	}

	if !matched {
		t.Error("Expected output does not match actual. Actual:", actual)
	}
}

func TestCheckCoins(t *testing.T) {
	coin := "btc"
	check, err := checkCoins(coin)
	if err != nil {
		t.Error("Unable to query API, test failed", err)
	}

	if check == "" {
		t.Error("Unable to query API, test failed")
	}
}
