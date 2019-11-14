package dependency

import (
	"github.com/msh5/boy/app/interactor"
	"github.com/msh5/boy/ifadapter/di"
	"github.com/msh5/boy/interface/persistence"
)

type CommandDIContainerBuildParameters struct {
	GitHubAccessToken string
}

func NewCommandDIContainer(params CommandDIContainerBuildParameters) di.Container {
	diBuilder := di.NewBuilder()

	gistEntryRepository := persistence.NewGistEntryPersistence(params.GitHubAccessToken)

	diBuilder.RegisterSnippetExecUsecase(
		interactor.NewSnippetExecInteractor(gistEntryRepository),
	)
	diBuilder.RegisterSnippetShowUsecase(
		interactor.NewSnippetShowInteractor(gistEntryRepository),
	)

	return diBuilder.BuildContainer()
}
