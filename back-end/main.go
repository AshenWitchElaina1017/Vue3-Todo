package main

import (
	"back-end/database"
	"back-end/todolist"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.InitDatabase()

	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 配置CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:3000"} // 允许前端端口
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// API路由组
	api := r.Group("/api")
	{
		api.POST("/add", todolist.AddTodo)
		api.GET("/get", todolist.GetTodo)
		api.PUT("/update", todolist.UpdateTodo)
		api.DELETE("/delete", todolist.DeleteTodo)
		api.DELETE("/deleteAll", todolist.DeleteAllTodo)
		api.PUT("/selectAll", todolist.SelectAll)
		api.PUT("/cancelSelectAll", todolist.CancelSelectAll)
		api.DELETE("/deleteCompleted", todolist.DeleteCompleted)
	}

	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Todo API服务运行正常",
		})
	})

	log.Println("Todo API服务启动在端口 :1017")
	if err := r.Run(":1017"); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
