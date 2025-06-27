package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/data"
	"logger-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	// This member of your struct is going to be required for pretty much every service you ever write over gRPC
	logs.UnimplementedLogServiceServer
	LogRepo data.Repository
}

// gRPC handler
func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {

	input := req.GetLogEntry() // gets the input from the request

	// write the Log
	logEntry := data.LogEntry{
		Name:    input.Name,
		Level:   input.Level,
		Message: input.Message,
	}

	err := l.LogRepo.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "Failed to write log"}
		return res, err
	}

	// return response
	res := &logs.LogResponse{Result: "logged!"}
	return res, nil
}

// Start listening for gRPC requests
func (app *Config) gRPCListen() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{
		LogRepo: app.Repo,
	})

	log.Printf("gRPC server started on port %s", gRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
