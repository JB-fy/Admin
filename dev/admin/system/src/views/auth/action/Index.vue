<script setup lang="ts">
import { index } from '@/api'

onMounted(async () => {
    await index()
})
let input = ref('')
</script>

<template>
    <div>
        <el-input v-model="input" placeholder="请输入" />
        <div style="text-align: center; font-size: 300px; color: #409EFF;">操作列表</div>
    </div>
</template>