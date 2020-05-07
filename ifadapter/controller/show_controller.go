package controller

import (
	"github.com/msh5/boy/app/usecase"
	"github.com/msh5/boy/ifadapter/di"
)

type ShowController struct {
	diContainer di.Container
}

func NewShowController(diContainer di.Container) *ShowController {
	return &ShowController{diContainer: diContainer}
}

func (c *ShowController) Handle(ref string) error {
	refType := detectReferenceType(ref)

	switch refType {
	case gistFileReferenceType:
		return c.showGistFileSnippet(ref)
	case gitHubBlobReferenceType:
		return c.showBlobSnippet(ref)
	default:
		return &unknownReferenceTypeError{}
	}
}

func (c *ShowController) showGistFileSnippet(ref string) error {
	refObj, err := parseGistFileReference(ref)
	if err != nil {
		return err
	}

	return c.diContainer.GetSnippetShowUsecase().Run(usecase.SnippetShowParameters{
		UserID:        refObj.UserID,
		GistEntryName: refObj.GistName,
	})
}

func (c *ShowController) showBlobSnippet(ref string) error {
	refObj, err := parseGitHubBlobReference(ref)
	if err != nil {
		return err
	}

	return c.diContainer.GetBlobShowUsecase().Run(usecase.BlobShowParameters{
		RepositoryOwner: refObj.RepositoryOwner,
		RepositoryName:  refObj.RepositoryName,
		BlobPath:        refObj.BlobPath,
	})
}
