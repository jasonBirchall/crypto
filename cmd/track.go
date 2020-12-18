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
	"strings"
	"time"

	"github.com/m7shapan/njson"
	"github.com/spf13/cobra"
)

type Coin struct {
	Rate string `njson:"bpi.GBP.rate"`
}

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Allows you to track the rise and fall of specific coins",

	// Args: cobra.MinimumNArgs(1),
	Run: checkCoins,
}

func checkCoins(cmd *cobra.Command, args []string) {
	data, err := queryApi()
	if err != nil {
		log.Fatalln(err)
	}

	price, err := grabPrice(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bitcoin:", price)
}

func queryApi() ([]byte, error) {
	url := "https://api.coindesk.com/v1/bpi/currentprice.json"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
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

func grabPrice(body []byte) (float64, error) {
	var c Coin

	err := njson.Unmarshal([]byte(body), &c)
	if err != nil {
		log.Fatal(err)
	}

	// Remove all commas from string
	res1 := strings.ReplaceAll(c.Rate, ",", "")

	// Convert string to float64
	v, err := strconv.ParseFloat(res1, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Round the pennies up
	r := math.Round(v)

	return r, nil
}

func init() {
	rootCmd.AddCommand(trackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
