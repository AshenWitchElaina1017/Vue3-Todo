package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Todo 数据库模型
type Todo struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"not null"`
	Edit     bool   `json:"edit" gorm:"default:false"`
	Complete bool   `json:"Complete" gorm:"default:false"`
}

// InitDatabase 初始化数据库连接
func InitDatabase() {
	var err error
	// 连接SQLite数据库
	DB, err = gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移数据库表
	err = DB.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully")
}

// GetDB 获取数据库连接实例
func GetDB() *gorm.DB {
	return DB
}
