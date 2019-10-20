package repository

import (
	"github.com/multimedia_ms/domain/model"
)

type EntryRepository interface {
	FindByID(id int) (*model.Files, error)
	Store(files *model.Files) error
	Update(files *model.Files) error
	Delete(files *model.Files) error
	FindAll() ([]*model.Files, error)
}
