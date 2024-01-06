package main

import (
	"template-echo/factory"
	"template-echo/middleware"
	"template-echo/routes"

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
