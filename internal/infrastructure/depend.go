package infrastructure

import (
	"client/internal/infrastructure/connect"
	"client/internal/infrastructure/provider"

	"go.uber.org/fx"
)

var InfraDepend = fx.Options(
	fx.Provide(
		fx.Annotated{
			Name:   "CommandConnector",
			Target: connect.NewCommandConnector,
		},
		fx.Annotated{
			Name:   "QueryConnector",
			Target: connect.NewQueryConnector,
		},
		fx.Annotated{
			Target: func(
				in struct {
					fx.In
					Connector connect.ServiceConnector `name:"CommandConnector"`
				}) (*provider.CommandClientProvider, error) {
				return provider.NewCommandClientProvider(in.Connector)
			},
		},
		fx.Annotated{
			Target: func(
				in struct {
					fx.In
					Connector connect.ServiceConnector `name:"QueryConnector"`
				}) (*provider.QueryClientProvider, error) {
				return provider.NewQueryClientProvider(in.Connector)
			},
		},
	),
)
