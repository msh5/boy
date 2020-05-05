package interactor

import (
	"fmt"

	"github.com/msh5/boy/app/presenter"
	"github.com/msh5/boy/app/repository"
	"github.com/msh5/boy/app/usecase"
)

type SnippetShowInteractor struct {
	gistEntryRepository repository.GistEntryRepository
	showPresenter       presenter.ShowPresenter
}

func NewSnippetShowInteractor(
	gistEntryRepository repository.GistEntryRepository,
	showPresenter presenter.ShowPresenter,
) usecase.SnippetShowUsecase {
	return &SnippetShowInteractor{
		gistEntryRepository: gistEntryRepository,
		showPresenter:       showPresenter,
	}
}

type showNoFileError struct {
	msg string
}

func (e *showNoFileError) Error() string {
	return fmt.Sprintf("error: %s", e.msg)
}

func (i *SnippetShowInteractor) Run(params usecase.SnippetShowParameters) error {
	gistEntry, err := i.gistEntryRepository.Load(params.UserID, params.GistEntryName)
	if err != nil {
		return err
	}

	if len(gistEntry.Files) == 0 {
		return &showNoFileError{"no files in gist entry"}
	}

	result := presenter.ShowResult{Text: gistEntry.Files[0].Text}
	i.showPresenter.Present(result)

	return nil
}
