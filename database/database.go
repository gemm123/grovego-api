package database

import (
	"gemm123/grovego-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.User{})

	return db, err
}

func CloseDB(db *gorm.DB) {
	dbSql, _ := db.DB()
	dbSql.Close()
}
