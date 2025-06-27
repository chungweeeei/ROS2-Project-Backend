package main

import (
	"authenticate-service/data"
	"fmt"
	"net/http"
	"os"

	"gorm.io/gorm"
)

const (
	serverPort = "80"
	gRPCPort   = "50001"
)

type Clients struct {
	LogHTTPClient *http.Client
	LoggRPCClient LogServiceClient
}

type Config struct {
	Repo    data.Repository
	Clients Clients
}

func main() {

	// Step1: connect to the database
	conn, err := connectToDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	logClient := NewgRPCLogClient("logger-service:50001")

	// Step2: setup the config
	app := Config{
		Clients: Clients{
			LogHTTPClient: &http.Client{},
			LoggRPCClient: logClient,
		},
	}
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
