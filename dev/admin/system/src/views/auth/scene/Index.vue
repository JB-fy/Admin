<script setup lang="ts">
const { t } = useI18n()

const saveDrawer = reactive({
    ref: null as any,
    visible: false,
    handleClose: (done: Function) => {
        ElMessageBox.confirm('是否退出当前操作？').then(() => {
            done()
        })
    }
})
const saveForm = reactive({
    ref: null as any,
    data: {
        sceneName: '',
        sceneCode: ''
    },
    rules: {
        sceneName: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ],
        sceneCode: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ]
    },
    loading: false,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            await request('index.index')
            saveForm.loading = false
        })
    }
})
</script>

<template>
    <div>
        <div @click="saveDrawer.visible = true">新增</div>
        <div style="text-align: center; font-size: 300px; color: #409EFF;">场景列表</div>
        <ElDrawer :ref="(el: any) => { saveDrawer.ref = el }" v-model="saveDrawer.visible" title="新增" size="50%"
            :show-close="true" :before-close="saveDrawer.handleClose" custom-class="saveDrawer">
            <ElScrollbar>
                <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules"
                    label-width="auto" :status-icon="true">
                    <ElFormItem label="名称" prop="sceneName">
                        <ElInput v-model="saveForm.data.sceneName" placeholder="名称" />
                    </ElFormItem>
                    <ElFormItem label="场景标识" prop="sceneCode">
                        <ElInput v-model="saveForm.data.sceneCode" placeholder="标识" />
                    </ElFormItem>
                </ElForm>
                <div style="text-align: center; font-size: 300px; color: #409EFF;">场景列表</div>
            </ElScrollbar>
            <template #footer>
                <ElButton @click="saveDrawer.ref.handleClose()">{{ t('common.cancel') }}</ElButton>
                <ElButton :loading="saveForm.loading" type="primary" @click="saveForm.submit">
                    {{ t('common.save') }}
                </ElButton>
            </template>
        </ElDrawer>
    </div>
</template>

<style scoped>
:deep(.saveDrawer .el-drawer__header) {
    box-shadow: var(--el-box-shadow-lighter);
    padding: 10px;
    margin-bottom: 0px;
}

:deep(.saveDrawer .el-drawer__body) {
    padding: 0;
}

:deep(.saveDrawer .el-form) {
    margin: 20px;
}

:deep(.saveDrawer .el-drawer__footer) {
    box-shadow: var(--el-box-shadow-lighter);
    padding: 10px;
}
</style>