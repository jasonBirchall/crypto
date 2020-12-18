package test

import (
	"log"
	"os/exec"
	"regexp"
	"testing"
)

const mainGo = "../main.go"

func TestCheckCoins(t *testing.T) {

	out, err := exec.Command("go", "run", mainGo, "track").Output()
	if err != nil {
		log.Fatal(err)
	}

	actual := string(out)

	matched, err := regexp.MatchString(`Bitcoin`, actual)
	if err != nil {
		log.Fatal(err)
	}

	if !matched {
		t.Error("Expected output does not match actual. Actual:", actual)
	}
}
