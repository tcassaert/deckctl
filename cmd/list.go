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
  "encoding/json"
  "fmt"
  "log"
  "io/ioutil"

  "github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List object",
	Long: `List objects in your NextCloud Desk app.

These objects can be cards, stacks and boards`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

type Board struct {
	Title        string        `json:"title"`
	Color        string        `json:"color"`
	Archived     bool          `json:"archived"`
	Labels       []interface{} `json:"labels"`
	ACL          []interface{} `json:"acl"`
	Users        []interface{} `json:"users"`
	Shared       int           `json:"shared"`
	Stacks       []interface{} `json:"stacks"`
	DeletedAt    int           `json:"deletedAt"`
	LastModified int           `json:"lastModified"`
	ID           int           `json:"id"`
}

var listBoards = &cobra.Command{
	Use:   "boards",
	Short: "List object",
	Long: `List objects in your NextCloud Deck app.

These objects can be cards, stacks and boards`,
	Run: func(cmd *cobra.Command, args []string) {
    c := NewHttpClient()
    resp, err := c.GetRequest(fmt.Sprintf("%s/index.php/apps/deck/api/v1.0/boards", c.Endpoint))
    if err != nil {
      log.Fatal(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    var decoded []Board
    jsonErr := json.Unmarshal(body, &decoded)
    if jsonErr != nil {
      fmt.Println(jsonErr)
	  }
    fmt.Println("Your boards are:")
    for i := 0; i < len(decoded); i++ {
      fmt.Println(decoded[i].Title)
    }
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
  listCmd.AddCommand(listBoards)
}
