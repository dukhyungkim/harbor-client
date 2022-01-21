package harbor

import (
	"encoding/json"
	"fmt"
	"github.com/dukhyungkim/harbor-client/model"
)

func (c *Client) ListArtifacts(projectName string, repositoryName string, params *model.ListArtifactsParams) ([]*model.Artifact, error) {
	if params == nil {
		params = model.NewListArtifactsParams()
	}

	url := fmt.Sprintf(urlListArtifacts, projectName, repositoryName) + params.ToParamString()
	data, err := c.getJSON(url, true)
	if err != nil {
		return nil, err
	}

	var artifacts []*model.Artifact
	if err := json.Unmarshal(data, &artifacts); err != nil {
		return nil, err
	}
	return artifacts, nil
}

func (c *Client) GetArtifact(projectName string, repositoryName string, reference string) (*model.Artifact, error) {
	url := fmt.Sprintf(urlGetArtifact, projectName, repositoryName, reference)
	data, err := c.getJSON(url, true)
	if err != nil {
		return nil, err
	}

	var artifact model.Artifact
	if err := json.Unmarshal(data, &artifact); err != nil {
		return nil, err
	}
	return &artifact, nil
}

func (c *Client) ListTags(projectName string, repositoryName string, reference string, params *model.ListTagsParams) ([]*model.Tag, error) {
	if params == nil {
		params = model.NewListTagsParams()
	}

	url := fmt.Sprintf(urlListTags, projectName, repositoryName, reference) + params.ToParamString()
	data, err := c.getJSON(url, true)
	if err != nil {
		return nil, err
	}

	var tags []*model.Tag
	if err := json.Unmarshal(data, &tags); err != nil {
		return nil, err
	}
	return tags, nil
}
