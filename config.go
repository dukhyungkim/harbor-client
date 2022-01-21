package harbor

const (
	urlPing           = "/ping"
	urlProjects       = "/projects"
	urlProjectInfo    = "/projects/%s"
	urlRepositories   = "/projects/%s/repositories"
	urlRepositoryInfo = "/projects/%s/repositories/%s"
	urlArtifacts      = "/projects/%s/repositories/%s/artifacts"
	urlArtifactInfo   = "/projects/%s/repositories/%s/artifacts/%s"
	urlTags           = "/projects/%s/repositories/%s/artifacts/%s/tags"
)

type Config struct {
	URL      string
	Username string
	Password string
}
