<script setup lang="ts">
const { t } = useI18n()

const queryData = inject('queryData') as { [propName: string]: any }

const emits = defineEmits(['query'])
const queryForm = reactive({
    ref: null as any,
    submit: () => {
        emits('query')
    },
    reset: () => {
        queryForm.ref.resetFields()
        //emits('query')
    }
})
</script>

<template>
    <ElForm class="query-form" :ref="(el: any) => { queryForm.ref = el }" :model="queryData" :inline="true"
        @keyup.enter="queryForm.submit">
        <ElFormItem prop="sceneName">
            <ElInput v-model="queryData.sceneName" placeholder="名称" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="sceneCode">
            <ElInput v-model="queryData.sceneCode" placeholder="标识" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="isStop" style="width: 100px;">
            <ElSelect v-model="queryData.isStop" placeholder="停用" :clearable="true">
                <ElOption :label="t('common.no')" value="0" />
                <ElOption :label="t('common.yes')" value="1" />
            </ElSelect>
        </ElFormItem>
        <ElFormItem>
            <ElButton type="primary" @click="queryForm.submit">
                <AutoiconEpSearch />{{ t('common.query') }}
            </ElButton>
            <ElButton type="info" @click="queryForm.reset">
                <AutoiconEpCircleClose />{{ t('common.reset') }}
            </ElButton>
        </ElFormItem>
    </ElForm>
</template>

<style scoped>
.query-form :deep(.el-form-item) {
    margin: 0 10px 10px 0;
}
</style>