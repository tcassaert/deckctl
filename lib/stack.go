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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Stack struct representing the Stack item
type Stack struct {
	BoardID int    `json:"boardId"`
	Title   string `json:"title"`
	ID      int    `json:"id"`
}

// Fetch stack
func (s *Stack) Fetch(c Client, title string) []Stack {
	boards := &Board{}
	boardid := boards.GetID(c, title)
	resp, err := c.GetRequest(fmt.Sprintf("%s/index.php/apps/deck/api/v1.0/boards/%d/stacks", c.Endpoint, boardid))
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var stacks []Stack
	jsonErr := json.Unmarshal(body, &stacks)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return stacks
}

// New Stack
func (s *Stack) New(c Client, board, title string, order int) error {
	boards := &Board{}
	boardid := boards.GetID(c, board)
	if title == "" {
		fmt.Println("Please provide a title for the new stack")
		os.Exit(1)
	}
	if board == "" {
		fmt.Println("Please provide a board where the stack needs to be made")
		os.Exit(1)
	}
	if order == 0 {
		fmt.Println("Please provide an order number for the stack")
		os.Exit(1)
	}
	jsonStr := fmt.Sprintf("{\"title\": \"%s\",\"order\": \"%d\"}", title, order)
	var jsonData = []byte(jsonStr)
	_, err := c.PostRequest(fmt.Sprintf("%s/index.php/apps/deck/api/v1.0/boards/%d/stacks", c.Endpoint, boardid), jsonData)
	if err != nil {
		log.Fatal(err)
	} else {
		return nil
	}
	return nil
}
