package harbor

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	baseURL string
	token   string
}

func NewClient(config *Config) *Client {
	return &Client{
		baseURL: config.URL + "/api/v2.0",
		token:   base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.Username, config.Password))),
	}
}

func (c *Client) getJSON(url string, useToken bool) ([]byte, error) {
	req, err := http.NewRequest("GET", c.baseURL+url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	if useToken {
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.token))
	}
	return c.doRequest(req)
}

func (c *Client) getText(url string, useToken bool) ([]byte, error) {
	req, err := http.NewRequest("GET", c.baseURL+url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "text/plain")
	if useToken {
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.token))
	}
	return c.doRequest(req)
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body cleanly; %v", err)
		}
	}()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(respData))
	}

	return respData, nil
}
