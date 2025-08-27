# Vue TodoList 前端

基于 Vue 3 + TypeScript + Vite 构建的现代化待办事项管理前端应用。

## 🚀 技术栈

- **Vue 3.5** - 使用 Composition API 的响应式框架
- **TypeScript 5.8** - 类型安全的 JavaScript 超集
- **Vite 7.1** - 极速的前端构建工具  
- **Axios 1.11** - 优雅的 HTTP 客户端库

## 📦 项目依赖

### 生产依赖
```json
{
  "axios": "^1.11.0",     // HTTP 请求库
  "nanoid": "^5.1.5",     // 轻量级唯一ID生成器
  "uuid": "^11.1.0",      // UUID 生成器  
  "vue": "^3.5.18"        // Vue 3 框架
}
```

### 开发依赖
```json
{
  "@vitejs/plugin-vue": "^6.0.1",  // Vue 插件
  "@vue/tsconfig": "^0.7.0",       // Vue TypeScript 配置
  "typescript": "~5.8.3",          // TypeScript 编译器
  "vite": "^7.1.2",               // 构建工具
  "vue-tsc": "^3.0.5"             // Vue TypeScript 检查器
}
```

## 🗂️ 项目结构

```
src/
├── api/                    # API 接口层
│   └── todoApi.ts         # Todo API 服务类
├── components/            # Vue 组件
│   ├── TodoNav.vue       # 顶部导航组件
│   ├── TodoMain.vue      # 主体组件(核心逻辑)
│   └── TodoFooter.vue    # 底部组件
├── assets/               # 静态资源
│   └── vue.svg          # Vue logo
├── App.vue              # 根组件
├── main.ts              # 应用入口文件
└── vite-env.d.ts        # Vite 类型声明
```

## ✨ 核心功能

### 📝 Todo 管理
- **添加待办事项** - 输入框回车或点击按钮添加
- **编辑待办事项** - 双击或点击编辑按钮进入编辑模式
- **删除待办事项** - 单个删除或批量删除
- **完成状态切换** - 复选框标记完成状态

### 🔄 批量操作
- **全选功能** - 一键标记所有项目为完成
- **取消全选** - 一键取消所有项目的完成状态
- **删除已完成** - 批量删除所有已完成的项目

### 🗂️ 筛选功能
- **全部** - 显示所有待办事项
- **待办事项** - 仅显示未完成的项目
- **已完成** - 仅显示已完成的项目

### 📱 用户体验优化
- **乐观更新** - 操作立即反映在UI上，提升响应速度
- **错误处理** - 完善的错误提示和回滚机制
- **加载状态** - 清晰的加载指示器
- **响应式设计** - 完美适配移动端设备

## 🎨 设计特点

### 界面设计
- **现代化UI** - 采用卡片式设计语言
- **渐变按钮** - 美观的渐变色彩搭配
- **动画效果** - 流畅的过渡动画
- **状态反馈** - 清晰的视觉状态反馈

### 响应式适配
- **桌面端** (>600px) - 横向布局，完整功能
- **移动端** (≤600px) - 纵向堆叠，触控优化
- **小屏设备** (≤480px) - 进一步压缩布局

## 🔧 开发命令

```bash
# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 构建生产版本
npm run build  

# 预览构建结果
npm run preview
```

## 🌐 API 集成

### 接口配置
```typescript
const API_BASE_URL = 'http://localhost:1017/api'
```

### 数据类型
```typescript
interface Todo {
  id: string
  title: string  
  edit: boolean
  Complete: boolean
}
```

### 服务方法
- `getTodos()` - 获取所有待办事项
- `addTodo(title)` - 添加新的待办事项
- `updateTodo(todo)` - 更新待办事项
- `deleteTodo(todo)` - 删除待办事项
- `selectAllTodos()` - 全选操作
- `cancelSelectAllTodos()` - 取消全选
- `deleteCompletedTodos()` - 删除已完成项目

## 🔄 状态管理

采用 Vue 3 Composition API 进行状态管理：

```typescript
// 响应式状态
const todos = reactive<Todo[]>([])      // 待办事项列表
const loading = ref(false)              // 全局加载状态
const addingTodo = ref(false)           // 添加状态
const error = ref('')                   // 错误信息
const currentFilter = ref('all')        // 当前过滤器
```

## ⚡ 性能优化

### 乐观更新策略
- **添加操作** - 立即显示在列表中，API失败后回滚
- **删除操作** - 立即从列表移除，API失败后恢复
- **状态切换** - 立即更新UI状态，API失败后回滚

### 用户体验
- **防重复操作** - 按钮禁用防止重复提交
- **输入验证** - 前端验证提升响应速度
- **错误恢复** - 自动恢复到操作前状态

## 🎯 开发最佳实践

### TypeScript 类型安全
- 完整的类型定义
- 接口类型约束
- 编译时类型检查

### 组件设计原则
- **单一职责** - 每个组件职责明确
- **组合优于继承** - 使用Composition API
- **响应式编程** - 充分利用Vue 3响应式系统

### 错误处理
- **统一错误处理** - 集中的错误处理逻辑
- **用户友好提示** - 清晰的错误信息展示
- **优雅降级** - 网络错误时的备用方案

## 🚀 部署配置

### 构建优化
```bash
# 构建生产版本
npm run build

# 构建产物位于 dist/ 目录
# 可直接部署到静态服务器
```

### 环境变量
可通过 `.env` 文件配置不同环境的API地址：

```bash
VITE_API_BASE_URL=http://localhost:1017/api
```

## 📱 移动端适配

### 触控优化
- 合适的按钮大小(最小44px)
- 触控友好的间距设计
- 防误触设计

### 布局适配
- Flexbox 响应式布局
- 媒体查询断点设计
- 移动端专用交互模式

## 🔍 调试和测试

### 开发工具
- Vue Devtools 扩展支持
- TypeScript 类型检查
- Vite HMR 热更新

### 错误监控
- Console 错误日志
- 网络请求状态监控
- 用户操作反馈

## 📄 许可证

MIT License
