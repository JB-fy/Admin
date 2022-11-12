<script setup lang="ts">
const { t } = useI18n()

const emits = defineEmits(['query'])

const handleQueryEvent = () => { emits('query', queryForm.data) }

const queryForm = reactive({
    data: {
        sceneName: '',
        sceneCode: '',
        isStop: ''
    },
    submit: () => {
        handleQueryEvent()
    }
})
</script>

<template>
    <ElForm class="query-form" :model="queryForm.data" :inline="true">
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
            <ElButton type="primary" @click="handleQueryEvent">
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