<script setup lang="ts">
import List from './List.vue'
import Query from './Query.vue'
import Save from './Save.vue'

//搜索
const queryCommon = reactive({
    data: {}
})
provide('queryCommon', queryCommon)

//列表
const listCommon = reactive({
    ref: null as any,
})
provide('listCommon', listCommon)

//保存
const saveCommon = reactive({
    visible: false,
    title: '',  //新增|编辑|复制
    data: {}
})
provide('saveCommon', saveCommon)
</script>

<template>
    <ElContainer class="app-container">
        <ElHeader>
            <Query />
        </ElHeader>

        <List :ref="(el: any) => { listCommon.ref = el }" />

        <Save v-if="saveCommon.visible" />
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