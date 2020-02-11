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
	"bytes"
	"fmt"
	"net/http"
)

// Client struct for http client
type Client struct {
	Endpoint string
	Username string
	Password string
}

// GetRequest method to make GET request to the NextCloud Deck API
func (c *Client) GetRequest(url string) (*http.Response, error) {
	return c.Request("GET", url, 200, nil)
}

// PostRequest method to make POST request to the NextCloud Deck API
func (c *Client) PostRequest(url string, i []byte) (*http.Response, error) {
	return c.Request("POST", url, 200, i)
}

// Request function to handle all API request
func (c *Client) Request(verb, url string, code int, payload []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(verb, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("OCI-APIRequest", "true")
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != code {
		return resp, fmt.Errorf("Received %d, expecting %d status code while fetching %s", resp.StatusCode, code, url)
	}
	return resp, err
}
