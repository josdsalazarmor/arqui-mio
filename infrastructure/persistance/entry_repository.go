package persistance

import (
	"github.com/jinzhu/gorm"
	"github.com/multimedia_ms/domain/model"
	"github.com/multimedia_ms/usecase/repository"
)

type entryRepository struct {
	db *gorm.DB
}

// NewEntryRepository creates a repository using MySQL for datasotre.
func NewEntryRepository(db *gorm.DB) repository.EntryRepository {
	return &entryRepository{db}
}

func (eR *entryRepository) FindByID(userId int) (*model.Files, error) {
	files := model.Files{UserId: userId}
	err := eR.db.First(&files, userId).Error
	if err != nil {
		return nil, err
	}

	return &files, nil
}

func (eR *entryRepository) Store(files *model.Files) error {
	return eR.db.Save(files).Error
}

func (eR *entryRepository) Update(files *model.Files) error {
	// Save will include all fields when perform the Updating SQL, even it is not changed
	return eR.db.Model(&model.Files{UserId: files.UserId}).Updates(files).Error
}

func (eR *entryRepository) Delete(files *model.Files) error {
	return eR.db.Delete(files).Error
}

func (eR *entryRepository) FindAll() ([]*model.Files, error) {
	filess := []*model.Files{}
		
	err := eR.db.Find(&filess).Error
	if err != nil {
		return nil, err

	}
	return filess,nil
}
