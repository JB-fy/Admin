<script setup lang="ts">
import QueryForm from './QueryForm.vue'
import SaveForm from './SaveForm.vue'
import List from './List.vue'

const { t } = useI18n()
const list = reactive({
    ref: null as any,
})

const queryFormData = reactive({})
provide('queryFormData', queryFormData)
// 搜索
const handleQuery = () => {
    list.ref.getList(true)
}

const saveFormVisible = ref(false)
provide('saveFormVisible', saveFormVisible)
</script>

<template>
    <ElContainer class="app-container">
        <ElHeader>
            <QueryForm @query="handleQuery" />
        </ElHeader>

        <ElMain style="padding: 0;">
            <List :ref="(el: any) => { list.ref = el }" />
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