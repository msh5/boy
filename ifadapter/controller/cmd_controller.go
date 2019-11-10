package controller

import (
	"errors"
	"net/url"
	"strings"

	"github.com/msh5/boy/app/usecase"
	"github.com/msh5/boy/ifadapter/di"
)

type CommandController struct {
	diContainer di.Container
}

func NewCommandController(diContainer di.Container) *CommandController {
	return &CommandController{diContainer: diContainer}
}

func (c *CommandController) ExecuteSnippet(gistRef string, commandArgs []string) error {
	gistHandle, err := parseGistReference(gistRef)
	if err != nil {
		return err
	}

	params := usecase.SnippetExecParameters{
		UserID:        gistHandle.UserID,
		GistEntryName: gistHandle.GistName,
		CommandArgs:   commandArgs,
	}

	return c.diContainer.GetSnippetExecUsecase().Run(params)
}

func (c *CommandController) ShowSnippet(gistRef string) error {
	gistHandle, err := parseGistReference(gistRef)
	if err != nil {
		return err
	}

	params := usecase.SnippetShowParameters{
		UserID:        gistHandle.UserID,
		GistEntryName: gistHandle.GistName,
	}

	return c.diContainer.GetSnippetShowUsecase().Run(params)
}

type gistHandle struct {
	UserID   string
	GistName string
}

func parseGistReference(ref string) (*gistHandle, error) {
	url, err := url.Parse("http://" + ref)
	if err != nil {
		return nil, err
	}

	segments := strings.Split(url.Path, "/")

	switch len(segments) {
	case 3:
		return &gistHandle{UserID: segments[1], GistName: segments[2]}, nil
	case 2:
		return &gistHandle{UserID: segments[0], GistName: segments[1]}, nil
	}

	return nil, errors.New("indicated gist reference is unexpected syntax")
}
