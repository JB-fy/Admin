<script setup lang="ts">
const { t } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    rules: {
        actionName: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ],
        actionCode: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ],
        remark: [
            { type: 'string', min: 0, max: 120, trigger: 'blur', message: t('validation.max.string', { max: 120 }) }
        ],
        isStop: [
            { type: 'enum', enum: [0, 1], trigger: 'change', message: t('validation.select') }
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
                await request('auth/action/save', param, true)
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } catch (error) { }
            saveForm.loading = false
        })
    }
})

const saveDrawer = reactive({
    ref: null as any,
    size: useSettingStore().saveDrawer.size,
    beforeClose: (done: Function) => {
        if (useSettingStore().saveDrawer.isTipClose) {
            ElMessageBox.confirm('', {
                type: 'info',
                title: t('common.tip.configExit'),
                center: true,
                showClose: false,
            }).then(() => {
                done()
            }).catch(() => { })
        } else {
            done()
        }
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
            :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
            <ElScrollbar>
                <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveCommon.data" :rules="saveForm.rules"
                    label-width="auto" :status-icon="true" :scroll-to-error="true">
                    <ElFormItem :label="t('common.name.auth.action.actionName')" prop="actionName">
                        <ElInput v-model="saveCommon.data.actionName"
                            :placeholder="t('common.name.auth.action.actionName')" minlength="1" maxlength="30"
                            :show-word-limit="true" :clearable="true" />
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.auth.action.actionCode')" prop="actionCode">
                        <ElInput v-model="saveCommon.data.actionCode"
                            :placeholder="t('common.name.auth.action.actionCode')" minlength="1" maxlength="30"
                            :show-word-limit="true" :clearable="true" style="max-width: 250px;" />
                        <label>
                            <ElAlert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true"
                                :closable="false" />
                        </label>
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.rel.sceneIdArr')" prop="sceneIdArr">
                        <MySelectScroll v-model="saveCommon.data.sceneIdArr"
                            :api="{ code: 'auth/scene/list', param: { field: ['id', 'sceneName'] }, limit:5 }" :multiple="true" />

                        <MyTransfer v-model="saveCommon.data.sceneIdArr"
                            :api="{ code: 'auth/scene/list', param: { field: ['id', 'sceneName'] }, limit:5 }" :multiple="true" />
                        <ElTransfer v-model="saveCommon.data.sceneIdArr"
                            :filter-placeholder="t('common.name.rel.sceneIdArr')" :filterable="true" />
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.remark')" prop="remark">
                        <ElInput v-model="saveCommon.data.remark" type="textarea" :autosize="{ minRows: 3 }"
                            minlength="0" maxlength="120" :show-word-limit="true" />
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