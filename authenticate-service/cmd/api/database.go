package main

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDB() (*gorm.DB, error) {

	// dsn := fmt.Sprintf("host=postgres user=root password=root dbname=%s port=5432 sslmode=disable", os.Getenv("DB_NAME"))
	dsn := "host=localhost user=root password=root dbname=go_db port=5432 sslmode=disable"

	var counts int = 0
	for {
		DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Postgres not yet ready, retrying...")
			counts++
		} else {
			sqlDB, err := DB.DB()
			if err != nil {
				return nil, errors.New("failed to get database instance")
			}

			// Set connection pool settings
			sqlDB.SetMaxIdleConns(5)
			sqlDB.SetMaxOpenConns(10)
			fmt.Println("Connected to Postgres database successfully")
			return DB, nil
		}

		if counts > 5 {
			return nil, errors.New("failed to connect to postgres database server")
		}

		fmt.Println("Backing off for 2 seconds before retrying...")
		time.Sleep(2 * time.Second)
	}
}
