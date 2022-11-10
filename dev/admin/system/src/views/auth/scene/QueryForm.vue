<script setup lang="ts">
const { t } = useI18n()

const emits = defineEmits(['search'])

const handleSearchEvent = () => emits('search', queryForm.data)

const queryForm = reactive({
    ref: null as any,
    data: {
        sceneName: '',
        sceneCode: '',
        isStop: ''
    },
    loading: false,
    submit: () => {
        queryForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            queryForm.loading = true
            await request('index.index', queryForm.data)
            handleSearchEvent()
            queryForm.loading = false
        })
    }
})
</script>

<template>
    <ElForm class="query-form" :ref="(el: any) => { queryForm.ref = el }" :model="queryForm.data" :inline="true">
        <ElFormItem prop="sceneName">
            <ElInput v-model="queryForm.data.sceneName" placeholder="名称" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="sceneCode">
            <ElInput v-model="queryForm.data.sceneCode" placeholder="标识" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="isStop" style="width: 100px;">
            <ElSelect v-model="queryForm.data.isStop" placeholder="停用" :clearable="true">
                <ElOption :label="t('common.no')" value="0" />
                <ElOption :label="t('common.yes')" value="1" />
            </ElSelect>
        </ElFormItem>
        <ElFormItem>
            <ElButton :loading="queryForm.loading" type="primary" @click="queryForm.submit">
                <AutoiconEpSearch />{{ t('common.query') }}
            </ElButton>
        </ElFormItem>
    </ElForm>
</template>

<style scoped>
.query-form :deep(.el-form-item) {
    margin: 0 10px 10px 0;
}
</style>