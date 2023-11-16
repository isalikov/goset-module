package service

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "goservice/pkg/api/grpc"
)

type grpcError struct {
	Message string
	Service string
}

func (e *grpcError) Error() string {
	return fmt.Sprintf("%d - %s", e.Message, e.Service)
}

func Connect(addr string) (proto.GoServiceClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, &grpcError{
			Message: err.Error(),
			Service: fmt.Sprintf("Service: %v, URL: %v", "goservice", addr),
		}
	}

	return proto.NewGoServiceClient(conn), nil
}
