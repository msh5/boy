package interactor

import (
	"github.com/msh5/boy/app/presenter"
	"github.com/msh5/boy/app/repository"
	"github.com/msh5/boy/app/usecase"
)

type BlobShowInteractor struct {
	gitHubBlobRepository repository.GitHubBlobRepository
	showPresenter        presenter.ShowPresenter
}

func NewBlobShowInteractor(
	gitHubBlobRepository repository.GitHubBlobRepository,
	showPresenter presenter.ShowPresenter,
) usecase.BlobShowUsecase {
	return &BlobShowInteractor{
		gitHubBlobRepository: gitHubBlobRepository,
		showPresenter:        showPresenter,
	}
}

func (i *BlobShowInteractor) Run(params usecase.BlobShowParameters) error {
	blob, err := i.gitHubBlobRepository.Load(params.RepositoryOwner, params.RepositoryName, params.BlobPath)
	if err != nil {
		return err
	}

	result := presenter.ShowResult{Text: blob.Text}
	i.showPresenter.Present(result)

	return nil
}
