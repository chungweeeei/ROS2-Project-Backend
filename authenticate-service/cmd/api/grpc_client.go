package main

import (
	"authenticate-service/proto/logs"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// LogServiceClient interface for dependency injection
type LogServiceClient interface {
	WriteLog(ctx context.Context, in *logs.LogRequest, opts ...grpc.CallOption) (*logs.LogResponse, error)
}

// gRPCLogClient implements LogServiceClient
type gRPCLogClient struct {
	address string
}

func NewgRPCLogClient(address string) *gRPCLogClient {
	return &gRPCLogClient{
		address: address,
	}
}

func (g *gRPCLogClient) WriteLog(ctx context.Context, in *logs.LogRequest, opts ...grpc.CallOption) (*logs.LogResponse, error) {
	conn, err := grpc.NewClient(g.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := logs.NewLogServiceClient(conn)
	return client.WriteLog(ctx, in, opts...)
}
