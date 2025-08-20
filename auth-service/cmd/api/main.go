package main

import (
	"auth-service/data"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	serverPort = "80"
	gRPCPort   = "50001"
)

func main() {
	// init database engine
	db := initDB()

	// init logger instance
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := Config{
		DB:            db,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		Models:        data.New(db),
		ErrorChan:     make(chan error),
		ErrorDoneChan: make(chan bool),
	}

	// listen for signal
	go app.listenForShutdown()

	// listen for errors
	go app.listenForErrors()

	// gRPC listener
	go app.gRPCListener()

	// execute api server
	app.serve()
}

func (app *Config) serve() {

	// start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: app.routes(),
	}

	app.InfoLog.Println("Starting auth service...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func (app *Config) listenForErrors() {
	for {
		select {
		case err := <-app.ErrorChan:
			app.ErrorLog.Println(err)
		case <-app.ErrorDoneChan:
			// close the goroutine
			return
		}
	}
}

func (app *Config) listenForShutdown() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {

	// perform any cleanup tasks
	app.InfoLog.Println("Would run cleanup tasks...")

	// send true via done chan
	app.ErrorDoneChan <- true

	// shutdown
	app.InfoLog.Println("closing channels and shutting down application...")
	close(app.ErrorChan)
	close(app.ErrorDoneChan)
}
