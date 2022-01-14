package harbor

const (
	urlProjects     = "/projects"
	urlPing         = "/ping"
	urlRepositories = "/projects/%s/repositories"
	urlArtifacts    = "/projects/%s/repositories/%s/artifacts"
	urlTags         = "/projects/%s/repositories/%s/artifacts/%s/tags"
)

type Config struct {
	URL      string
	Username string
	Password string
}
