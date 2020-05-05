package persistence

import (
	"context"
	"fmt"

	"github.com/msh5/boy/domain/entity"
	"github.com/msh5/boy/interface/driver"
)

type GistEntryPersistence struct {
	client *driver.GitHubClient
}

type gistError struct {
	msg string
}

func (e *gistError) Error() string {
	return fmt.Sprintf("err: %s", e.msg)
}

func NewGistEntryPersistence(accessToken string) *GistEntryPersistence {
	return &GistEntryPersistence{
		client: driver.NewGitHubClient(accessToken),
	}
}

func (r *GistEntryPersistence) Load(userID string, gistEntryName string) (*entity.GistEntry, error) {
	obj, err := r.client.QueryGistFiles(context.Background(), userID, gistEntryName)
	if err != nil {
		return nil, err
	}

	if len(obj.User.Gist.Files) == 0 {
		return nil, &gistError{msg: "no files in gist entry"}
	}

	var entry entity.GistEntry
	for _, file := range obj.User.Gist.Files {
		entry.Files = append(entry.Files, entity.GistFile{Name: file.Name, Text: file.Text})
	}

	return &entry, nil
}
