package repositories

import (
	"template-fiber/entities"

	"gorm.io/gorm"
)

type ExampleRepository struct {
}

func NewExampleRepository() *ExampleRepository {
	return &ExampleRepository{}
}

func (r *ExampleRepository) Create(db *gorm.DB, data *entities.Example) error {
	return db.Create(data).Error
}
