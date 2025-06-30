package presentation

import (
	"client/internal/infrastructure"
	"client/internal/presentation/gin"
	"client/internal/presentation/gin/handler"
	"client/internal/presentation/gin/helper"

	"go.uber.org/fx"
)

var PresenDepend = fx.Options(
	infrastructure.InfraDepend,
	fx.Provide(
		helper.NewCommandHelper,
		helper.NewQueryHelper,
		handler.NewCommandHandler,
		handler.NewQueryHandler,
		gin.NewRouter,
	),
	fx.Invoke(gin.RegisterHooks),
)
