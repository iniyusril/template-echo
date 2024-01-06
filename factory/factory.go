package factory

import (
	"laundry-api/db"
	"laundry-api/repositories"

	"gorm.io/gorm"
)

type Factory struct {
	ExampleRepository repositories.ExampleRepository
	DB                *gorm.DB
}

func NewFactory() *Factory {
	dbf := db.Get()
	return &Factory{
		ExampleRepository: *repositories.NewExampleRepository(),
		DB:                dbf,
	}
}
