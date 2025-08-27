<template>
    <div class="todo-footer">
        <div class="footer-content">
            <p class="copyright">© 2025 待办事项管理系统</p>
            <p class="tips">提示：双击任务可快速编辑，使用 Enter 键快速添加任务</p>
        </div>
    </div>
</template>

<script setup>

</script>

<style scoped>
.todo-footer {
    background: white;
    border-radius: 10px;
    padding: 20px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    margin-top: 20px;
}

.footer-content {
    text-align: center;
}

.copyright {
    color: #2c3e50;
    font-weight: 500;
    margin: 0 0 10px 0;
    font-size: 14px;
}

.tips {
    color: #7f8c8d;
    font-size: 12px;
    margin: 0;
    font-style: italic;
}

@media (max-width: 480px) {
    .todo-footer {
        padding: 15px;
    }
    
    .tips {
        font-size: 11px;
    }
}
</style>