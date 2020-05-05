package controller

import (
	"github.com/msh5/boy/app/usecase"
	"github.com/msh5/boy/ifadapter/di"
	"golang.org/x/xerrors"
)

type ExecController struct {
	diContainer di.Container
}

func NewExecController(diContainer di.Container) *ExecController {
	return &ExecController{diContainer: diContainer}
}

func (c *ExecController) Handle(ref string, commandArgs []string) error {
	refType := detectReferenceType(ref)

	switch refType {
	case gistFileReferenceType:
		return c.executeGistFileSnippet(ref, commandArgs)
	case gitHubBlobReferenceType:
		return c.executeBlobSnippet(ref, commandArgs)
	default:
		return xerrors.Errorf("reference is unexpected syntax")
	}
}

func (c *ExecController) executeGistFileSnippet(ref string, commandArgs []string) error {
	refObj, err := parseGistFileReference(ref)
	if err != nil {
		return err
	}

	return c.diContainer.GetSnippetExecUsecase().Run(usecase.SnippetExecParameters{
		UserID:        refObj.UserID,
		GistEntryName: refObj.GistName,
		CommandArgs:   commandArgs,
	})
}

func (c *ExecController) executeBlobSnippet(ref string, commandArgs []string) error {
	refObj, err := parseGitHubBlobReference(ref)
	if err != nil {
		return err
	}

	return c.diContainer.GetBlobExecUsecase().Run(usecase.BlobExecParameters{
		RepositoryOwner: refObj.RepositoryOwner,
		RepositoryName:  refObj.RepositoryName,
		BlobPath:        refObj.BlobPath,
		CommandArgs:     commandArgs,
	})
}
