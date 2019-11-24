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
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"os"
)

// Card struct representing the Stack item
type Card struct {
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	StackID         int           `json:"stackId"`
	Type            string        `json:"type"`
	LastModified    int           `json:"lastModified"`
	LastEditor      interface{}   `json:"lastEditor"`
	CreatedAt       int           `json:"createdAt"`
	Labels          interface{}   `json:"labels"`
	AssignedUsers   []interface{} `json:"assignedUsers"`
	Attachments     interface{}   `json:"attachments"`
	AttachmentCount int           `json:"attachmentCount"`
	Owner           interface{}   `json:"owner"`
	Order           int           `json:"order"`
	Archived        bool          `json:"archived"`
	Duedate         interface{}   `json:"duedate"`
	DeletedAt       int           `json:"deletedAt"`
	CommentsUnread  int           `json:"commentsUnread"`
	ID              int           `json:"id"`
	Overdue         int           `json:"overdue"`
}

// Fetch list of cards
func (cd *Card) Fetch(c Client, boardtitle, stacktitle string) []gjson.Result {
	boards := &Board{}
	stacks := &Stack{}
	boardid := boards.GetID(c, boardtitle)
	stackid := stacks.GetID(c, boardtitle, stacktitle)

	resp, err := c.GetRequest(fmt.Sprintf("%s/index.php/apps/deck/api/v1.0/boards/%d/stacks/%d", c.Endpoint, boardid, stackid))
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	stringjson := (string(body))
	cardtitles := gjson.Get(stringjson, "cards.#.title")
	if len(cardtitles.Array()) <= 0 {
		fmt.Println(fmt.Errorf("No cards on stack %s found", stacktitle))
		os.Exit(1)
	}
	return cardtitles.Array()
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
