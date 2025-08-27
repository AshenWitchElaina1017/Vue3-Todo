package main

import (
	"back-end/todolist"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	todo := r.Group("api")
	{
		todo.POST("/add", todolist.AddTodo)
		todo.GET("/get", todolist.GetTodo)
		todo.PUT("/update", todolist.UpdateTodo)
		todo.DELETE("/delete", todolist.DeleteTodo)
		todo.DELETE("/deleteAll", todolist.DeleteAllTodo)
		todo.PUT("/selectAll", todolist.SelectAll)
		todo.PUT("/cancelSelectAll", todolist.CancelSelectAll)
		todo.DELETE("/deleteCompleted", todolist.DeleteCompleted)
	}
	r.Run(":1017")
}
