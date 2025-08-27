<template>
    <div class="todo-main">
        <div class="input-section">
            <div class="input-group">
                <input 
                    type="text" 
                    v-model="todo" 
                    @keyup.enter="addTodo" 
                    placeholder="请输入待办事项"
                    class="todo-input"
                >
                <button @click="addTodo" class="add-btn primary">
                    <span class="btn-icon">+</span>
                    添加
                </button>
            </div>
            <div class="action-buttons" v-if="todos.length>0">
                <div class="button-group">
                    <button @click="selectAll" class="action-btn secondary">
                        <span class="btn-icon">✓</span>
                        全选
                    </button>
                    <button @click="cancelSelectAll" class="action-btn secondary">
                        <span class="btn-icon">✗</span>
                        取消全选
                    </button>
                </div>
                <button @click="deleteComplete" class="action-btn danger">
                    <span class="btn-icon">🗑</span>
                    删除已完成
                </button>
            </div>
        </div>
        <div class="todo-list-box" v-if="todos.length>0">
            <button 
                class="todo-list-box-btn" 
                :class="{ active: currentFilter === 'all' }"
                @click="setFilter('all')"
            >全部<span>{{todos.length}}</span></button>
            <button 
                class="todo-list-box-btn"
                :class="{ active: currentFilter === 'pending' }"
                @click="setFilter('pending')"
            >待办事项<span>{{todos.filter(item=>!item.isComplete).length}}</span></button>
            <button 
                class="todo-list-box-btn"
                :class="{ active: currentFilter === 'completed' }"
                @click="setFilter('completed')"
            >已完成<span>{{todos.filter(item=>item.isComplete).length}}</span></button>
        </div>
        <ul class="todo-list">
            <li v-for="item in filteredTodos" :key="item.id" class="todo-item" :class="{ completed: item.isComplete }">
                <div class="todo-content">
                    <input 
                        type="checkbox" 
                        v-model="item.isComplete" 
                        class="todo-checkbox"
                    >
                    <input 
                        v-if="item.edit" 
                        type="text" 
                        v-model="item.title"
                        class="edit-input"
                    >
                    <span v-else class="todo-text" :class="{ 'completed-text': item.isComplete }">
                        {{ item.title }}
                    </span>
                </div>
                <div class="todo-actions">
                    <button @click="todoDelete(item.id)" class="delete-btn">删除</button>
                    <button 
                        v-if="!item.edit" 
                        @click="todoEdit(item.id)" 
                        class="edit-btn"
                    >编辑</button>
                    <button 
                        v-else 
                        @click="todoComplete(item.id)" 
                        class="complete-btn"
                    >完成</button>
                </div>
            </li>
        </ul>
        <div v-if="filteredTodos.length === 0" class="empty-state">
            <template v-if="todos.length === 0">
                暂无待办事项，添加一个开始吧！
            </template>
            <template v-else-if="currentFilter === 'pending'">
                暂无待办事项，所有任务都已完成！
            </template>
            <template v-else-if="currentFilter === 'completed'">
                暂无已完成的任务
            </template>
        </div>
    </div>
</template>

<script setup>
import { ref,reactive,watch,computed } from 'vue'
import { nanoid } from 'nanoid'
let todo=ref("")
let todos=reactive(JSON.parse(localStorage.getItem("todos"))||[])
let currentFilter=ref('all') // 当前过滤状态：all, pending, completed

// 计算过滤后的任务列表
const filteredTodos = computed(() => {
    switch (currentFilter.value) {
        case 'pending':
            return todos.filter(item => !item.isComplete)
        case 'completed':
            return todos.filter(item => item.isComplete)
        case 'all':
        default:
            return todos
    }
})

function addTodo(){
    let title=todo.value.trim()
    if(title!==""){
        todos.push({id:nanoid(),title:title,edit:false,isComplete:false})
        localStorage.setItem("todos",JSON.stringify(todos))
        todo.value=""
    }
}
function todoDelete(id){
    todos.splice(todos.findIndex(item=>item.id===id),1)
    localStorage.setItem("todos",JSON.stringify(todos))
}
function todoEdit(id){
    todos.find(item=>item.id===id).edit=true
    localStorage.setItem("todos",JSON.stringify(todos))
}
function todoComplete(id){
    todos.find(item=>item.id===id).edit=false
    localStorage.setItem("todos",JSON.stringify(todos))
}
function deleteComplete(){
    const remain = todos.filter(item => !item.isComplete)
    todos.length = 0               // 清空原数组
    todos.push(...remain)          // 再把保留项推回去（保持同一个代理对象）
    localStorage.setItem("todos",JSON.stringify(todos))
}
function selectAll(){
    todos.forEach(item => {
        item.isComplete = true
    })
    localStorage.setItem("todos",JSON.stringify(todos))
}
function cancelSelectAll(){
    todos.forEach(item => {
        item.isComplete = false
    })
    localStorage.setItem("todos",JSON.stringify(todos))
}

// 设置过滤状态
function setFilter(filter) {
    currentFilter.value = filter
}
</script>

<style scoped>
.todo-main {
    background: white;
    border-radius: 10px;
    padding: 25px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    margin-bottom: 20px;
}

.input-section {
    margin-bottom: 25px;
    display: flex;
    flex-direction: column;
    gap: 15px;
}

.input-group {
    display: flex;
    gap: 10px;
    align-items: stretch;
}

.action-buttons {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 15px;
    flex-wrap: wrap;
}

.button-group {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
}

.todo-list-box {
    display: flex;
    gap: 8px;
    margin-bottom: 20px;
    justify-content: center;
    flex-wrap: wrap;
}

.todo-list-box-btn {
    padding: 8px 16px;
    border: 2px solid #e1e8ed;
    border-radius: 20px;
    background: white;
    color: #7f8c8d;
    font-size: 14px;
    cursor: pointer;
    transition: all 0.3s ease;
    position: relative;
}

.todo-list-box-btn:hover {
    border-color: #3498db;
    color: #3498db;
    transform: translateY(-1px);
}

.todo-list-box-btn.active {
    background: #3498db;
    border-color: #3498db;
    color: white;
    font-weight: 500;
}

.todo-list-box-btn span {
    background: rgba(255, 255, 255, 0.2);
    border-radius: 12px;
    padding: 2px 6px;
    margin-left: 6px;
    font-size: 12px;
    font-weight: bold;
}

.todo-list-box-btn.active span {
    background: rgba(255, 255, 255, 0.3);
}

.todo-input {
    flex: 1;
    padding: 12px 16px;
    border: 2px solid #e1e8ed;
    border-radius: 6px;
    font-size: 16px;
    outline: none;
    transition: border-color 0.3s ease;
    min-width: 200px;
}

.todo-input:focus {
    border-color: #3498db;
    box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.1);
}

.add-btn, .action-btn, .delete-btn, .edit-btn, .complete-btn {
    padding: 12px 20px;
    border: none;
    border-radius: 8px;
    font-size: 14px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.3s ease;
    min-width: 80px;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 6px;
    position: relative;
    overflow: hidden;
}

.add-btn.primary {
    background: linear-gradient(135deg, #3498db, #2980b9);
    color: white;
    box-shadow: 0 2px 8px rgba(52, 152, 219, 0.3);
}

.add-btn.primary:hover {
    background: linear-gradient(135deg, #2980b9, #21618c);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(52, 152, 219, 0.4);
}

.action-btn.secondary {
    background: linear-gradient(135deg, #95a5a6, #7f8c8d);
    color: white;
    box-shadow: 0 2px 6px rgba(149, 165, 166, 0.3);
}

.action-btn.secondary:hover {
    background: linear-gradient(135deg, #7f8c8d, #5d6d7e);
    transform: translateY(-1px);
    box-shadow: 0 3px 8px rgba(149, 165, 166, 0.4);
}

.action-btn.danger {
    background: linear-gradient(135deg, #e74c3c, #c0392b);
    color: white;
    box-shadow: 0 2px 8px rgba(231, 76, 60, 0.3);
}

.action-btn.danger:hover {
    background: linear-gradient(135deg, #c0392b, #922b21);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(231, 76, 60, 0.4);
}

.btn-icon {
    font-size: 16px;
    font-weight: bold;
}

.todo-list {
    list-style: none;
    padding: 0;
    margin: 0;
}

.todo-item {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 15px;
    margin-bottom: 10px;
    background: #f8f9fa;
    border-radius: 8px;
    border-left: 4px solid #3498db;
    transition: all 0.3s ease;
}

.todo-item:hover {
    background: #e9ecef;
    transform: translateX(5px);
}

.todo-item.completed {
    border-left-color: #27ae60;
    background: #d5edda;
}

.todo-content {
    display: flex;
    align-items: center;
    flex: 1;
    gap: 12px;
}

.todo-checkbox {
    width: 18px;
    height: 18px;
    cursor: pointer;
    accent-color: #27ae60;
}

.todo-text {
    font-size: 16px;
    color: #2c3e50;
    flex: 1;
}

.completed-text {
    text-decoration: line-through;
    color: #7f8c8d;
}

.edit-input {
    flex: 1;
    padding: 8px 12px;
    border: 2px solid #3498db;
    border-radius: 4px;
    font-size: 16px;
    outline: none;
}

.todo-actions {
    display: flex;
    gap: 8px;
}

.delete-btn {
    background-color: #e74c3c;
    color: white;
    padding: 8px 12px;
    font-size: 12px;
}

.delete-btn:hover {
    background-color: #c0392b;
}

.edit-btn {
    background-color: #f39c12;
    color: white;
    padding: 8px 12px;
    font-size: 12px;
}

.edit-btn:hover {
    background-color: #e67e22;
}

.complete-btn {
    background-color: #27ae60;
    color: white;
    padding: 8px 12px;
    font-size: 12px;
}

.complete-btn:hover {
    background-color: #229954;
}

.empty-state {
    text-align: center;
    color: #7f8c8d;
    font-style: italic;
    padding: 40px 20px;
    font-size: 16px;
}

@media (max-width: 600px) {
    .input-section {
        gap: 12px;
    }
    
    .input-group {
        flex-direction: column;
    }
    
    .todo-input {
        min-width: 100%;
        margin-bottom: 0;
    }
    
    .action-buttons {
        flex-direction: column;
        align-items: stretch;
        gap: 10px;
    }
    
    .button-group {
        justify-content: center;
    }
    
    .add-btn, .action-btn {
        padding: 10px 16px;
        font-size: 13px;
        min-width: 70px;
    }
    
    .btn-icon {
        font-size: 14px;
    }
    
    .todo-list-box {
        gap: 6px;
    }
    
    .todo-list-box-btn {
        padding: 6px 12px;
        font-size: 12px;
    }
    
    .todo-list-box-btn span {
        padding: 1px 4px;
        margin-left: 4px;
        font-size: 10px;
    }
    
    .todo-item {
        flex-direction: column;
        align-items: stretch;
        gap: 10px;
    }
    
    .todo-content {
        margin-bottom: 10px;
    }
    
    .todo-actions {
        justify-content: center;
    }
}

@media (max-width: 480px) {
    .action-buttons {
        gap: 8px;
    }
    
    .button-group {
        gap: 6px;
        flex-direction: column;
    }
    
    .action-btn.danger {
        margin-top: 5px;
    }
}
</style>