package todolist

import (
	"back-end/database"
	"back-end/services"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// TodoController Todo控制器
type TodoController struct {
	service *services.TodoService
}

// NewTodoController 创建新的Todo控制器
func NewTodoController() *TodoController {
	return &TodoController{
		service: services.NewTodoService(),
	}
}

// 全局控制器实例（惰性初始化，确保在数据库初始化之后创建）
var controller *TodoController
var controllerOnce sync.Once

func getController() *TodoController {
	controllerOnce.Do(func() {
		controller = NewTodoController()
	})
	return controller
}

// AddTodo 添加新的Todo
func AddTodo(c *gin.Context) {
	var request struct {
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误", "details": err.Error()})
		return
	}

	todo, err := getController().service.CreateTodo(request.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建Todo失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo添加成功", "todo": todo})
}

// GetTodo 获取所有Todo
func GetTodo(c *gin.Context) {
	todos, err := getController().service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取Todo失败", "details": err.Error()})
		return
	}

	// 获取统计信息
	stats, err := controller.service.GetTodoStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取统计信息失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  map[string]interface{}{"todos": todos},
		"stats": stats,
	})
}

// UpdateTodo 更新Todo
func UpdateTodo(c *gin.Context) {
	var request struct {
		ID       string `json:"id" binding:"required"`
		Title    string `json:"title" binding:"required"`
		Edit     bool   `json:"edit"`
		Complete bool   `json:"Complete"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误", "details": err.Error()})
		return
	}

	todo := &database.Todo{
		ID:       request.ID,
		Title:    request.Title,
		Edit:     request.Edit,
		Complete: request.Complete,
	}

	err := getController().service.UpdateTodo(request.ID, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新Todo失败", "details": err.Error()})
		return
	}

	// 获取更新后的Todo并返回
	updatedTodo, err := getController().service.GetTodoByID(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取更新后的Todo失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo更新成功", "todo": updatedTodo})
}

// DeleteTodo 删除单个Todo
func DeleteTodo(c *gin.Context) {
	var request struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误", "details": err.Error()})
		return
	}

	err := getController().service.DeleteTodo(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除Todo失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo删除成功"})
}

// DeleteAllTodo 删除所有Todo
func DeleteAllTodo(c *gin.Context) {
	err := getController().service.DeleteAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除所有Todo失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "所有Todo已删除"})
}

// SelectAll 全选所有待办事项
func SelectAll(c *gin.Context) {
	err := getController().service.SelectAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "全选Todo失败", "details": err.Error()})
		return
	}

	// 获取更新后的todos
	todos, err := getController().service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取Todo失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "所有Todo已选中", "data": map[string]interface{}{"todos": todos}})
}

// CancelSelectAll 取消全选所有待办事项
func CancelSelectAll(c *gin.Context) {
	err := getController().service.CancelSelectAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "取消全选Todo失败", "details": err.Error()})
		return
	}

	// 获取更新后的todos
	todos, err := getController().service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取Todo失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "取消全选成功", "data": map[string]interface{}{"todos": todos}})
}

// DeleteCompleted 删除已完成的待办事项
func DeleteCompleted(c *gin.Context) {
	err := getController().service.DeleteCompletedTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除已完成Todo失败", "details": err.Error()})
		return
	}

	// 获取更新后的todos
	todos, err := getController().service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取Todo失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "已完成的Todo已删除", "data": map[string]interface{}{"todos": todos}})
}
