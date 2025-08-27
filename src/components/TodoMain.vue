<template>
    <div class="todo-main">
        <div class="input-section">
            <input 
                type="text" 
                v-model="todo" 
                @keyup.enter="addTodo" 
                placeholder="请输入待办事项"
                class="todo-input"
            >
            <button @click="addTodo" class="add-btn">添加</button>
            <button @click="selectAll" class="add-btn">全选</button>
            <button @click="deleteComplete" class="clear-btn">删除已完成</button>
        </div>
        <ul class="todo-list">
            <li v-for="item in todos" :key="item.id" class="todo-item" :class="{ completed: item.isComplete }">
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
        <div v-if="todos.length === 0" class="empty-state">
            暂无待办事项，添加一个开始吧！
        </div>
    </div>
</template>

<script setup>
import { ref,reactive } from 'vue'
import { nanoid } from 'nanoid'
let todo=ref("")
// let todos=reactive([])
let todos=reactive(JSON.parse(localStorage.getItem("todos"))||[])
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
}
function todoEdit(id){
    todos.find(item=>item.id===id).edit=true
}
function todoComplete(id){
    todos.find(item=>item.id===id).edit=false
}
function deleteComplete(){
    const remain = todos.filter(item => !item.isComplete)
    todos.length = 0               // 清空原数组
    todos.push(...remain)          // 再把保留项推回去（保持同一个代理对象）
}
function selectAll(){
    todos.forEach(item => {
        item.isComplete = true
    })
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
    display: flex;
    gap: 10px;
    margin-bottom: 25px;
    flex-wrap: wrap;
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

.add-btn, .clear-btn, .delete-btn, .edit-btn, .complete-btn {
    padding: 12px 20px;
    border: none;
    border-radius: 6px;
    font-size: 14px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.3s ease;
    min-width: 80px;
}

.add-btn {
    background-color: #3498db;
    color: white;
}

.add-btn:hover {
    background-color: #2980b9;
    transform: translateY(-1px);
}

.clear-btn {
    background-color: #e74c3c;
    color: white;
}

.clear-btn:hover {
    background-color: #c0392b;
    transform: translateY(-1px);
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
        flex-direction: column;
    }
    
    .todo-input {
        min-width: 100%;
        margin-bottom: 10px;
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
</style>