package cmd

import (
	"net/url"

	"github.com/msh5/boy/interface/dependency"
)

const (
	GithubHost = "github.com"
	GistHost   = "gist.github.com"
)

func newDependencies(config commandConfig, ref string) (*dependency.CLIDependencies, error) {
	params := config.toDIContainerBuildParams()

	if err := setParamsIfEnterprise(&params, ref); err != nil {
		return nil, err
	}

	dependencies, err := dependency.NewCLIDependencies(params)
	if err != nil {
		return nil, err
	}

	return dependencies, nil
}

func setParamsIfEnterprise(params *dependency.CommandDIContainerBuildParameters, ref string) error {
	parsedURL, err := url.Parse("https://" + ref)
	if err != nil {
		return err
	}

	if parsedURL.Host != GithubHost && parsedURL.Host != GistHost {
		params.EnterpriseHostname = parsedURL.Hostname()
		params.IsEnterprise = true
	}

	return nil
}
