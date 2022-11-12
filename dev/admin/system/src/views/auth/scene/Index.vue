<script setup lang="ts">
import QueryForm from './QueryForm.vue'
import SaveForm from './SaveForm.vue'
import List from './List.vue'

const { t } = useI18n()

const queryFormData = ref({})
provide('queryFormData', queryFormData)
// 搜索
const queryFormEmit = (data) => {
    queryFormData.value = data
    console.log(data)
}

const saveDrawerVisible = ref(false)
provide('saveDrawerVisible', saveDrawerVisible)
</script>

<template>
    <ElContainer class="app-container">
        <ElHeader>
            <QueryForm @query="queryFormEmit" />
        </ElHeader>

        <ElMain style="padding: 0;">
            <List />
        </ElMain>

        <SaveForm />
    </ElContainer>
</template>

<style scoped>
.app-container {
    height: 100%;
}

.app-container :deep(.el-header) {
    height: auto;
    padding: 0;
}
</style>