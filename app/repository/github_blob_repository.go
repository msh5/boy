package repository

import (
	"github.com/msh5/boy/domain/entity"
)

type GitHubBlobRepository interface {
	Load(repoOwner string, repoName string, path string) (*entity.GitHubBlob, error)
}
