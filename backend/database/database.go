package database

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		dsn := "postgres://school:schoolmanagement@localhost:5432/schoolmanagement"
		if dsn == "" {
			log.Fatal("DATABASE_URL environment variable is required")
		}

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
	})
	return db
}
