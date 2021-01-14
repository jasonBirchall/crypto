/*
Copyright © 2021 json.birchall@gmail.com

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
	"encoding/json"
	"fmt"

	api "github.com/jasonbirchall/crypto-tracker/internal/api"
	"github.com/spf13/cobra"
)

type Price struct {
	DataPoints []string `json:"datapoints"`
}

var coin string
var height int

// graphCmd represents the graph command. You can only return an error statement
// using RunE.
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Graphs the recent price points of a certain coin",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := getData(coin)
		if err != nil {
			return err
		}

		// graph := asciigraph.Plot(data, asciigraph.Height(height))

		fmt.Println(data)

		return nil
	},
}

// getData takes a string as an argument, this is passed as a flag by the
// graph command. It queries the API package and collates a collection of
// float64 data type. This collection is then returned back to the command
// function.
func getData(coin string) ([]string, error) {
	var p Price
	var arr []string

	data, err := api.Query(coin)

	err = json.Unmarshal(data, &p)
	if err != nil {
		return arr, err
	}

	for _, v := range p.DataPoints {
		fmt.Println(v)
	}

	return arr, nil
}

func (p *Price) UnmarshalJSON(data []byte) error {
	var v struct {
		Data struct {
			Coins []struct {
				History []string `json:"history"`
			} `json:"coins"`
		} `json:"data"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if len(v.Data.Coins) > 0 {
		p.DataPoints = v.Data.Coins[0].History
	}

	return nil
}

func init() {
	rootCmd.AddCommand(graphCmd)

	graphCmd.Flags().StringVarP(&coin, "coin", "c", "", "Coin to place in graph")
	graphCmd.Flags().IntVarP(&height, "height", "H", 10, "Height of the graph")
	graphCmd.MarkFlagRequired("coin")
}
