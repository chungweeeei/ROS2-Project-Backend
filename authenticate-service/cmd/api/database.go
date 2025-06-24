package main

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDB() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=postgres user=root password=root dbname=%s port=5432 sslmode=disable", os.Getenv("DB_NAME"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to postgres database server")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, errors.New("failed to get database instance")
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)

	return DB, nil
}
