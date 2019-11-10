package di

import (
	"github.com/sarulabs/di"

	"github.com/msh5/boy/app/usecase"
)

type Container struct {
	di.Container
}

func (c *Container) GetSnippetExecUsecase() usecase.SnippetExecUsecase {
	objectInterface := c.Get(snippetExecUsecaseDIObject)

	usecase, ok := objectInterface.(usecase.SnippetExecUsecase)
	if !ok {
		panic("Cannot cast to SnippetExecUsecase")
	}

	return usecase
}

func (c *Container) GetSnippetShowUsecase() usecase.SnippetShowUsecase {
	objectInterface := c.Get(snippetShowUsecaseDIObject)

	usecase, ok := objectInterface.(usecase.SnippetShowUsecase)
	if !ok {
		panic("Cannot cast to SnippetShowUsecase")
	}

	return usecase
}
