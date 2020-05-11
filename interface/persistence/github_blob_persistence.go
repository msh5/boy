package persistence

import (
	"context"

	"github.com/msh5/boy/domain/entity"
	"github.com/msh5/boy/interface/driver"
)

type GitHubBlobPersistence struct {
	client *driver.GitHubClient
}

func NewGitHubBlobPersistence(accessToken, url string) *GitHubBlobPersistence {
	return &GitHubBlobPersistence{
		client: driver.NewGitHubClient(accessToken, url),
	}
}

func (r *GitHubBlobPersistence) Load(repoOwner string, repoName string, path string) (*entity.GitHubBlob, error) {
	obj, err := r.client.QueryRepositoryBlob(context.Background(), repoOwner, repoName, path)
	if err != nil {
		return nil, err
	}

	return &entity.GitHubBlob{Text: obj.Repository.Object.Blob.Text}, nil
}
