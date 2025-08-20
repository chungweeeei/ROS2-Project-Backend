package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {

	conn := connectToDB()
	if conn == nil {
		log.Panic("Can not connect to database.")
	}

	return conn
}

func connectToDB() *gorm.DB {

	count := 0
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	for {
		// use gorm connect to postgres database
		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Postgres not yet ready, retrying...")
		} else {
			DB, err := connection.DB()
			if err != nil {
				fmt.Println("Failed connect to database")
				return nil
			}

			// setup connection pool settings
			DB.SetMaxIdleConns(5)
			DB.SetMaxOpenConns(10)
			fmt.Println("Connected to Postgres database successfully")
			return connection
		}

		if count > 10 {
			return nil
		}

		log.Println("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		count++
	}

}
