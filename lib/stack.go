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
	Title        string `json:"title"`
	BoardID      int    `json:"boardId"`
	DeletedAt    int    `json:"deletedAt"`
	LastModified int    `json:"lastModified"`
	Order        int    `json:"order"`
	ID           int    `json:"id"`
}

// Fetch stack
func (s *Stack) Fetch(c Client, boardtitle string) []Stack {
	boards := &Board{}
	boardid := boards.GetID(c, boardtitle)
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

// GetID from stack
func (s *Stack) GetID(c Client, boardtitle string, title string) int {
	stacks := &Stack{}
	stacklist := stacks.Fetch(c, boardtitle)
	var id int
	for i := 0; i < len(stacklist); i++ {
		if stacklist[i].Title == title {
			id = stacklist[i].ID
		}
	}
	if id == 0 {
		fmt.Println(fmt.Errorf("No stack with title %s found", title))
		os.Exit(1)
	}
	return id
}

// New Stack
func (s *Stack) New(c Client, board, title string, order int) error {
	boards := &Board{}
	boardid := boards.GetID(c, board)
	var _order int
	if title == "" {
		fmt.Println("Please provide a title for the new stack")
		os.Exit(1)
	}
	if board == "" {
		fmt.Println("Please provide a board where the stack needs to be made")
		os.Exit(1)
	}
	if order == 0 {
		_order = 999
	} else {
		_order = order
	}
	jsonStr := fmt.Sprintf("{\"title\": \"%s\",\"order\": \"%d\"}", title, _order)
	var jsonData = []byte(jsonStr)
	_, err := c.PostRequest(fmt.Sprintf("%s/index.php/apps/deck/api/v1.0/boards/%d/stacks", c.Endpoint, boardid), jsonData)
	if err != nil {
		log.Fatal(err)
	} else {
		return nil
	}
	return nil
}
