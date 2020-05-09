package dependency

import (
	"github.com/msh5/boy/app/interactor"
	"github.com/msh5/boy/ifadapter/controller"
	"github.com/msh5/boy/ifadapter/di"
	"github.com/msh5/boy/ifadapter/output"
	"github.com/msh5/boy/interface/persistence"
	"github.com/msh5/boy/interface/view"
)

type CommandDIContainerBuildParameters struct {
	GitHubAccessToken string
	Ref               string
}

type CLIDependencies struct {
	ExecController *controller.ExecController
	ShowController *controller.ShowController

	ExecView *view.ExecConsoleView
	ShowView *view.ShowConsoleView
}

func NewCLIDependencies(params CommandDIContainerBuildParameters) (*CLIDependencies, error) {
	gistEntryPersistence, err := persistence.NewGistEntryPersistence(params.GitHubAccessToken, params.Ref)
	if err != nil {
		return nil, err
	}

	gitHubBlobPersistence, err := persistence.NewGitHubBlobPersistence(params.GitHubAccessToken, params.Ref)
	if err != nil {
		return nil, err
	}

	var execViewModel output.ExecViewModel
	snippetExecInteractor := interactor.NewSnippetExecInteractor(
		gistEntryPersistence,
		output.NewExecOutput(&execViewModel),
	)
	blobExecInteractor := interactor.NewBlobExecInteractor(
		gitHubBlobPersistence,
		output.NewExecOutput(&execViewModel),
	)

	var showViewModel output.ShowViewModel
	snippetShowInteractor := interactor.NewSnippetShowInteractor(
		gistEntryPersistence,
		output.NewShowOutput(&showViewModel),
	)
	blobShowInteractor := interactor.NewBlobShowInteractor(
		gitHubBlobPersistence,
		output.NewShowOutput(&showViewModel),
	)

	diBuilder := di.NewBuilder()
	diBuilder.RegisterSnippetExecUsecase(snippetExecInteractor)
	diBuilder.RegisterSnippetShowUsecase(snippetShowInteractor)
	diBuilder.RegisterBlobExecUsecase(blobExecInteractor)
	diBuilder.RegisterBlobShowUsecase(blobShowInteractor)

	diContainer := diBuilder.BuildContainer()

	return &CLIDependencies{
		ExecController: controller.NewExecController(diContainer),
		ShowController: controller.NewShowController(diContainer),

		ExecView: view.NewExecConsoleView(&execViewModel),
		ShowView: view.NewShowConsoleView(&showViewModel),
	}, nil
}
