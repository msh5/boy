package cmd

import (
	"net/url"

	"github.com/msh5/boy/interface/dependency"
)

const (
	GithubHost = "github.com"
	GistHost   = "gist.github.com"
)

func setDependencies(config commandConfig, ref string) (*dependency.CLIDependencies, error) {
	params := config.toDIContainerBuildParams()

	parsedURL, err := url.Parse("https://" + ref)
	if err != nil {
		return nil, err
	}

	if parsedURL.Host != GithubHost && parsedURL.Host != GistHost {
		parsedURL.Path = "/api/graphql"
		params.EnterpriseHostname = parsedURL.String()
		params.IsEnterprise = true
	}

	dependencies := dependency.NewCLIDependencies(params)

	return dependencies, nil
}
