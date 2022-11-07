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
        account: '',
        password: ''
    },
    rules: {
        account: [
            { type: 'string', required: true, min: 4, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 4, max: 30 }) }
        ],
        password: [
            { type: 'string', required: true, min: 6, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 6, max: 30 }) }
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
                <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules">
                    <ElFormItem prop="account">
                        <ElInput v-model="saveForm.data.account" :placeholder="t('common.account')">
                            <template #prefix>
                                <AutoiconEpUser />
                            </template>
                        </ElInput>
                    </ElFormItem>
                    <ElFormItem prop="password">
                        <ElInput v-model="saveForm.data.password" type="password" :placeholder="t('common.password')"
                            :show-password="true">
                            <template #prefix>
                                <AutoiconEpLock />
                            </template>
                        </ElInput>
                    </ElFormItem>
                </ElForm>
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
:deep(.saveDrawer .el-drawer__body) {
    padding: 0;
}
</style>