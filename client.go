package harbor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dukhyungkim/harbor-client/model"
	"io/ioutil"
	"log"
	"net/http"
)

type harborClient struct {
	baseURL string
	token   string
}

func NewHarborClient(config *HarborConfig) HarborClient {
	return &harborClient{
		baseURL: config.URL + "/api/v2.0",
		token:   base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.Username, config.Password))),
	}
}

func (c *harborClient) getJSON(url string, useToken bool) ([]byte, error) {
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

func (c *harborClient) getText(url string, useToken bool) ([]byte, error) {
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

func (c *harborClient) doRequest(req *http.Request) ([]byte, error) {
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

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var errors []*model.Error
		if err := json.Unmarshal(respData, &errors); err != nil {
			return nil, err
		}
		return nil, nil
	}

	return respData, nil
}
