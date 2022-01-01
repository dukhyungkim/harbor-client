package harbor

import (
	"github.com/dukhyungkim/harbor-client/model"
)

type Client interface {
	ListProjects(params *model.ListProjectParams) ([]*model.Project, error)
	Ping() (string, error)
	ListRepositories(projectName string, params *model.ListRepositoriesParams) ([]*model.Repository, error)
	ListArtifacts(projectName string, repositoryName string, params *model.ListArtifactsParams) ([]*model.Artifact, error)
}
