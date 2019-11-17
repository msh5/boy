package interactor

import (
	"github.com/msh5/boy/app/presenter"
	"github.com/msh5/boy/app/repository"
	"github.com/msh5/boy/app/usecase"
)

type BlobExecInteractor struct {
	gitHubBlobRepository repository.GitHubBlobRepository
	execPresenter        presenter.ExecPresenter
}

func NewBlobExecInteractor(
	gitHubBlobRepository repository.GitHubBlobRepository,
	execPresenter presenter.ExecPresenter,
) usecase.BlobExecUsecase {
	return &BlobExecInteractor{
		gitHubBlobRepository: gitHubBlobRepository,
		execPresenter:        execPresenter,
	}
}

func (i *BlobExecInteractor) Run(params usecase.BlobExecParameters) error {
	blob, err := i.gitHubBlobRepository.Load(params.RepositoryOwner, params.RepositoryName, params.BlobPath)
	if err != nil {
		return err
	}

	exitStatus, err := runBytesAsCommand([]byte(blob.Text), params.CommandArgs)
	if err != nil {
		return err
	}

	result := presenter.ExecResult{ExitStatus: exitStatus}
	i.execPresenter.Present(result)

	return nil
}
