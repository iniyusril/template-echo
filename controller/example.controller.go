package controller

import (
	"net/http"
	"template-fiber/context"
	"template-fiber/dto"
	"template-fiber/dto/response"
	"template-fiber/factory"
	"template-fiber/services"

	"github.com/labstack/echo/v4"
)

type ExampleController struct {
	ExampleService *services.ExampleService
}

func NewExampleController(f *factory.Factory) *ExampleController {
	return &ExampleController{
		ExampleService: services.NewExampleService(f),
	}
}

func (s *ExampleController) Create(c echo.Context) error {

	var payload = new(dto.ExampleRequest)
	var cc = c.(*context.CustomContex)

	if err := c.Bind(payload); err != nil {
		res := response.ErrorBuilder(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	data, err := s.ExampleService.Create(cc, payload)
	if err != nil {
		errRes := response.ErrorBuilder(err)
		return c.JSON(http.StatusInternalServerError, errRes)
	}

	return c.JSON(http.StatusOK, response.SuccessBuilder(data))
}
