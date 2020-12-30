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
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/m7shapan/njson"
	"github.com/spf13/cobra"
)

type Coin struct {
	Rate   string  `njson:"data.coins.0.price"`
	Change float64 `njson:"data.coins.0.change"`
}

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Allows you to track the rise and fall of specific coins",

	RunE: func(cmd *cobra.Command, args []string) error {
		coinName, err := cmd.Flags().GetString("coin")
		if err != nil {
			return err
		}

		price, change, err := checkCoins(coinName)
		if err != nil {
			return err
		}

		fmt.Println("£", price, "|", change, "%")
		return nil
	},
}

func checkCoins(c string) (float64, float64, error) {
	data, err := queryApi(c)
	if err != nil {
		return -1, -1, err
	}

	price, change, err := grabPrice(data)
	if err != nil {
		return -1, -1, err
	}

	return price, change, nil
}

func queryApi(c string) ([]byte, error) {
	apiUrl := "https://api.coinranking.com/v1/public/coins?base=gbp&prefix=" + c
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "crypto-tracker")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}

func grabPrice(body []byte) (float64, float64, error) {
	var c Coin

	err := njson.Unmarshal([]byte(body), &c)
	if err != nil {
		log.Fatal(err)
	}

	// Convert string to float64
	v, err := strconv.ParseFloat(c.Rate, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Round the pennies to the nearest
	r := math.Round(v*100) / 100

	return r, c.Change, nil
}

func init() {
	rootCmd.AddCommand(trackCmd)
	trackCmd.Flags().StringP("coin", "c", "bitcoin", "You must specify a coin (required)")
	trackCmd.MarkPersistentFlagRequired("coin")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
