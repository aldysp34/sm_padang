package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabaseConfig() *gorm.DB {
	dbURL := "postgres://postgres:postgres@localhost:5432/sm_padang"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db

}
