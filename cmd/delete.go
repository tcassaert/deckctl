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
	"github.com/tcassaert/deckctl/lib"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete items",
	Long: `Delete items in your NextCloud Deck app.

These items can be boards, stacks and cards.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Possible items to delete:")
		fmt.Println("\n  Boards\n  Stacks\n  Cards")
	},
}

var deleteBoardCmd = &cobra.Command{
	Use:   "board",
	Short: "Delete board",
	Run: func(cmd *cobra.Command, args []string) {
		boards := &lib.Board{}
		c := NewHTTPClient()
		title, _ := cmd.Flags().GetString("title")
		err := boards.Delete(c, title)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Deleted %s board\n", title)
		}
	},
}

var deleteStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Delete stack",
	Run: func(cmd *cobra.Command, args []string) {
		stacks := &lib.Stack{}
		c := NewHTTPClient()
		board, _ := cmd.Flags().GetString("board")
		title, _ := cmd.Flags().GetString("stack")
		err := stacks.Delete(c, board, title)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Deleted \"%s\" stack on board %s\n", title, board)
		}
	},
}

var deleteCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Delete card",
	Run: func(cmd *cobra.Command, args []string) {
		card := &lib.Card{}
		c := NewHTTPClient()
		board, _ := cmd.Flags().GetString("board")
		stack, _ := cmd.Flags().GetString("stack")
		title, _ := cmd.Flags().GetString("title")
		err := card.Delete(c, board, stack, title)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Deleted \"%s\" card on %s stack on board %s\n", title, stack, board)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteBoardCmd)
	deleteCmd.AddCommand(deleteStackCmd)
	deleteCmd.AddCommand(deleteCardCmd)
	deleteBoardCmd.Flags().StringP("title", "t", "", "The title of the board")
	deleteStackCmd.Flags().StringP("board", "b", "", "The title of the board")
	deleteStackCmd.Flags().StringP("stack", "s", "", "The title of the stack")
	deleteCardCmd.Flags().StringP("board", "b", "", "The title of the board")
	deleteCardCmd.Flags().StringP("stack", "s", "", "The title of the stack")
	deleteCardCmd.Flags().StringP("title", "t", "", "The title of the card")
}
