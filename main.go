package main

import "github.com/gin-gonic/gin"

// タスク情報を表す構造体
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// メモリ上にタスクを保持するためのスライス
var tasks []Task

func main() {
	// ルーティングの設定
	r := gin.Default()

	// タスク一覧を取得するエンドポイント（GET）
	r.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, tasks)
	})

	// サーバー
	r.Run(":8080")
}
