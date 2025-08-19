package main

import (
	"auth-service/proto/auth"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type AuthenticateServer struct {
	auth.UnimplementedAuthenticateServiceServer
}

// gRPC handler for authentication
func (a *AuthenticateServer) CheckAuthenticate(ctx context.Context, req *auth.AuthenticateRequest) (resp *auth.AuthenticateResponse, err error) {

	// check token from request
	token := req.GetToken()

	// verify the token
	email, err := verifyToken(token)
	if err != nil {
		return &auth.AuthenticateResponse{
			IsAuthenticated: false,
			Email:           "",
		}, err
	}

	// if token is valid, return authenticated response
	return &auth.AuthenticateResponse{
		IsAuthenticated: true,
		Email:           email,
	}, nil
}

func (app *Config) gRPCListener() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		app.ErrorLog.Fatalf("Failed to listen on port %s: %v", gRPCPort, err)
	}

	s := grpc.NewServer()

	auth.RegisterAuthenticateServiceServer(s, &AuthenticateServer{})

	app.InfoLog.Printf("Start gRPC server listening on port %s", gRPCPort)
	if err := s.Serve(lis); err != nil {
		app.ErrorLog.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
