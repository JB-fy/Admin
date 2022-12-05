<script setup lang="ts">
import List from './List.vue'
import Query from './Query.vue'
import Save from './Save.vue'


/*--------列表 开始--------*/
const list = reactive({
    ref: null as any,
})
/*--------列表 开始--------*/

/*--------搜索 开始--------*/
const query = reactive({
    data: {},
    handleQuery: () => {
        list.ref.getList(true)
    }
})
provide('query', query)

const queryData = ref({})
provide('queryData', queryData)
//搜索
const handleQuery = () => {
    list.ref.getList(true)
}
/*--------搜索 结束--------*/

/*--------保存（新增|编辑） 开始--------*/
const save = reactive({
    visible: false,
    title: '新增',
    data: {},
    handleSave: () => {
        list.ref.getList(true)
        save.visible = false
    }
})
provide('save', save)

const saveData = ref({})
provide('saveData', saveData)
const saveVisible = ref(false)
provide('saveVisible', saveVisible)
//保存成功处理
const handleSave = () => {
    list.ref.getList(true)
    saveVisible.value = false
}
/*--------保存（新增|编辑） 结束--------*/
</script>

<template>
    <ElContainer class="app-container">
        <ElHeader>
            <Query @query="handleQuery" />
        </ElHeader>

        <List :ref="(el: any) => { list.ref = el }" />

        <Save @save="handleSave" />
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