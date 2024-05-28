package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	ApiUrl  string
	Username string
	Password string
}

func (c *Config) Client() (*Client, error) {
	client := &Client{
		ApiUrl:  c.ApiUrl,
		Username: c.Username,
		Password: c.Password,
	}
	return client, nil
}

type Client struct {
	ApiUrl  string
	Username string
	Password string
}

func (c *Client) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.ApiUrl, endpoint)
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}
