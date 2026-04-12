package db_service

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitializeDB() error {
	var dsm = fmt.Sprintf("host=%s user=%s password=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	var err error

	Database, err = gorm.Open(postgres.Open(dsm), &gorm.Config{})

	return err
}
