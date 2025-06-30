package provider

import (
	"client/internal/infrastructure/connect"

	"github.com/watarui-go-cqrs/pb/pb"
)

type CommandClientProvider struct {
	Category pb.CategoryCommandClient
	Product  pb.ProductCommandClient
}

func NewCommandClientProvider(connector connect.ServiceConnector) (*CommandClientProvider, error) {
	conn, err := connector.Connect()
	if err != nil {
		return nil, err
	}
	return &CommandClientProvider{
		Category: pb.NewCategoryCommandClient(conn),
		Product:  pb.NewProductCommandClient(conn),
	}, nil
}
