package middleware

import (
	"laundry-api/context"
	"laundry-api/db"

	"github.com/labstack/echo/v4"
)

func FillCustomContex(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &context.CustomContex{
			Context: c,
			DB:      db.Get(),
		}
		return next(cc)
	}
}
