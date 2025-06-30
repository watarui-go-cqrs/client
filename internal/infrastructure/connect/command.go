package connect

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type commandConnector struct{}

func NewCommandConnector() ServiceConnector {
	return &commandConnector{}
}

func (c *commandConnector) Connect() (*grpc.ClientConn, error) {
	server := "command-service:8082"
	conn, err := grpc.Dial(
		server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
