package main

import (
	"authenticate-service/data"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

const serverPort = "80"

type Config struct {
	DB     *gorm.DB
	Models data.Models
}

func main() {

	// Step1: connect to the database
	conn, err := connectToDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}

	fmt.Println("Connected to database successfully")

	// Step2: setup the config
	models, err := data.New(conn)
	if err != nil {
		fmt.Println(err)
	}

	app := Config{
		DB:     conn,
		Models: models,
	}

	// Step3: setup the server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: app.routes(),
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
