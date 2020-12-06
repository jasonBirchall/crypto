/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"strings"

	"github.com/spf13/cobra"
)

// echoCmd represents the echo command
var cmdEcho = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, " "))
	},
}

// var echoCmd = &cobra.Command{
// 	Use:   "echo [string to echo]",
// 	Short: "Echo anything to the screen.",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Args: cobra.MaximumNArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("Echo: " + strings.Join(args, " "))
// 	},
// }

func init() {
	rootCmd.AddCommand(echoCmd)
}
