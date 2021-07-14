/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"

	"makefiletool/parser"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "makefiletool",
	Short: "",
	Long:  "parse makefile",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("%#v\n", cmd.Flag("file"))

		if fileFlag := cmd.Flag("file"); fileFlag != nil {
			fileName := fileFlag.Value.String()
			if fileName != "" {
				parser.SetFile(cmd.Flag("file").Value.String())
			}
		}

		targets, _ := parser.Parse()
		for id, target := range targets {
			fmt.Printf("%d. %v\n", id+1, target.Name)
			// fmt.Printf("deps: %v\n", len(target.Deps))
			for _, dep := range target.Deps {
				fmt.Printf("\t%v\n", dep)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.makefiletool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().StringP("file", "f", "", "file name default Makefile")
}
