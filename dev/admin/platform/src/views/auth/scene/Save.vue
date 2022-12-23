<script setup lang="ts">
const { t } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    rules: {
        sceneName: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ],
        sceneCode: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ],
        sceneConfig: [
            {
                validator: (rule: any, value: any, callback: any) => {
                    try {
                        if (value === '' || value === null || value === undefined) {
                            callback()
                        }
                        JSON.parse(value)
                        callback()
                    } catch (e) {
                        callback(new Error())
                    }
                },
                trigger: 'blur',
                message: t('validation.json')
            },
        ],
        isStop: [
            { type: 'enum', enum: [0, 1]/* Object.keys(customOption.yesOrNo).map(Number) */, trigger: 'change', message: t('validation.select') }
        ]
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = {
                ...removeEmptyOfObj(saveCommon.data, false)
            }
            try {
                await request('auth/scene/save', param, true)
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } catch (error) { }
            saveForm.loading = false
        })
    }
})

const saveDrawer = reactive({
    ref: null as any,
    beforeClose: (done: Function) => {
        //确定退出当前操作？
        ElMessageBox.confirm('', {
            type: 'info',
            title: t('common.tip.configExit'),
            center: true,
            showClose: false,
        }).then(() => {
            done()
        }).catch(() => { })
    },
    closed: () => {
        saveForm.ref.clearValidate()    //清理表单验证错误提示
    },
    buttonClose: () => {
        //saveCommon.visible = false
        saveDrawer.ref.handleClose()    //会触发beforeClose
    }
})
</script>

<template>
    <div class="save-drawer">
        <ElDrawer :ref="(el: any) => { saveDrawer.ref = el }" v-model="saveCommon.visible" :title="saveCommon.title"
            size="50%" :before-close="saveDrawer.beforeClose" @closed="saveDrawer.closed">
            <ElScrollbar>
                <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveCommon.data" :rules="saveForm.rules"
                    label-width="auto" :status-icon="true" :scroll-to-error="true">
                    <ElFormItem :label="t('view.auth.scene.sceneName')" prop="sceneName">
                        <ElInput v-model="saveCommon.data.sceneName" :placeholder="t('view.auth.scene.sceneName')"
                            minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
                    </ElFormItem>
                    <ElFormItem :label="t('view.auth.scene.sceneCode')" prop="sceneCode">
                        <ElAlert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true"
                            :closable="false" />
                        <ElInput v-model="saveCommon.data.sceneCode" :placeholder="t('view.auth.scene.sceneCode')"
                            minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
                    </ElFormItem>
                    <ElFormItem :label="t('view.auth.scene.sceneConfig')" prop="sceneConfig">
                        <ElInput v-model="saveCommon.data.sceneConfig" type="textarea" :autosize="{ minRows: 3 }" />
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.isStop')" prop="isStop">
                        <ElSwitch v-model="saveCommon.data.isStop" :active-value="1" :inactive-value="0"
                            :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')"
                            style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)" />
                    </ElFormItem>
                </ElForm>
            </ElScrollbar>
            <template #footer>
                <ElButton @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</ElButton>
                <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                    {{ t('common.save') }}
                </ElButton>
            </template>
        </ElDrawer>
    </div>
</template>

<style scoped>
.save-drawer :deep(.el-drawer .el-drawer__header) {
    box-shadow: var(--el-box-shadow-lighter);
    padding: 10px;
    margin-bottom: 0px;
}

.save-drawer :deep(.el-drawer .el-drawer__body) {
    padding: 0;
}

.save-drawer :deep(.el-drawer .el-form) {
    margin: 20px;
}

.save-drawer :deep(.el-drawer .el-drawer__footer) {
    box-shadow: var(--el-box-shadow-lighter);
    padding: 10px;
}

.save-drawer :deep(.el-alert) {
    padding: 0 0.5rem;
}
</style>