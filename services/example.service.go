package services

import (
	"errors"
	"laundry-api/context"
	"laundry-api/dto"
	"laundry-api/entities"
	"laundry-api/factory"
	"laundry-api/repositories"

	"github.com/stroiman/go-automapper"
	"gorm.io/gorm"
)

type ExampleService struct {
	ExampleRepository repositories.ExampleRepository
}

func NewExampleService(f *factory.Factory) *ExampleService {
	return &ExampleService{
		ExampleRepository: f.ExampleRepository,
	}
}

func (s *ExampleService) Create(c *context.CustomContex, payload *dto.ExampleRequest) (*entities.Example, error) {
	var entity = new(entities.Example)

	automapper.MapLoose(payload, &entity)

	err := c.Transaction(func(tx *gorm.DB) error {

		err := s.ExampleRepository.Create(tx, entity)
		if err != nil {
			return err
		}

		return errors.New("throw error")
		// return nil
	})

	return entity, err
}
