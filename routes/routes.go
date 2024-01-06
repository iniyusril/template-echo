package routes

import (
	"template-fiber/controller"
	"template-fiber/factory"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	app               *echo.Echo
	ExampleController *controller.ExampleController
}

func NewRoutes(app *echo.Echo, f *factory.Factory) Routes {
	return Routes{
		app:               app,
		ExampleController: controller.NewExampleController(f),
	}
}

func (r Routes) RegisterRoutes() {
	//register routes here
	app := r.app

	exampleGroup := app.Group("v1/example")
	{
		exampleGroup.POST("", r.ExampleController.Create)
	}

}
