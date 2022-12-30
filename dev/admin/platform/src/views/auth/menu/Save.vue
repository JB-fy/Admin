<script setup lang="ts">
const { t } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        pid: 0,
        sort: 50,
        ...saveCommon.data
    } as { [propName: string]: any },
    rules: {
        menuName: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        sceneId: [
            { type: 'integer', required: true, min: 1, trigger: 'change', message: t('validation.select') }
        ],
        pid: [
            { type: 'integer', min: 0, trigger: 'change', message: t('validation.select') }
        ],
        extraData: [
            {
                type: 'object',
                fields: {
                    title: {
                        type: 'object',
                        fields: {
                            'zh-cn': { type: 'string', required: true, min: 1, message: 'title.zh-cn' + t('validation.min.string', { min: 1 }) },
                            en: { type: 'string', min: 1, message: 'title.en' + t('validation.min.string', { min: 1 }) }
                        },
                        message: 'title' + t('validation.regex')
                    },
                    icon: { type: 'string', min: 1, message: 'icon' + t('validation.min.string', { min: 1 }) },
                    url: { type: 'string', min: 1, message: 'url' + t('validation.min.string', { min: 1 }) }
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
        ],
        sort: [
            { type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) }
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
                ...removeEmptyOfObj(saveForm.data, false)
            }
            try {
                await request('auth/menu/save', param, true)
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
                <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules"
                    label-width="auto" :status-icon="true" :scroll-to-error="true">
                    <ElFormItem :label="t('common.name.auth.menu.menuName')" prop="menuName">
                        <ElInput v-model="saveForm.data.menuName" :placeholder="t('common.name.auth.menu.menuName')"
                            minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.rel.sceneId')" prop="sceneId">
                        <MySelect v-model="saveForm.data.sceneId"
                            :api="{ code: 'auth/scene/list', param: { field: ['id', 'sceneName'] } }"
                            @change="() => { saveForm.data.pid = 0 }" />
                    </ElFormItem>
                    <ElFormItem v-if="saveForm.data.sceneId" :label="t('common.name.rel.pid')" prop="pid">
                        <MyCascader v-model="saveForm.data.pid"
                            :api="{ code: 'auth/menu/tree', param: { field: ['id', 'menuName'], where: { sceneId: saveForm.data.sceneId, excId: saveForm.data.id } } }"
                            :defaultOptions="[{ id: 0, menuName: t('common.name.without') }]" :clearable="false" />
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.extraData')" prop="extraData">
                        <ElAlert :title="t('view.auth.menu.tip.extraData')" type="info" :show-icon="true"
                            :closable="false" />
                        <ElInput v-model="saveForm.data.extraData" type="textarea" :autosize="{ minRows: 3 }" />
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.sort')" prop="sort">
                        <ElInputNumber v-model="saveForm.data.sort" :precision="0" :min="0" :max="100" :step="1"
                            :step-strictly="true" controls-position="right" :value-on-clear="50" />
                        <label>
                            <ElAlert :title="t('common.tip.sort')" type="info" :show-icon="true" :closable="false" />
                        </label>
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.isStop')" prop="isStop">
                        <ElSwitch v-model="saveForm.data.isStop" :active-value="1" :inactive-value="0"
                            :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')"
                            style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
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