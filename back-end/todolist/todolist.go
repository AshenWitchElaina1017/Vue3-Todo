package todolist

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Todo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Edit     bool   `json:"edit"`
	Complete bool   `json:"Complete"`
}
type TodoList struct {
	Todos []Todo `json:"todos"`
}

// 内存存储（生产环境中应该使用数据库）
var todoList = TodoList{
	Todos: []Todo{},
}

func AddTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 生成UUID
	todo.ID = uuid.New().String()
	todo.Edit = false
	todo.Complete = false

	// 添加到todo列表
	todoList.Todos = append(todoList.Todos, todo)

	c.JSON(http.StatusOK, gin.H{"message": "Todo添加成功", "todo": todo})
}
func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": todoList})
}
func UpdateTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查找并更新todo
	for i, existingTodo := range todoList.Todos {
		if existingTodo.ID == todo.ID {
			todoList.Todos[i] = todo
			c.JSON(http.StatusOK, gin.H{"message": "Todo更新成功", "todo": todo})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo未找到"})
}
func DeleteTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查找并删除todo
	for i, existingTodo := range todoList.Todos {
		if existingTodo.ID == todo.ID {
			todoList.Todos = append(todoList.Todos[:i], todoList.Todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo删除成功"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo未找到"})
}
func DeleteAllTodo(c *gin.Context) {
	// 清空所有todos
	todoList.Todos = []Todo{}

	c.JSON(http.StatusOK, gin.H{"message": "所有Todo已删除"})
}

// SelectAll 全选所有待办事项
func SelectAll(c *gin.Context) {
	for i := range todoList.Todos {
		todoList.Todos[i].Complete = true
	}

	c.JSON(http.StatusOK, gin.H{"message": "所有Todo已选中", "data": todoList})
}

// CancelSelectAll 取消全选所有待办事项
func CancelSelectAll(c *gin.Context) {
	for i := range todoList.Todos {
		todoList.Todos[i].Complete = false
	}

	c.JSON(http.StatusOK, gin.H{"message": "取消全选成功", "data": todoList})
}

// DeleteCompleted 删除已完成的待办事项
func DeleteCompleted(c *gin.Context) {
	var remainingTodos []Todo
	for _, todo := range todoList.Todos {
		if !todo.Complete {
			remainingTodos = append(remainingTodos, todo)
		}
	}

	todoList.Todos = remainingTodos

	c.JSON(http.StatusOK, gin.H{"message": "已完成的Todo已删除", "data": todoList})
}
