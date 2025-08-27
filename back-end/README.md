# Go TodoList 后端

基于 Go + Gin + GORM + SQLite 构建的高性能待办事项管理后端服务。

## 🚀 技术栈

- **Go 1.24.4** - 高性能后端编程语言
- **Gin 1.10.1** - 轻量级Web框架
- **GORM 1.25.11** - 功能强大的ORM框架
- **SQLite** - 轻量级嵌入式数据库
- **UUID** - 唯一标识符生成

## 📦 项目依赖

### 核心依赖
```go
require (
    github.com/gin-contrib/cors v1.7.6    // CORS 中间件
    github.com/gin-gonic/gin v1.10.1      // Web 框架
    github.com/google/uuid v1.6.0         // UUID 生成器
    gorm.io/driver/sqlite v1.5.4          // SQLite 驱动
    gorm.io/gorm v1.25.11                 // ORM 框架
)
```

## 🗂️ 项目架构

```
back-end/
├── database/              # 数据库层
│   └── database.go       # 数据库连接和模型定义
├── services/             # 业务逻辑层  
│   └── todo_service.go   # Todo 服务实现
├── todolist/             # 控制器层
│   └── todolist.go       # HTTP 请求处理
├── main.go              # 应用程序入口
├── go.mod               # Go 模块定义
├── go.sum               # 依赖版本锁定
└── todolist.db          # SQLite 数据库文件
```

## 🏗️ 分层架构设计

### 🗄️ 数据库层 (Database Layer)
**文件**: `database/database.go`

负责数据库连接管理和数据模型定义：

```go
type Todo struct {
    ID       string `json:"id" gorm:"primaryKey"`
    Title    string `json:"title" gorm:"not null"`  
    Edit     bool   `json:"edit" gorm:"default:false"`
    Complete bool   `json:"Complete" gorm:"default:false"`
}
```

**功能特性**:
- SQLite 数据库初始化
- 自动数据表迁移
- 数据库连接池管理
- 结构化日志记录

### 🔧 服务层 (Service Layer)  
**文件**: `services/todo_service.go`

封装业务逻辑和数据操作：

```go
type TodoService struct {
    db *gorm.DB
}
```

**核心方法**:
- `CreateTodo(title)` - 创建新的待办事项
- `GetAllTodos()` - 获取所有待办事项
- `GetTodoByID(id)` - 根据ID获取待办事项
- `UpdateTodo(id, todo)` - 更新待办事项
- `DeleteTodo(id)` - 删除单个待办事项
- `SelectAllTodos()` - 批量标记为完成
- `CancelSelectAllTodos()` - 批量取消完成
- `DeleteCompletedTodos()` - 删除已完成项目
- `GetTodoStats()` - 获取统计信息

### 🌐 控制器层 (Controller Layer)
**文件**: `todolist/todolist.go`

处理HTTP请求和响应：

```go
type TodoController struct {
    service *services.TodoService
}
```

**API 端点**:
- `POST /api/add` - 添加待办事项
- `GET /api/get` - 获取待办事项列表
- `PUT /api/update` - 更新待办事项
- `DELETE /api/delete` - 删除单个待办事项
- `DELETE /api/deleteAll` - 删除所有待办事项
- `PUT /api/selectAll` - 全选操作
- `PUT /api/cancelSelectAll` - 取消全选
- `DELETE /api/deleteCompleted` - 删除已完成项目

## 🔧 快速开始

### 环境要求
- Go 1.20+ 
- Git

### 安装和运行

```bash
# 克隆项目
git clone <项目地址>
cd vue-todo/back-end

# 安装依赖
go mod download

# 运行开发服务器
go run main.go

# 或者编译后运行
go build -o todo-server
./todo-server
```

### 服务启动信息
```
Todo API服务启动在端口 :1017
Database connected and migrated successfully
```

## 🌐 API 接口文档

### 1. 添加待办事项
```http
POST /api/add
Content-Type: application/json

{
    "title": "学习Go语言"
}
```

**响应**:
```json
{
    "message": "Todo添加成功",
    "todo": {
        "id": "uuid-string",
        "title": "学习Go语言", 
        "edit": false,
        "Complete": false
    }
}
```

### 2. 获取待办事项列表
```http
GET /api/get
```

**响应**:
```json
{
    "data": {
        "todos": [...]
    },
    "stats": {
        "total": 5,
        "completed": 2, 
        "pending": 3
    }
}
```

### 3. 更新待办事项
```http
PUT /api/update
Content-Type: application/json

{
    "id": "uuid-string",
    "title": "更新后的标题",
    "edit": false,
    "Complete": true
}
```

### 4. 删除待办事项
```http
DELETE /api/delete
Content-Type: application/json

{
    "id": "uuid-string"
}
```

### 5. 批量操作

**全选**:
```http
PUT /api/selectAll
```

**取消全选**:
```http  
PUT /api/cancelSelectAll
```

**删除已完成**:
```http
DELETE /api/deleteCompleted
```

**删除所有**:
```http
DELETE /api/deleteAll
```

### 6. 健康检查
```http
GET /health
```

**响应**:
```json
{
    "status": "ok",
    "message": "Todo API服务运行正常"
}
```

## 🔒 CORS 配置

支持跨域请求，允许来自前端开发服务器的请求：

```go
config := cors.DefaultConfig()
config.AllowOrigins = []string{
    "http://localhost:5173", // Vite 开发服务器
    "http://localhost:3000"  // 备用端口
}
config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
```

## 🗄️ 数据库设计

### Todo 表结构

| 字段 | 类型 | 约束 | 描述 |
|------|------|------|------|
| id | STRING | PRIMARY KEY | 唯一标识符 |
| title | STRING | NOT NULL | 待办事项标题 |
| edit | BOOLEAN | DEFAULT false | 是否处于编辑状态 |
| Complete | BOOLEAN | DEFAULT false | 是否已完成 |

### 数据库特性
- **自动迁移** - 应用启动时自动创建/更新表结构
- **事务支持** - 确保数据一致性
- **索引优化** - 主键索引提升查询性能
- **数据持久化** - SQLite 文件持久化存储

## ⚡ 性能优化

### 数据库优化
- **连接池管理** - 复用数据库连接
- **预编译语句** - GORM 自动优化查询
- **批量操作优化** - 减少数据库交互次数

### 并发处理
- **Gin 协程池** - 高并发请求处理
- **线程安全** - 服务层单例模式设计
- **惰性初始化** - 使用 sync.Once 确保线程安全

### 内存管理
- **对象复用** - 减少GC压力
- **结构体优化** - 合理的字段布局

## 🛡️ 错误处理

### 统一错误响应格式
```json
{
    "error": "错误描述",
    "details": "详细错误信息"  
}
```

### 错误类型
- **400 Bad Request** - 请求参数错误
- **404 Not Found** - 资源未找到  
- **500 Internal Server Error** - 服务器内部错误

### 错误处理策略
- **输入验证** - 请求参数合法性检查
- **数据库错误** - GORM 错误统一处理
- **业务逻辑错误** - 自定义错误类型

## 🔍 日志监控

### 日志级别
- **Info** - 服务启动、数据库连接等信息
- **Error** - 错误信息记录
- **Debug** - 开发调试信息

### 监控指标
- API 请求响应时间
- 数据库查询性能
- 错误率统计

## 🚀 部署配置

### 生产环境构建
```bash
# 构建二进制文件
go build -ldflags="-w -s" -o todo-server

# 运行服务
./todo-server
```

### Docker 部署
```dockerfile
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o todo-server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/todo-server .
EXPOSE 1017
CMD ["./todo-server"]
```

### 服务器配置
- **端口**: 1017
- **数据库文件**: `todolist.db` (自动创建)
- **日志**: 控制台输出

## 🧪 测试

### 功能测试
```bash
# 健康检查
curl http://localhost:1017/health

# 获取待办事项
curl http://localhost:1017/api/get

# 添加待办事项
curl -X POST http://localhost:1017/api/add \
  -H "Content-Type: application/json" \
  -d '{"title":"测试项目"}'
```

### 性能测试
可使用 `ab` 或 `wrk` 进行API性能测试。

## 🔧 开发工具

### 推荐IDE
- **VS Code** + Go扩展
- **GoLand** (JetBrains)

### 代码质量
- `go fmt` - 代码格式化
- `go vet` - 静态分析检查
- `golint` - 代码风格检查

### 调试
- **Delve** - Go 调试器
- **pprof** - 性能分析工具

## 📈 扩展计划

### 功能扩展
- [ ] 用户认证和授权
- [ ] 待办事项分类和标签  
- [ ] 到期时间和提醒
- [ ] 文件附件支持

### 技术升级
- [ ] Redis 缓存集成
- [ ] PostgreSQL 数据库支持
- [ ] API 限流和监控
- [ ] 微服务架构重构

## 📄 许可证

MIT License