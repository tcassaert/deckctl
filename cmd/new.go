/*
Copyright © 2019 Thomas Cassaert <tcassaert@inuits.eu>

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
	"github.com/tcassaert/deckctl/lib"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create items",
	Long: `Create items in your NextCloud Deck app.

These items can be boards, stacks and cards.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Possible items to create:")
		fmt.Println("\n  Boards\n  Stacks\n  Cards")
	},
}

var newBoardCmd = &cobra.Command{
	Use:   "board",
	Short: "Create new board",
	Run: func(cmd *cobra.Command, args []string) {
		board := &lib.Board{}
		c := NewHTTPClient()
		title, _ := cmd.Flags().GetString("title")
		color, _ := cmd.Flags().GetString("color")
		err := board.New(c, title, color)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Created %s board\n", title)
		}
	},
}

var newStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Create new stack",
	Run: func(cmd *cobra.Command, args []string) {
		stack := &lib.Stack{}
		c := NewHTTPClient()
		board, _ := cmd.Flags().GetString("board")
		order, _ := cmd.Flags().GetInt("order")
		title, _ := cmd.Flags().GetString("title")
		err := stack.New(c, board, title, order)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Created %s stack on board %s\n", title, board)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.AddCommand(newBoardCmd)
	newCmd.AddCommand(newStackCmd)
	newBoardCmd.Flags().StringP("title", "t", "", "The title of the board")
	newBoardCmd.Flags().StringP("color", "c", "", "The color of the board")
	newStackCmd.Flags().StringP("board", "b", "", "The title of the board")
	newStackCmd.Flags().IntP("order", "o", 0, "Order of the stack")
	newStackCmd.Flags().StringP("title", "t", "", "The title of the stack")
}
