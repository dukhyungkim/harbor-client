package harbor

const (
	urlPing             = "/ping"
	urlListProjects     = "/projects"
	urlGetProject       = "/projects/%s"
	urlListRepositories = "/projects/%s/repositories"
	urlGetRepository    = "/projects/%s/repositories/%s"
	urlListArtifacts    = "/projects/%s/repositories/%s/artifacts"
	urlGetArtifact      = "/projects/%s/repositories/%s/artifacts/%s"
	urlListTags         = "/projects/%s/repositories/%s/artifacts/%s/tags"
)

type Config struct {
	URL      string
	Username string
	Password string
}
