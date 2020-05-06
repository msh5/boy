package persistence

import (
	"context"

	"github.com/msh5/boy/domain/entity"
	"github.com/msh5/boy/interface/driver"
)

type GistEntryPersistence struct {
	client *driver.GitHubClient
}

func NewGistEntryPersistence(accessToken, ref string) *GistEntryPersistence {
	return &GistEntryPersistence{
		client: driver.NewGitHubClient(accessToken, ref),
	}
}

func (r *GistEntryPersistence) Load(userID string, gistEntryName string) (*entity.GistEntry, error) {
	obj, err := r.client.QueryGistFiles(context.Background(), userID, gistEntryName)
	if err != nil {
		return nil, err
	}

	var entry entity.GistEntry
	for _, file := range obj.User.Gist.Files {
		entry.Files = append(entry.Files, entity.GistFile{Name: file.Name, Text: file.Text})
	}

	return &entry, nil
}
