package test

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os/exec"
	"testing"
)

const mainGo = "../../main.go"

func TestCheckCoins(t *testing.T) {

	out, err := exec.Command("go", "run", mainGo, "track").Output()
	if err != nil {
		log.Fatal(err)
	}

	// actual, _ := strconv.ParseFloat(out, 64)
	bits := binary.LittleEndian.Uint64(out)
	actual := math.Float64bits(bits)
	fmt.Printf("%T", actual)

	if actual < "" {
		t.Error("Expected output does not match actual. Actual:", actual)
	}
}
