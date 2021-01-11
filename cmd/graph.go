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
	"math"
	"reflect"

	"github.com/guptarohit/asciigraph"
	api "github.com/jasonbirchall/crypto-tracker/pkg/api"
	"github.com/m7shapan/njson"
	"github.com/spf13/cobra"
)

type Price struct {
	DataPoint5  float64 `njson:"data.coins.0.history.5"`
	DataPoint6  float64 `njson:"data.coins.0.history.6"`
	DataPoint7  float64 `njson:"data.coins.0.history.7"`
	DataPoint8  float64 `njson:"data.coins.0.history.8"`
	DataPoint9  float64 `njson:"data.coins.0.history.9"`
	DataPoint10 float64 `njson:"data.coins.0.history.10"`
	DataPoint11 float64 `njson:"data.coins.0.history.11"`
	DataPoint12 float64 `njson:"data.coins.0.history.12"`
	DataPoint13 float64 `njson:"data.coins.0.history.13"`
	DataPoint14 float64 `njson:"data.coins.0.history.14"`
	DataPoint15 float64 `njson:"data.coins.0.history.15"`
	DataPoint16 float64 `njson:"data.coins.0.history.16"`
	DataPoint17 float64 `njson:"data.coins.0.history.17"`
	DataPoint18 float64 `njson:"data.coins.0.history.18"`
	DataPoint19 float64 `njson:"data.coins.0.history.19"`
	DataPoint20 float64 `njson:"data.coins.0.history.20"`
	DataPoint21 float64 `njson:"data.coins.0.history.21"`
	DataPoint22 float64 `njson:"data.coins.0.history.22"`
	DataPoint23 float64 `njson:"data.coins.0.history.23"`
	DataPoint24 float64 `njson:"data.coins.0.history.24"`
	DataPoint25 float64 `njson:"data.coins.0.history.25"`
	DataPoint26 float64 `njson:"data.coins.0.history.26"`
}

var coin string

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

		graph := asciigraph.Plot(data, asciigraph.Height(10))

		fmt.Println(graph)

		return nil
	},
}

// getCoin takes a string as an argument, this is passed as a flag by the
// graph command. It queries the API package and collates a collection of
// float64 data type. This collection is then returned back to the command
// function.
func getData(coin string) ([]float64, error) {
	var c Price
	var arr []float64

	data, err := api.Query(coin)
	if err != nil {
		return arr, err
	}

	err = njson.Unmarshal([]byte(data), &c)
	if err != nil {
		return arr, err
	}

	fields := reflect.TypeOf(c)
	values := reflect.ValueOf(c)
	num := fields.NumField()

	for i := 0; i < num; i++ {
		value := values.Field(i)
		v := value.Interface().(float64)
		c := math.Round(v*100) / 100

		arr = append(arr, c)
	}

	return arr, nil
}

func init() {
	rootCmd.AddCommand(graphCmd)

	graphCmd.Flags().StringVarP(&coin, "coin", "c", "", "Coin to place in graph")
	graphCmd.MarkFlagRequired("coin")
}
