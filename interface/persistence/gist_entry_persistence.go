package persistence

import (
	"context"

	"github.com/msh5/boy/domain/entity"
	"github.com/msh5/boy/interface/driver"
)

type GistEntryPersistence struct {
	client *driver.GitHubClient
}

func NewGistEntryPersistence(accessToken string) *GistEntryPersistence {
	return &GistEntryPersistence{
		client: driver.NewGitHubClient(accessToken),
	}
}

func (r *GistEntryPersistence) Load(handle entity.GistHandle) (*entity.GistEntry, error) {
	q, err := r.client.QueryUserGist(context.Background(), handle.UserID, handle.GistEntryName)
	if err != nil {
		return nil, err
	}

	var entry entity.GistEntry
	for _, file := range q.User.Gist.Files {
		entry.Files = append(entry.Files, entity.GistFile{Name: file.Name, Text: file.Text})
	}

	return &entry, nil
}
