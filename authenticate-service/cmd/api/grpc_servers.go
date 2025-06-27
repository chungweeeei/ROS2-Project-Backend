package main

import (
	"authenticate-service/proto/auth"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type AuthenticateServer struct {
	auth.UnimplementedAuthenticateServiceServer
}

// gRPC handler for authentication
func (a *AuthenticateServer) CheckAuthenticate(ctx context.Context, req *auth.AuthenticateRequest) (*auth.AuthenticateResponse, error) {

	// check token
	token := req.GetToken()

	err := verifyToken(token)
	if err != nil {
		return &auth.AuthenticateResponse{
			IsAuthenticated: false,
			Message:         "Token is not valid",
		}, err
	}

	return &auth.AuthenticateResponse{
		IsAuthenticated: true,
		Message:         "Token is valid",
	}, nil
}

// Start listening for gRPC requests
func (app *Config) gRPCListen() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	auth.RegisterAuthenticateServiceServer(s, &AuthenticateServer{})

	log.Printf("gRPC server started on port %s", gRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
