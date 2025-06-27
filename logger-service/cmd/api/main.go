package main

import (
	"fmt"
	"logger-service/data"
	"net/http"
	"os"

	"gorm.io/gorm"
)

const (
	serverPort = "80"
	gRPCPort   = "50001"
)

type Config struct {
	Repo data.Repository
}

func main() {

	// Step1: connect to the database
	conn, err := connectToDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	// Step2: setup the config
	app := Config{}
	app.setupRepo(conn)

	// Step3: setup the server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: app.routes(),
	}

	// Start gRPC server
	go app.gRPCListen()

	// Step4: start the server
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func (app *Config) setupRepo(conn *gorm.DB) {
	db, err := data.NewPostgresRepository(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	app.Repo = db
}
