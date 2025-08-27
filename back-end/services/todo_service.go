package services

import (
	"back-end/database"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoService struct {
	db *gorm.DB
}

// NewTodoService 创建新的Todo服务实例
func NewTodoService() *TodoService {
	return &TodoService{
		db: database.GetDB(),
	}
}

// CreateTodo 创建新的Todo
func (s *TodoService) CreateTodo(title string) (*database.Todo, error) {
	if title == "" {
		return nil, errors.New("todo标题不能为空")
	}

	todo := &database.Todo{
		ID:       uuid.New().String(),
		Title:    title,
		Edit:     false,
		Complete: false,
	}

	if err := s.db.Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// GetAllTodos 获取所有Todo
func (s *TodoService) GetAllTodos() ([]database.Todo, error) {
	var todos []database.Todo
	if err := s.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// GetTodoByID 根据ID获取Todo
func (s *TodoService) GetTodoByID(id string) (*database.Todo, error) {
	var todo database.Todo
	if err := s.db.First(&todo, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Todo未找到")
		}
		return nil, err
	}
	return &todo, nil
}

// UpdateTodo 更新Todo
func (s *TodoService) UpdateTodo(id string, todo *database.Todo) error {
	existingTodo, err := s.GetTodoByID(id)
	if err != nil {
		return err
	}

	// 更新字段
	existingTodo.Title = todo.Title
	existingTodo.Edit = todo.Edit
	existingTodo.Complete = todo.Complete

	if err := s.db.Save(existingTodo).Error; err != nil {
		return err
	}

	return nil
}

// DeleteTodo 删除单个Todo
func (s *TodoService) DeleteTodo(id string) error {
	result := s.db.Delete(&database.Todo{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Todo未找到")
	}
	return nil
}

// DeleteAllTodos 删除所有Todo
func (s *TodoService) DeleteAllTodos() error {
	if err := s.db.Where("1 = 1").Delete(&database.Todo{}).Error; err != nil {
		return err
	}
	return nil
}

// SelectAllTodos 全选所有Todo
func (s *TodoService) SelectAllTodos() error {
	if err := s.db.Model(&database.Todo{}).Where("1 = 1").Update("Complete", true).Error; err != nil {
		return err
	}
	return nil
}

// CancelSelectAllTodos 取消全选所有Todo
func (s *TodoService) CancelSelectAllTodos() error {
	if err := s.db.Model(&database.Todo{}).Where("1 = 1").Update("Complete", false).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCompletedTodos 删除已完成的Todo
func (s *TodoService) DeleteCompletedTodos() error {
	if err := s.db.Where("Complete = ?", true).Delete(&database.Todo{}).Error; err != nil {
		return err
	}
	return nil
}

// GetTodoStats 获取Todo统计信息
func (s *TodoService) GetTodoStats() (map[string]int64, error) {
	var total, completed, pending int64

	// 获取总数
	if err := s.db.Model(&database.Todo{}).Count(&total).Error; err != nil {
		return nil, err
	}

	// 获取已完成数量
	if err := s.db.Model(&database.Todo{}).Where("Complete = ?", true).Count(&completed).Error; err != nil {
		return nil, err
	}

	pending = total - completed

	return map[string]int64{
		"total":     total,
		"completed": completed,
		"pending":   pending,
	}, nil
}
