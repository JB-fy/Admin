<script setup lang="ts">
let input = ref('')
</script>

<template>
    <div>
        <el-input v-model="input" placeholder="请输入" />
        <div style="text-align: center; font-size: 300px; color: #409EFF;">管理员</div>
    </div>
</template>