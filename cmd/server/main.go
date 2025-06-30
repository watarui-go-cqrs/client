package main

import (
	"client/internal/presentation"

	"go.uber.org/fx"
)

// @Title gRPCサンプルプログラム(CQRSサービス)
// @Version 1.0
// @Description カテゴリ情報と商品情報を管理するAPIサービス
// @TermOfService http://localhost:8081/
// @Contact.name example
// @Contact.url https://www.example.com
// @Contact.email example@example.com
// @Host localhost:8081
// @Basepath /
func main() {
	fx.New(
		presentation.PresenDepend,
	).Run()
}
