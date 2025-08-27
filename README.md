# Vue + Go TodoList 应用

一个现代化的全栈待办事项管理应用，采用 Vue3 前端 + Go 后端的架构设计。

## 🚀 技术栈

### 前端
- **Vue 3** - 响应式前端框架
- **TypeScript** - 类型安全的JavaScript
- **Vite** - 快速构建工具
- **Axios** - HTTP 客户端

### 后端  
- **Go** - 高性能后端语言
- **Gin** - Web 框架
- **GORM** - ORM 框架
- **SQLite** - 轻量级数据库

## 📁 项目结构

```
vue-todo/
├── src/                    # 前端源码
│   ├── api/               # API 接口层
│   ├── components/        # Vue 组件
│   ├── assets/           # 静态资源
│   └── main.ts           # 应用入口
├── back-end/              # 后端源码
│   ├── database/         # 数据库层
│   ├── services/         # 业务逻辑层  
│   ├── todolist/         # 控制器层
│   └── main.go           # 服务入口
├── public/               # 静态文件
└── package.json          # 前端依赖配置
```

## ✨ 功能特性

- ✅ 添加/删除/编辑待办事项
- ✅ 标记完成状态
- ✅ 批量操作（全选/取消全选）
- ✅ 状态筛选（全部/待办/已完成）
- ✅ 实时数据同步
- ✅ 乐观更新策略
- ✅ 响应式设计
- ✅ 错误处理机制

## 🛠️ 快速开始

### 前端启动

详见：[前端 README](./src/README.md)

### 后端启动

详见：[后端 README](./back-end/README.md)

## 🌐 API 接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/get` | 获取所有待办事项 |
| POST | `/api/add` | 添加待办事项 |
| PUT | `/api/update` | 更新待办事项 |
| DELETE | `/api/delete` | 删除单个待办事项 |
| DELETE | `/api/deleteAll` | 删除所有待办事项 |
| PUT | `/api/selectAll` | 全选待办事项 |
| PUT | `/api/cancelSelectAll` | 取消全选 |
| DELETE | `/api/deleteCompleted` | 删除已完成项目 |
| GET | `/health` | 健康检查 |

## 🔧 开发配置

### 端口配置
- 前端开发服务器: `http://localhost:5173`
- 后端API服务器: `http://localhost:1017`

### CORS 配置
后端已配置 CORS，允许来自前端开发服务器的请求。

## 📱 响应式设计

应用完全适配移动端设备，提供优秀的移动端用户体验。

## 🚀 部署说明

1. 前端构建：`npm run build`
2. 后端编译：`go build -o todo-server`
3. 部署构建产物到服务器

## 📄 许可证

MIT License

## 🤝 贡献

欢迎提交 Pull Request 来改进项目！