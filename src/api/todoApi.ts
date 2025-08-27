import axios from 'axios'

// 后端API基础URL
const API_BASE_URL = 'http://localhost:1017/api'

// 创建axios实例
const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Todo数据类型定义
export interface Todo {
  id: string
  title: string
  edit: boolean
  Complete: boolean  // 注意：后端使用大写C
}

export interface TodoListResponse {
  data: {
    todos: Todo[]
  }
}

export interface ApiResponse {
  message?: string
  todo?: Todo
  data?: {
    todos: Todo[]
  }
  error?: string
}

// API服务类
export class TodoApiService {
  
  // 获取所有待办事项
  static async getTodos(): Promise<Todo[]> {
    try {
      const response = await api.get<ApiResponse>('/get')
      return response.data.data?.todos || []
    } catch (error) {
      console.error('获取待办事项失败:', error)
      throw new Error('获取待办事项失败')
    }
  }

  // 添加待办事项
  static async addTodo(title: string): Promise<Todo> {
    try {
      const response = await api.post<ApiResponse>('/add', {
        title: title
      })
      
      if (response.data.todo) {
        return response.data.todo
      }
      throw new Error('添加失败')
    } catch (error) {
      console.error('添加待办事项失败:', error)
      throw new Error('添加待办事项失败')
    }
  }

  // 更新待办事项
  static async updateTodo(todo: Todo): Promise<Todo> {
    try {
      const response = await api.put<ApiResponse>('/update', todo)
      
      if (response.data.todo) {
        return response.data.todo
      }
      throw new Error('更新失败')
    } catch (error) {
      console.error('更新待办事项失败:', error)
      throw new Error('更新待办事项失败')
    }
  }

  // 删除待办事项
  static async deleteTodo(todo: Todo): Promise<void> {
    try {
      await api.delete('/delete', {
        data: todo
      })
    } catch (error) {
      console.error('删除待办事项失败:', error)
      throw new Error('删除待办事项失败')
    }
  }

  // 删除所有待办事项
  static async deleteAllTodos(): Promise<void> {
    try {
      await api.delete('/deleteAll')
    } catch (error) {
      console.error('删除所有待办事项失败:', error)
      throw new Error('删除所有待办事项失败')
    }
  }

  // 全选所有待办事项
  static async selectAllTodos(): Promise<Todo[]> {
    try {
      const response = await api.put<ApiResponse>('/selectAll')
      return response.data.data?.todos || []
    } catch (error) {
      console.error('全选待办事项失败:', error)
      throw new Error('全选待办事项失败')
    }
  }

  // 取消全选
  static async cancelSelectAllTodos(): Promise<Todo[]> {
    try {
      const response = await api.put<ApiResponse>('/cancelSelectAll')
      return response.data.data?.todos || []
    } catch (error) {
      console.error('取消全选失败:', error)
      throw new Error('取消全选失败')
    }
  }

  // 删除已完成的待办事项
  static async deleteCompletedTodos(): Promise<Todo[]> {
    try {
      const response = await api.delete<ApiResponse>('/deleteCompleted')
      return response.data.data?.todos || []
    } catch (error) {
      console.error('删除已完成待办事项失败:', error)
      throw new Error('删除已完成待办事项失败')
    }
  }
}
