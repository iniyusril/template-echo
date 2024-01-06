package db

import (
	"sync"
	"template-fiber/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once
var err error

func Init() {
	Get()
}

func Get() *gorm.DB {

	once.Do(func() {
		dsn := "host=localhost user=yusril password=hahahehe dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&entities.Example{})
	})

	return db
}
