package gin

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

func RegisterHooks(lifecycle fx.Lifecycle, router *Router) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("CQRSClientの開始 Post:8081 !!!")
				go router.Engine.Run(":8081")
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("CQRSClientの停止 !!!")
				return nil
			},
		},
	)
}
