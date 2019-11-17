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
		panic("cannot cast to SnippetExecUsecase")
	}

	return usecase
}

func (c *Container) GetSnippetShowUsecase() usecase.SnippetShowUsecase {
	objectInterface := c.Get(snippetShowUsecaseDIObject)

	usecase, ok := objectInterface.(usecase.SnippetShowUsecase)
	if !ok {
		panic("cannot cast to SnippetShowUsecase")
	}

	return usecase
}

func (c *Container) GetBlobExecUsecase() usecase.BlobExecUsecase {
	objectInterface := c.Get(blobExecUsecaseDIObject)

	usecase, ok := objectInterface.(usecase.BlobExecUsecase)
	if !ok {
		panic("cannot cast to BlobExecUsecase")
	}

	return usecase
}

func (c *Container) GetBlobShowUsecase() usecase.BlobShowUsecase {
	objectInterface := c.Get(blobShowUsecaseDIObject)

	usecase, ok := objectInterface.(usecase.BlobShowUsecase)
	if !ok {
		panic("cannot cast to BlobShowUsecase")
	}

	return usecase
}
