package repository

import (
	"github.com/msh5/boy/domain/entity"
)

type GistEntryRepository interface {
	Load(handle entity.GistHandle) (*entity.GistEntry, error)
}
