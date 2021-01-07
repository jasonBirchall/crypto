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
	DataPoint0 float64 `njson:"data.coins.0.history.18"`
	DataPoint1 float64 `njson:"data.coins.0.history.19"`
	DataPoint2 float64 `njson:"data.coins.0.history.20"`
	DataPoint3 float64 `njson:"data.coins.0.history.21"`
	DataPoint4 float64 `njson:"data.coins.0.history.22"`
	DataPoint5 float64 `njson:"data.coins.0.history.23"`
	DataPoint6 float64 `njson:"data.coins.0.history.24"`
	DataPoint7 float64 `njson:"data.coins.0.history.25"`
	DataPoint8 float64 `njson:"data.coins.0.history.26"`
}

var coin string

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Graphs the recent price points of a certain coin",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := getData(coin)
		if err != nil {
			return err
		}
		graph := asciigraph.Plot(data)

		fmt.Println(graph)

		return nil
	},
}

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
