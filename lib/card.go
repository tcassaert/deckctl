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

package lib

import (
	"fmt"
	"log"
	"os"
)

// Card struct representing the Stack item
type Card struct {
	BoardID     int    `json:"boardId"`
	Description string `json:"description"`
	ID          int    `json:"id"`
	Order       int    `json:"order"`
	StackID     int    `json:"stackId"`
	Title       string `json:"title"`
}

// New Card
func (cd *Card) New(c Client, board, stack, title string, order int) error {
	boards := &Board{}
	stacks := &Stack{}
	boardid := boards.GetID(c, board)
	stackid := stacks.GetID(c, board, stack)
	var _order int
	if title == "" {
		fmt.Println("Please provide a title for the new stack")
		os.Exit(1)
	}
	if board == "" {
		fmt.Println("Please provide a board where the stack needs to be made")
		os.Exit(1)
	}
	if stack == "" {
		fmt.Println("Please provide a stack where the card needs to be made")
		os.Exit(1)
	}
	if order == 0 {
		_order = 999
	} else {
		_order = order
	}
	jsonStr := fmt.Sprintf("{\"title\": \"%s\",\"order\": \"%d\", \"type\": \"plain\"}", title, _order)
	var jsonData = []byte(jsonStr)
	_, err := c.PostRequest(fmt.Sprintf("%s/index.php/apps/deck/api/v1.0/boards/%d/stacks/%d/cards", c.Endpoint, boardid, stackid), jsonData)
	if err != nil {
		log.Fatal(err)
	} else {
		return nil
	}
	return nil
}
