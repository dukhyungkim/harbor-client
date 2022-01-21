package harbor

import (
	"encoding/json"
	"fmt"
	"github.com/dukhyungkim/harbor-client/model"
)

func (c *Client) ListRepositories(projectName string, params *model.ListRepositoriesParams) ([]*model.Repository, error) {
	if params == nil {
		params = model.NewListRepositoriesParams()
	}

	url := fmt.Sprintf(urlListRepositories, projectName) + params.ToParamString()
	data, err := c.getJSON(url, true)
	if err != nil {
		return nil, err
	}

	var repositories []*model.Repository
	if err := json.Unmarshal(data, &repositories); err != nil {
		return nil, err
	}
	return repositories, nil
}

func (c *Client) GetRepository(projectName string, repositoryName string) (*model.Repository, error) {
	url := fmt.Sprintf(urlGetRepository, projectName, repositoryName)
	data, err := c.getJSON(url, true)
	if err != nil {
		return nil, err
	}

	var repository model.Repository
	if err := json.Unmarshal(data, &repository); err != nil {
		return nil, err
	}
	return &repository, nil
}
