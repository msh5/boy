package interactor

import (
	"fmt"

	"github.com/msh5/boy/app/presenter"
	"github.com/msh5/boy/app/repository"
	"github.com/msh5/boy/app/usecase"
)

type SnippetExecInteractor struct {
	gistEntryRepository repository.GistEntryRepository
	execPresenter       presenter.ExecPresenter
}

func NewSnippetExecInteractor(
	gistEntryRepository repository.GistEntryRepository,
	execPresenter presenter.ExecPresenter,
) usecase.SnippetExecUsecase {
	return &SnippetExecInteractor{
		gistEntryRepository: gistEntryRepository,
		execPresenter:       execPresenter,
	}
}

type execNoFileError struct {
	msg string
}

func (e *execNoFileError) Error() string {
	return fmt.Sprintf("error: %s", e.msg)
}

func (i *SnippetExecInteractor) Run(params usecase.SnippetExecParameters) error {
	gistEntry, err := i.gistEntryRepository.Load(params.UserID, params.GistEntryName)
	if err != nil {
		return err
	}

	if len(gistEntry.Files) == 0 {
		return &execNoFileError{msg: "no files in gist entry"}
	}

	exitStatus, err := runBytesAsCommand([]byte(gistEntry.Files[0].Text), params.CommandArgs)
	if err != nil {
		return err
	}

	result := presenter.ExecResult{ExitStatus: exitStatus}
	i.execPresenter.Present(result)

	return nil
}
