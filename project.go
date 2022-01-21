package harbor

import (
	"encoding/json"
	"fmt"
	"github.com/dukhyungkim/harbor-client/model"
)

func (c *Client) ListProjects(params *model.ListProjectParams) ([]*model.Project, error) {
	if params == nil {
		params = model.NewListProjectsParams()
	}

	url := urlProjects + params.ToParamString()
	data, err := c.getJSON(url, true)
	if err != nil {
		return nil, err
	}

	var projects []*model.Project
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (c *Client) GetProject(projectName string) (*model.Project, error) {
	url := fmt.Sprintf(urlProjectInfo, projectName)
	data, err := c.getJSON(url, true)
	if err != nil {
		return nil, err
	}

	var project model.Project
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, err
	}
	return &project, nil
}
