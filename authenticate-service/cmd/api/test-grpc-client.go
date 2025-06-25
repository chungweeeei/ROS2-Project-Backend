// mocks.go
package main

import (
	"authenticate-service/logs"
	"context"
)

type MockLogServiceClient struct {
	WriteLogFunc func(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error)
}

func (m *MockLogServiceClient) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	if m.WriteLogFunc != nil {
		return m.WriteLogFunc(ctx, req)
	}

	// Default mock response
	return &logs.LogResponse{
		Result: "success",
	}, nil
}
