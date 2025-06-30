package connect

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type queryConnector struct{}

func NewQueryConnector() ServiceConnector {
	return &queryConnector{}
}

func (c *queryConnector) Connect() (*grpc.ClientConn, error) {
	server := "query-service:8083"
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
