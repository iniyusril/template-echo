package context

import (
	"context"
	"database/sql"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomContex struct {
	echo.Context
	DB *gorm.DB
}

func (c *CustomContex) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return c.DB.WithContext(c.Request().Context()).Transaction(
		fc, opts...,
	)

}

type DbContext struct {
	DB  *gorm.DB
	Ctx context.Context
}
