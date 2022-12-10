<script setup lang="ts">
onMounted(() => {
    request('index/index')
})
onActivated(() => {
    request('index/index')
})
let input = ref('')
</script>

<template>
    <div>
        <el-input v-model="input" placeholder="请输入" />
        <div style="text-align: center; font-size: 300px; color: #409EFF;">个人中心</div>
    </div>
</template>