<script setup lang="ts">
const { t } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }
//可不做。主要作用：新增时设置默认值；知道有哪些字段
/* saveCommon.data = {
    sceneName: '',
    sceneCode: '',
    sceneConfig: '',
    isStop: 0,
    ...saveCommon.data
} */

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
                        const valueTmp = JSON.parse(value)
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
    },
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
                await request('auth.scene.save', param, true)
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } catch (error) {
            }
            saveForm.loading = false
        })
    }
})

const saveDrawer = reactive({
    ref: null as any,
    beforeClose: (done: Function) => {
        ElMessageBox.confirm('确定放弃当前操作？', { center: true }).then(() => {
            done()
        }).catch((error) => {
        })
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
                    <ElFormItem label="名称" prop="sceneName">
                        <ElInput v-model="saveCommon.data.sceneName" placeholder="名称" minlength="1" maxlength="30"
                            :show-word-limit="true" :clearable="true" />
                    </ElFormItem>
                    <ElFormItem label="场景标识" prop="sceneCode">
                        <ElAlert title="值不能与现有记录重复" type="info" :show-icon="true" :closable="false" />
                        <ElInput v-model="saveCommon.data.sceneCode" placeholder="场景标识" minlength="1" maxlength="30"
                            :show-word-limit="true" :clearable="true" />
                        <!-- <ElInput v-model="saveCommon.data.sceneCode" placeholder="场景标识" minlength="1" maxlength="30"
                            :show-word-limit="true" :clearable="true" style="max-width: 300px;" />
                        <label>
                            <ElAlert title="值不能与现有记录重复" type="info" :show-icon="true" :closable="false" />
                        </label> -->
                    </ElFormItem>
                    <ElFormItem label="场景配置" prop="sceneConfig">
                        <ElInput v-model="saveCommon.data.sceneConfig" type="textarea" :autosize="{ minRows: 3 }" />
                    </ElFormItem>
                    <ElFormItem label="停用" prop="isStop">
                        <ElSwitch v-model="saveCommon.data.isStop" :active-value="1" :inactive-value="0"
                            :inline-prompt="true" active-text="是" inactive-text="否"
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