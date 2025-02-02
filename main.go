package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	// タスクを追加するエンドポイント（POST）
	r.POST("/tasks", func(ctx *gin.Context) {
		var newTask Task
		if err := ctx.ShouldBindJSON(&newTask); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		newTask.ID = len(tasks) + 1
		tasks = append(tasks, newTask)
		ctx.JSON(201, newTask)
	})

	// タスクを更新するエンドポイント（PUT）
	r.PUT("/tasks/:id", func(ctx *gin.Context) {
		// idを文字列→数値に変換
		id := ctx.Param("id")
		taskID, err := strconv.Atoi(id)

		// IDが数値でない場合
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		var updateTask Task // 更新するタスク

		// リクエストボディをJSONにバインド
		if err := ctx.ShouldBindJSON(&updateTask); err != nil {
			ctx.JSON(400, gin.H{"error": "Invaild JSON"})
			return
		}

		// 該当するタスクを検索して更新
		for i, t := range tasks {
			if t.ID == taskID {
				tasks[i] = updateTask
				ctx.JSON(200, tasks[i])
				return
			}
		}

		// 該当するタスクが見つからない場合
		ctx.JSON(404, gin.H{"error": "Task not found"})

	})

	// タスクを削除するエンドポイント（DELETE）
	r.DELETE("/tasks/:id", func(ctx *gin.Context) {
		// パスパラメータからIDを取得
		id := ctx.Param("id")
		taskID, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		// 該当するタスクを検索して削除
		found := false
		newTasks := []Task{} // 削除後のタスクリストを作成
		for _, t := range tasks {
			if t.ID == taskID {
				found = true // 該当タスクが見つかった
				continue     // 削除するタスクはスキップ
			}
			newTasks = append(newTasks, t) // 他のタスクは新しいリストに追加
		}

		if !found {
			// 該当するタスクが見つからない場合
			ctx.JSON(404, gin.H{"error": "Task not found"})
			return
		}

		// タスクリストを更新
		tasks = newTasks
		ctx.JSON(204, nil) // 成功（204 No Content）
	})

	// サーバー
	r.Run(":8080")
}
