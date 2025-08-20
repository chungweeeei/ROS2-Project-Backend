package gateways

import (
	"context"
	"record-service/proto/auth"

	"google.golang.org/grpc"
)

type TestAuthClient struct {
	address string
}

func NewTestAuthClient(address string) *TestAuthClient {
	return &TestAuthClient{address: address}
}

func (tc *TestAuthClient) CheckAuthenticate(ctx context.Context, in *auth.AuthenticateRequest, opts ...grpc.CallOption) (*auth.AuthenticateResponse, error) {
	return &auth.AuthenticateResponse{
		IsAuthenticated: true,
		Email:           "test@test.com",
	}, nil
}
