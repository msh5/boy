package interactor

import (
	"fmt"

	"github.com/msh5/boy/app/repository"
	"github.com/msh5/boy/app/usecase"
	"github.com/msh5/boy/domain/entity"
)

type SnippetShowInteractor struct {
	gistEntryRepository repository.GistEntryRepository
}

func NewSnippetShowInteractor(repo repository.GistEntryRepository) usecase.SnippetShowUsecase {
	return &SnippetShowInteractor{gistEntryRepository: repo}
}

func (i *SnippetShowInteractor) Run(params usecase.SnippetShowParameters) error {
	gistEntry, err := i.gistEntryRepository.Load(entity.GistHandle{
		UserID:        params.UserID,
		GistEntryName: params.GistEntryName,
	})
	if err != nil {
		return err
	}

	if len(gistEntry.Files) == 0 {
		return fmt.Errorf("no files in gist entry")
	}

	fmt.Printf("%s\n", gistEntry.Files[0].Text)

	return nil
}
