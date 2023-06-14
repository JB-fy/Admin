<script setup lang="ts">
const { t } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        ...saveCommon.data
    } as { [propName: string]: any },
    rules: {
        sceneName: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        sceneCode: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        sceneConfig: [
            {
                type: 'object',
                fields: {
                    signKey: { type: 'string', min: 1, message: 'signKey' + t('validation.min.string', { min: 1 }) },
                    signType: { type: 'string', min: 1, message: 'signType' + t('validation.min.string', { min: 1 }) },
                    expireTime: { type: 'integer', min: 1, message: 'expireTime' + t('validation.min.number', { min: 1 }) }
                },
                transform(value: any) {
                    if (value === '' || value === null || value === undefined) {
                        return undefined
                    }
                    try {
                        return JSON.parse(value)
                    } catch (e) {
                        return value
                    }
                },
                trigger: 'blur',
                message: t('validation.json')
            },
            /* {
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
            }, */
        ],
        isStop: [
            /* { type: 'enum', enum: tm('common.status.whether').map((item) => item.value), trigger: 'change', message: t('validation.select') } */
            { type: 'enum', enum: [0, 1], trigger: 'change', message: t('validation.select') }
        ]
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data, false)
            try {
                if (param?.idArr?.length > 0) {
                    await request('/auth/scene/update', param, true)
                } else {
                    await request('/auth/scene/create', param, true)
                }
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
    <ElDrawer class="save-drawer" :ref="(el: any) => { saveDrawer.ref = el }" v-model="saveCommon.visible"
        :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <ElScrollbar>
            <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules"
                label-width="auto" :status-icon="true" :scroll-to-error="true">
                <ElFormItem :label="t('auth.scene.name.sceneName')" prop="sceneName">
                    <ElInput v-model="saveForm.data.sceneName" :placeholder="t('auth.scene.name.sceneName')"
                        minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('auth.scene.name.sceneCode')" prop="sceneCode">
                    <ElInput v-model="saveForm.data.sceneCode" :placeholder="t('auth.scene.name.sceneCode')"
                        minlength="1" maxlength="30" :show-word-limit="true" :clearable="true"
                        style="max-width: 250px;" />
                    <label>
                        <ElAlert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true"
                            :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem :label="t('auth.scene.name.sceneConfig')" prop="sceneConfig">
                    <ElAlert :title="t('auth.scene.tip.sceneConfig')" type="info" :show-icon="true"
                        :closable="false" />
                    <ElInput v-model="saveForm.data.sceneConfig" type="textarea" :autosize="{ minRows: 3 }" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.isStop')" prop="isStop">
                    <ElSwitch v-model="saveForm.data.isStop" :active-value="1" :inactive-value="0" :inline-prompt="true"
                        :active-text="t('common.yes')" :inactive-text="t('common.no')"
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
</template>