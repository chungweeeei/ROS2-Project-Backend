package gateways

import (
	"context"
	"record-service/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient interface {
	CheckAuthenticate(ctx context.Context, in *auth.AuthenticateRequest, opts ...grpc.CallOption) (*auth.AuthenticateResponse, error)
}

type AuthClient struct {
	address string
}

func NewAuthClient(address string) *AuthClient {
	return &AuthClient{address: address}
}

func (ac *AuthClient) CheckAuthenticate(ctx context.Context, in *auth.AuthenticateRequest, opts ...grpc.CallOption) (*auth.AuthenticateResponse, error) {

	conn, err := grpc.NewClient(ac.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := auth.NewAuthenticateServiceClient(conn)
	return client.CheckAuthenticate(ctx, in, opts...)

}
