package harbor

import (
	"encoding/json"
	"fmt"
	"github.com/dukhyungkim/harbor-client/model"
)

func (c *client) ListArtifacts(projectName string, repositoryName string, params *model.ListArtifactsParams) ([]*model.Artifact, error) {
	if params == nil {
		params = model.NewListArtifactsParams()
	}

	url := fmt.Sprintf(urlArtifacts, projectName, repositoryName) + params.ToParamString()
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

func (c *client) ListTags(projectName string, repositoryName string, reference string, params *model.ListTagsParams) ([]*model.Tag, error) {
	if params == nil {
		params = model.NewListTagsParams()
	}

	url := fmt.Sprintf(urlTags, projectName, repositoryName, reference) + params.ToParamString()
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
