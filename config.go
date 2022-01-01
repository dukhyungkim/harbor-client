package harbor

const (
	urlProjects     = "/projects"
	urlPing         = "/ping"
	urlRepositories = "/projects/%s/repositories"
	urlArtifacts    = "/projects/%s/repositories/%s/artifacts"
)

type Config struct {
	URL      string
	Username string
	Password string
}
