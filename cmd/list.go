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

	"github.com/spf13/cobra"
	"github.com/tcassaert/deckctl/lib"
	"github.com/tidwall/gjson"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items",
	Long: `List items in your NextCloud Deck app.

These items can be boards, stacks and cards.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Possible items to list:")
		fmt.Println("\n  Boards\n  Stacks\n  Cards")
	},
}

var listBoardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "List boards",
	Run: func(cmd *cobra.Command, args []string) {
		boards := &lib.Board{}
		c := NewHTTPClient()
		boardlist := boards.Fetch(c)
		fmt.Printf("\nYour boards are:\n\n")
		for i := 0; i < len(boardlist); i++ {
			fmt.Printf(" %s\n", boardlist[i].Title)
		}
	},
}

var listStacksCmd = &cobra.Command{
	Use:   "stacks",
	Short: "List stacks",
	Run: func(cmd *cobra.Command, args []string) {
		stacks := &lib.Stack{}
		c := NewHTTPClient()
		board, _ := cmd.Flags().GetString("board")
		stacklist := stacks.Fetch(c, board)
		fmt.Printf("\nYour stacks on board %s are:\n\n", board)
		for i := 0; i < len(stacklist); i++ {
			fmt.Printf(" %s\n", stacklist[i].Title)
		}
	},
}

var listCardsCmd = &cobra.Command{
	Use:   "cards",
	Short: "List cards",
	Run: func(cmd *cobra.Command, args []string) {
		cards := &lib.Card{}
		c := NewHTTPClient()
		board, _ := cmd.Flags().GetString("board")
		stack, _ := cmd.Flags().GetString("stack")
		cardlist := cards.Fetch(c, board, stack)
		fmt.Printf("\nYour cards on board %s, stack %s are:\n\n", board, stack)
		for _, card := range cardlist {
			cardstring := card.String()
			cardtitle := gjson.Get(cardstring, "title")
			fmt.Printf("  %s\n", cardtitle)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listBoardsCmd)
	listCmd.AddCommand(listStacksCmd)
	listCmd.AddCommand(listCardsCmd)
	listStacksCmd.Flags().StringP("board", "b", "", "The title of the board")
	listCardsCmd.Flags().StringP("board", "b", "", "The title of the board")
	listCardsCmd.Flags().StringP("stack", "s", "", "The title of the stack")
}
