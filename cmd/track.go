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

	api "github.com/jasonbirchall/crypto-tracker/pkg/api"
	"github.com/m7shapan/njson"
	"github.com/spf13/cobra"
)

type Coin struct {
	Rate   string  `njson:"data.coins.0.price"`
	Change float64 `njson:"data.coins.0.change"`
}

var coinsArg []string

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Allows you to track the rise and fall of specific coins",

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, v := range coinsArg {
			price, err := checkCoins(v)
			if err != nil {
				return err
			}

			fmt.Print(strings.ToUpper(v)+" ", price)
		}

		return nil
	},
}

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

func grabPrice(body []byte) (string, error) {
	var c Coin

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
		return fmt.Sprintf("£%.2f | %.2f%%   ", v, c.Change), nil
	} else {
		return fmt.Sprintf("£%.2f | +%.2f%%   ", v, c.Change), nil
	}
}

func init() {
	rootCmd.AddCommand(trackCmd)
	trackCmd.Flags().StringSliceVarP(&coinsArg, "coin", "c", []string{}, "")
	trackCmd.MarkPersistentFlagRequired("coin")
}
