<script setup lang="ts">
const { t } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }
//可不做。主要作用：新增时设置默认值；知道有哪些字段
saveCommon.data = {
    sort: 50,
    ...saveCommon.data
}

const saveForm = reactive({
    ref: null as any,
    loading: false,
    rules: {
        menuName: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ],
        sceneId: [
            { type: 'integer', required: true, min: 1, trigger: 'change', message: t('validation.select') }
        ],
        pid: [
            { type: 'integer', min: 0, trigger: 'change', message: t('validation.select') }
        ],
        extraData: [
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
                ...removeEmptyOfObj(saveCommon.data, false)
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
                    <ElFormItem :label="t('common.name.auth.menu.menuName')" prop="menuName">
                        <ElInput v-model="saveCommon.data.menuName" :placeholder="t('common.name.auth.menu.menuName')"
                            minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.rel.sceneId')" prop="sceneId">
                        <MySelectScroll v-model="saveCommon.data.sceneId"
                            :api="{ code: 'auth/scene/list', param: { field: ['id', 'sceneName'] } }" />
                    </ElFormItem>
                    <ElFormItem v-if="saveCommon.data.sceneId" :label="t('common.name.rel.pid')" prop="pid">
                        <MyCascader v-model="saveCommon.data.pid"
                            :api="{ code: 'auth/menu/tree', param: { field: ['id', 'menuName'], where: { sceneId: saveCommon.data.sceneId } } }" />
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.extraData')" prop="extraData">
                        <ElAlert :title="t('view.auth.menu.tip.extraData')" type="info" :show-icon="true"
                            :closable="false" />
                        <ElInput v-model="saveCommon.data.extraData" type="textarea" :autosize="{ minRows: 3 }" />
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.sort')" prop="sort">
                        <ElInputNumber v-model="saveCommon.data.sort" :precision="0" :min="0" :max="100" :step="1"
                            :step-strictly="true" controls-position="right" :value-on-clear="50" />
                        <label>
                            <ElAlert :title="t('common.tip.sort')" type="info" :show-icon="true" :closable="false" />
                        </label>
                    </ElFormItem>
                    <ElFormItem :label="t('common.name.isStop')" prop="isStop">
                        <ElSwitch v-model="saveCommon.data.isStop" :active-value="1" :inactive-value="0"
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