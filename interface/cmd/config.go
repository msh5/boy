package cmd

import (
	"github.com/spf13/viper"

	"github.com/msh5/boy/interface/dependency"
)

type commandConfig struct {
	GitHubAccessToken string
}

func loadCommandConfig() commandConfig {
	var config commandConfig

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}

func (c *commandConfig) toDIContainerBuildParams() dependency.CommandDIContainerBuildParameters {
	return dependency.CommandDIContainerBuildParameters{
		GitHubAccessToken: c.GitHubAccessToken,
	}
}
