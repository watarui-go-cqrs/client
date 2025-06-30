package provider

import (
	"client/internal/infrastructure/connect"

	"github.com/watarui-go-cqrs/pb/pb"
)

type QueryClientProvider struct {
	Category pb.CategoryQueryClient
	Product  pb.ProductQueryClient
}

func NewQueryClientProvider(connector connect.ServiceConnector) (*QueryClientProvider, error) {
	conn, err := connector.Connect()
	if err != nil {
		return nil, err
	}
	return &QueryClientProvider{
		Category: pb.NewCategoryQueryClient(conn),
		Product:  pb.NewProductQueryClient(conn),
	}, nil
}
