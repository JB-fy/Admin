<script setup lang="ts">
const { t } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }
const queryForm = reactive({
    ref: null as any,
    loading: false,
    submit: () => {
        queryForm.loading = true
        listCommon.ref.getList(true).finally(() => {
            queryForm.loading = false
        })
    },
    reset: () => {
        queryForm.ref.resetFields()
        //queryForm.submit()
    }
})
</script>

<template>
    <ElForm class="query-form" :ref="(el: any) => { queryForm.ref = el }" :model="queryCommon.data" :inline="true"
        @keyup.enter="queryForm.submit">
        <ElFormItem prop="sceneName">
            <ElInput v-model="queryCommon.data.sceneName" :placeholder="t('view.auth.scene.sceneName')"
                :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="sceneCode">
            <ElInput v-model="queryCommon.data.sceneCode" :placeholder="t('view.auth.scene.sceneCode')"
                :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="isStop" style="width: 100px;">
            <ElSelect v-model="queryCommon.data.isStop" :placeholder="t('common.name.isStop')" :clearable="true">
                <ElOption :label="t('common.no')" value="0" />
                <ElOption :label="t('common.yes')" value="1" />
            </ElSelect>
        </ElFormItem>
        <ElFormItem>
            <ElButton type="primary" @click="queryForm.submit" :loading="queryForm.loading">
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