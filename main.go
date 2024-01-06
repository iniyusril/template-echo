package main

import (
	"laundry-api/factory"
	"laundry-api/middleware"
	"laundry-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	f := factory.NewFactory()
	app.Use(middleware.FillCustomContex)
	route := routes.NewRoutes(app, f)
	route.RegisterRoutes()

	app.Logger.Fatal(app.Start(":9000"))
}
