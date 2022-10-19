<script setup lang="ts">
import { index } from '@/api'

onMounted(async () => {
    try {
        await index()
    } catch (e) {
        await errorHandle(e)
    }
})
let input = ref('')
</script>

<template>
    <div>
        <el-input v-model="input" placeholder="请输入" />
        <div style="text-align: center; font-size: 300px; color: #409EFF;">权限菜单</div>
    </div>
</template>