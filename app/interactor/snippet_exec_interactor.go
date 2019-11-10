package interactor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/msh5/boy/app/repository"
	"github.com/msh5/boy/app/usecase"
	"github.com/msh5/boy/domain/entity"
)

type SnippetExecInteractor struct {
	gistEntryRepository repository.GistEntryRepositoryInterface
}

func NewSnippetExecInteractor(repo repository.GistEntryRepositoryInterface) usecase.SnippetExecUsecase {
	return &SnippetExecInteractor{gistEntryRepository: repo}
}

func (i *SnippetExecInteractor) Run(params usecase.SnippetExecParameters) error {
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

	tempFile, err := writeToTempFile([]byte(gistEntry.Files[0].Text))
	if err != nil {
		return err
	}

	defer os.Remove(tempFile.Name())

	if err := os.Chmod(tempFile.Name(), 0755); err != nil {
		return err
	}

	command := exec.Command(tempFile.Name(), params.CommandArgs...) //nolint:gosec
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()
}

func writeToTempFile(b []byte) (*os.File, error) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "boy_*")
	if err != nil {
		return nil, err
	}

	_, err = tempFile.Write(b)

	tempFile.Close()

	if err != nil {
		os.Remove(tempFile.Name())
		return nil, err
	}

	return tempFile, nil
}
