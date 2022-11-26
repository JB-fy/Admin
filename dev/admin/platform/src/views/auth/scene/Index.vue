<script setup lang="ts">
import List from './List.vue'
import Query from './Query.vue'
import Save from './Save.vue'

const { t } = useI18n()
const list = reactive({
    ref: null as any,
})

const queryData = reactive({})
provide('queryData', queryData)
// 搜索
const handleQuery = () => {
    list.ref.getList(true)
}

const saveVisible = ref(false)
provide('saveVisible', saveVisible)
</script>

<template>
    <ElContainer class="app-container">
        <ElHeader>
            <Query @query="handleQuery" />
        </ElHeader>

        <List :ref="(el: any) => { list.ref = el }" />

        <Save />
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