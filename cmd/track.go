/*
Copyright © 2020 json.birchall@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	tm "github.com/buger/goterm"
	"github.com/fatih/color"
	api "github.com/jasonbirchall/crypto/pkg/api"
	"github.com/m7shapan/njson"
	"github.com/spf13/cobra"
)

// Coin defines the current price rate of a coin and subsequently
// a change in rate using a float value.
type Coin struct {
	Rate   string  `njson:"data.coins.0.price"`
	Change float64 `njson:"data.coins.0.change"`
}

// coinsArg is gathered using the --coin or -c flag.
var coinsArg []string

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Allows you to track the rise and fall of specific coins",

	RunE: func(cmd *cobra.Command, args []string) error {
		if watch {
			// Clear the current screen.
			tm.Clear()
			// Loop over track command, executing track command every second.
			for {
				err := loopTrack()
				if err != nil {
					return err
				}
			}
		} else {
			p, err := execute()
			if err != nil {
				return err
			}
			fmt.Println(p)
		}
		return nil
	},
}

// execute forms a map and loops its values to create a string to print.
// An example of this would be BTC £12322 | +4.32%
func execute() (string, error) {
	var s string
	m, err := createMap()
	if err != nil {
		return "An error occured creating a map:", err
	}

	for k, v := range m {
		s = s + strings.ToUpper(k) + " " + v
	}

	return s, nil
}

// loopTrack is used to loop over the execute function every second. This function is
// called if the --watch flag is set.
func loopTrack() error {
	tm.MoveCursor(1, 1)
	tm.Println("Current Time:", time.Now().Format(time.RFC1123), "\n-------------")
	p, err := execute()
	if err != nil {
		return err
	}
	tm.Println(p)
	// Call it every time at the end of rendering
	tm.Flush()
	time.Sleep(time.Second)

	return nil
}

// createMap takes no arguments and calls checkCoins to query an API.
// It returns a map of coin(string): price(string).
func createMap() (map[string]string, error) {
	m := make(map[string]string)
	for _, c := range coinsArg {
		price, err := checkCoins(c)
		if err != nil {
			return m, err
		}
		m[c] = price
	}
	return m, nil

}

// checkCoins takes a coin shorthand as a string, i.e. btc and queries the
// api package to retrieve a collection of bytes. It then calls grabPrice and
// creates a Coin property. Finally, checkCoins will return a string to the main
// Cobra command.
func checkCoins(c string) (string, error) {
	data, err := api.Query(c)
	if err != nil {
		return "An error has occured queying the API:", err
	}

	price, err := grabPrice(data)
	if err != nil {
		return "An error has occurred grabbing the json object:", err
	}

	return price, nil
}

// grabPrice accepts a slice of bytes from the checkCoins function and unmarshalls it into
// a coin object. This object is then converted to a float to show only two decimal places and
// then returns the object values depending on if they're positive or negative.
func grabPrice(body []byte) (string, error) {
	var c Coin
	pos := color.New(color.FgGreen)
	neg := color.New(color.FgRed)

	err := njson.Unmarshal([]byte(body), &c)
	if err != nil {
		log.Fatal(err)
	}

	// Convert string to float64 to show two decimal places only.
	v, err := strconv.ParseFloat(c.Rate, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Check to see if the difference is positive or negative. If
	// positive then add a + symbol.
	isNeg := math.Signbit(c.Change)
	if isNeg {
		return neg.Sprintf("£%.2f | %.2f%%   ", v, c.Change), nil
	} else {
		return pos.Sprintf("£%.2f | +%.2f%%   ", v, c.Change), nil
	}
}

func init() {
	rootCmd.AddCommand(trackCmd)
	trackCmd.Flags().StringSliceVarP(&coinsArg, "coin", "c", []string{}, "")
	trackCmd.MarkPersistentFlagRequired("coin")
}
