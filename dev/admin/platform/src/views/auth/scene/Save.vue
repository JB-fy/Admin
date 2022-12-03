<script setup lang="ts">
const { t } = useI18n()

const saveVisible = inject('saveVisible')

const saveDrawer = reactive({
    ref: null as any,
    handleClose: (done: Function) => {
        //done()
        ElMessageBox.confirm('是否退出当前操作？').then(() => {
            done()
        })
    }
})

const saveForm = reactive({
    ref: null as any,
    data: {
        sceneName: '',
        sceneCode: '',
        sceneConfig: '',
        isStop: 0
    },
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
    loading: false,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            try {
                await request('auth.scene.save', param, true)
            } catch (error) {
            }
            saveForm.loading = false
        })
    }
})
</script>

<template>
    <div class="save-drawer">
        <ElDrawer :ref="(el: any) => { saveDrawer.ref = el }" v-model="saveVisible" title="新增" size="50%"
            :before-close="saveDrawer.handleClose">
            <ElScrollbar>
                <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules"
                    label-width="auto" :status-icon="true">
                    <ElFormItem label="名称" prop="sceneName">
                        <ElInput v-model="saveForm.data.sceneName" placeholder="名称" minlength="1" maxlength="30"
                            :show-word-limit="true" :clearable="true" />
                        <!-- <ElInput style="max-width: 300px;" v-model="saveForm.data.sceneName" placeholder="名称"
                            minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
                        <label>
                            <ElAlert title="标题" type="info" :show-icon="true" :closable="false" />
                        </label> -->
                    </ElFormItem>
                    <ElFormItem label="场景标识" prop="sceneCode">
                        <ElInput v-model="saveForm.data.sceneCode" placeholder="场景标识" minlength="1" maxlength="30" />
                    </ElFormItem>
                    <ElFormItem label="场景配置" prop="sceneConfig">
                        <!-- <ElAlert title="标题" type="info" :show-icon="true" :closable="false" /> -->
                        <ElInput v-model="saveForm.data.sceneConfig" type="textarea" :autosize="{ minRows: 3 }" />
                    </ElFormItem>
                    <ElFormItem label="停用" prop="isStop">
                        <ElSwitch v-model="saveForm.data.isStop" :active-value="1" :inactive-value="0"
                            :inline-prompt="true" active-text="是" inactive-text="否"
                            style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)" />
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