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

// Board struct representing the Board item
type Board struct {
	Title string `json:"title"`
	Color string `json:"color"`
	ID    int    `json:"id"`
}

// Fetch board
func (b *Board) Fetch(c Client) []Board {
	resp, err := c.GetRequest(fmt.Sprintf("%s/index.php/apps/deck/api/v1.0/boards", c.Endpoint))
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var boards []Board
	jsonErr := json.Unmarshal(body, &boards)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return boards
}

// GetID from Board
func (b *Board) GetID(c Client, title string) int {
	boards := &Board{}
	boardlist := boards.Fetch(c)
	var id int
	for i := 0; i < len(boardlist); i++ {
		if boardlist[i].Title == title {
			id = boardlist[i].ID
		}
	}
	if id == 0 {
		fmt.Println(fmt.Errorf("No board with title %s found", title))
		os.Exit(1)
	}
	return id
}

// Delete Board
func (b *Board) Delete(c Client, title string) error {
	if title == "" {
		fmt.Println("Please provide a title")
		os.Exit(1)
	}
	boards := &Board{}
	boardid := boards.GetID(c, title)
	_, err := c.DeleteRequest(fmt.Sprintf("%s/index.php/apps/deck/api/v1.0/boards/%d", c.Endpoint, boardid))
	if err != nil {
		log.Fatal(err)
	} else {
		return nil
	}
	return nil
}

// New Board
func (b *Board) New(c Client, title, color string) error {
	if title == "" {
		fmt.Println("Please provide a title")
		os.Exit(1)
	}
	jsonStr := fmt.Sprintf("{\"title\": \"%s\", \"color\": \"%s\"}", title, color)
	var jsonData = []byte(jsonStr)
	_, err := c.PostRequest(fmt.Sprintf("%s/index.php/apps/deck/api/v1.0/boards", c.Endpoint), jsonData)
	if err != nil {
		log.Fatal(err)
	} else {
		return nil
	}
	return nil
}
