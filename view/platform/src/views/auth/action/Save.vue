<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        ...saveCommon.data,
    } as { [propName: string]: any },
    rules: {
        actionName: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
        ],
        actionCode: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
            { trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') },
        ],
        remark: [{ type: 'string', trigger: 'blur', max: 120, message: t('validation.max.string', { max: 120 }) }],
        isStop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
        sceneIdArr: [
            { required: true, message: t('validation.required') },
            { type: 'array', trigger: 'change', min: 1, message: t('validation.select'), defaultField: { type: 'integer', min: 1, message: t('validation.min.number', { min: 1 }) } },
        ],
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            try {
                if (param?.idArr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/auth/action/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/auth/action/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } catch (error) {}
            saveForm.loading = false
        })
    },
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
            })
                .then(() => {
                    done()
                })
                .catch(() => {})
        } else {
            done()
        }
    },
    buttonClose: () => {
        //saveCommon.visible = false
        saveDrawer.ref.handleClose() //会触发beforeClose
    },
})
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
                <el-form-item :label="t('auth.action.name.actionName')" prop="actionName">
                    <el-input v-model="saveForm.data.actionName" :placeholder="t('auth.action.name.actionName')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('auth.action.name.actionCode')" prop="actionCode">
                    <el-input v-model="saveForm.data.actionCode" :placeholder="t('auth.action.name.actionCode')" maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('auth.action.name.remark')" prop="remark">
                    <el-input v-model="saveForm.data.remark" type="textarea" :autosize="{ minRows: 3 }" maxlength="120" :show-word-limit="true" />
                </el-form-item>
                <el-form-item :label="t('auth.action.name.isStop')" prop="isStop">
                    <el-switch
                        v-model="saveForm.data.isStop"
                        :active-value="1"
                        :inactive-value="0"
                        :inline-prompt="true"
                        :active-text="t('common.yes')"
                        :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </el-form-item>
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                {{ t('common.save') }}
            </el-button>
        </template>
    </el-drawer>
</template>
