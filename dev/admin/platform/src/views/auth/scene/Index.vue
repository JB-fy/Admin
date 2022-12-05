<script setup lang="ts">
import List from './List.vue'
import Query from './Query.vue'
import Save from './Save.vue'

/*--------搜索 开始--------*/
const queryData = ref({})
provide('queryData', queryData)
//搜索
const handleQuery = () => {
    list.ref.getList(true)
}
/*--------搜索 结束--------*/

/*--------列表 开始--------*/
const list = reactive({
    ref: null as any,
})
/*--------列表 开始--------*/

/*--------保存（新增|编辑） 开始--------*/
const save = reactive({
    visible: false,
    title: '',  //新增|编辑|复制
    data: {},
    handleSave: () => {
        //保存成功处理
        list.ref.getList(true)
        save.visible = false
    }
})
provide('save', save)
/*--------保存（新增|编辑） 结束--------*/
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