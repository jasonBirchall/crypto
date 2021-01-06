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
	"fmt"

	api "github.com/jasonbirchall/crypto-tracker/pkg/api"
	"github.com/m7shapan/njson"
	"github.com/spf13/cobra"
)

type Price struct {
	History string `njson:"data.coins.0.history"`
}

var coin string

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Graphs the recent price points of a certain coin",
	Run: func(cmd *cobra.Command, args []string) {
		// data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 5, 6}
		data, err := getData(coin)
		// graph := asciigraph.Plot(data)
		fmt.Println(data, err)

		// fmt.Println(graph)
	},
}

func getData(coin string) (float64, error) {
	var c Price

	data, err := api.Query(coin)
	if err != nil {
		return -1, err
	}

	err = njson.Unmarshal([]byte(data), &c)
	if err != nil {
		return -1, err
	}

	fmt.Println(c.History)
	for i, v := range c.History {
		fmt.Println(i, v)
	}

	return -1, nil

}

func init() {
	rootCmd.AddCommand(graphCmd)

	graphCmd.Flags().StringVarP(&coin, "coin", "c", "", "Coin to place in graph")
	graphCmd.MarkFlagRequired("coin")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
