/*
Copyright © 2020 Thomas Cassaert <tcassaert@inuits.eu>

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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tcassaert/deckctl/lib"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a board based on config",
	Long:  `Initialize a board based on a predefined layout in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		board := &lib.Board{}
		stack := &lib.Stack{}
		c := NewHTTPClient()
		title, _ := cmd.Flags().GetString("title")
		color, _ := cmd.Flags().GetString("color")
		errBoard := board.New(c, title, color)
		if errBoard != nil {
			log.Fatal(errBoard)
		}
		stacklist := viper.GetStringSlice("init.stacks")
		for _, stackname := range stacklist {
			errStack := stack.New(c, title, stackname, 0)
			if errStack != nil {
				log.Fatal(errStack)
			}
		}
		fmt.Printf("Board \"%s\" created with the configured stacks.\n", title)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("title", "t", "", "The title of the board")
	initCmd.Flags().StringP("color", "c", "", "The color of the board")
}
