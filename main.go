// Ginフレームワークの挙動を確認

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	// ルートパスにアクセスした際にHello Worldを返す
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello World")
	})

	// ポート8080でサーバを起動
	r.Run(":8080")
}
