package factory

import (
	"template-echo/db"
	"template-echo/repositories"

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
