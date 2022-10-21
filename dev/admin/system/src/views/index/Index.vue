<script setup lang="ts">
onMounted(() => {
    request('index.index', undefined, false)
        .then((data) => {
            console.log(data)
        })
        .catch((error) => {
            errorHandle(<Error>error)
        })
})
let input = ref('')
</script>

<template>
    <div>
        <el-input v-model="input" placeholder="请输入" />
        <div style="text-align: center; font-size: 300px; color: #409EFF;">欢迎</div>
    </div>
</template>