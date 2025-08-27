<template>
    <div class="app">
        <TodoNav />
        <TodoMain />
        <TodoFooter />
    </div>
</template>

<script setup>
import TodoNav from './components/TodoNav.vue'
import TodoMain from './components/TodoMain.vue'
import TodoFooter from './components/TodoFooter.vue'
</script>

<style scoped>
.app {
    max-width: 600px;
    margin: 0 auto;
    padding: 20px;
    font-family: 'Arial', sans-serif;
    background-color: #f5f5f5;
    min-height: 100vh;
    box-sizing: border-box;
}

/* 全局重置样式 */
:global(*) {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

:global(body) {
    background-color: #e8f4f8;
    font-family: 'Arial', sans-serif;
    line-height: 1.6;
}
</style>