package controller

import (
	"github.com/msh5/boy/app/usecase"
	"github.com/msh5/boy/ifadapter/di"
)

type CommandController struct {
	diContainer di.Container
}

func NewCommandController(diContainer di.Container) *CommandController {
	return &CommandController{diContainer: diContainer}
}

func (c *CommandController) ExecuteSnippet(args []string) error {
	params := usecase.SnippetExecParameters{
		UserID:        args[0],
		GistEntryName: args[1],
		CommandArgs:   args[2:],
	}

	return c.diContainer.GetSnippetExecUsecase().Run(params)
}

func (c *CommandController) ShowSnippet(args []string) error {
	params := usecase.SnippetShowParameters{
		UserID:        args[0],
		GistEntryName: args[1],
	}

	return c.diContainer.GetSnippetShowUsecase().Run(params)
}
