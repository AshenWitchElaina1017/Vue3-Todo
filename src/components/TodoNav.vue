<template>
    <div class="todo-nav">
        <h1>待办事项管理</h1>
    </div>
</template>

<script setup>
import { ref } from 'vue'

const inputValue = ref('')

const addTodo = () => {
    console.log(inputValue.value)
}
</script>

<style scoped>
.todo-nav {
    text-align: center;
    margin-bottom: 30px;
    padding: 20px 0;
}

.todo-nav h1 {
    color: #2c3e50;
    font-size: 2.5rem;
    font-weight: 300;
    margin: 0;
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

@media (max-width: 480px) {
    .todo-nav h1 {
        font-size: 2rem;
    }
}
</style>