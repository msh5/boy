package repository

import (
	"github.com/msh5/boy/domain/entity"
)

type GistEntryRepository interface {
	Load(userID string, gistEntryName string) (*entity.GistEntry, error)
}
