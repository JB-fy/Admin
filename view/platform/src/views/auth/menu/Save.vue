<script setup lang="ts">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        sort: 50,
        isStop: 0,
        ...saveCommon.data,
    } as { [propName: string]: any },
    rules: {
        menuName: [
            { type: 'string', required: true, max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
        ],
        sceneId: [{ type: 'integer', required: true, min: 1, trigger: 'change', message: t('validation.select') }],
        pid: [{ type: 'integer', min: 0, trigger: 'change', message: t('validation.select') }],
        menuIcon: [
            { type: 'string', max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
        ],
        menuUrl: [{ type: 'string', max: 120, trigger: 'blur', message: t('validation.max.string', { max: 120 }) }],
        extraData: [
            {
                type: 'object',
                /* fields: {
					xxxx: { type: 'string', required: true, message: 'xxxx' + t('validation.required') },
					xxxx: { type: 'integer', required: true, min: 1, message: 'xxxx' + t('validation.min.number', { min: 1 }) },
				}, */
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
                message: t('validation.json'),
            },
        ],
        sort: [{ type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) }],
        isStop: [{ type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') }],
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data, false)
            if (param.sceneId === undefined) {
                param.sceneId = 0
            }
            if (param.pid === undefined) {
                param.pid = 0
            }
            try {
                if (param?.idArr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/create', param, true)
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
    <ElDrawer class="save-drawer" :ref="(el: any) => (saveDrawer.ref = el)" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <ElScrollbar>
            <ElForm :ref="(el: any) => (saveForm.ref = el)" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
                <ElFormItem :label="t('auth.menu.name.menuName')" prop="menuName">
                    <ElInput v-model="saveForm.data.menuName" :placeholder="t('auth.menu.name.menuName')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('auth.menu.name.sceneId')" prop="sceneId">
                    <MySelect
                        v-model="saveForm.data.sceneId"
                        :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list' }"
                        @change="
                            () => {
                                saveForm.data.pid = 0
                            }
                        "
                    />
                </ElFormItem>
                <ElFormItem v-if="saveForm.data.sceneId" :label="t('auth.menu.name.pid')" prop="pid">
                    <MyCascader
                        v-model="saveForm.data.pid"
                        :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { filter: { sceneId: saveForm.data.sceneId, excIdArr: saveForm.data.idArr } } }"
                        :props="{ checkStrictly: true, emitPath: false }"
                    />
                </ElFormItem>
                <ElFormItem :label="t('auth.menu.name.menuIcon')" prop="menuIcon">
                    <ElInput v-model="saveForm.data.menuIcon" :placeholder="t('auth.menu.name.menuIcon')" maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <label>
                        <ElAlert :title="t('auth.menu.tip.menuIcon')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem :label="t('auth.menu.name.menuUrl')" prop="menuUrl">
                    <ElInput v-model="saveForm.data.menuUrl" :placeholder="t('auth.menu.name.menuUrl')" maxlength="120" :show-word-limit="true" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('auth.menu.name.extraData')" prop="extraData">
                    <ElAlert :title="t('auth.menu.tip.extraData')" type="info" :show-icon="true" :closable="false" />
                    <ElInput v-model="saveForm.data.extraData" type="textarea" :autosize="{ minRows: 3 }" />
                </ElFormItem>
                <ElFormItem :label="t('auth.menu.name.sort')" prop="sort">
                    <ElInputNumber v-model="saveForm.data.sort" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="50" />
                    <label>
                        <ElAlert :title="t('auth.menu.tip.sort')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem :label="t('auth.menu.name.isStop')" prop="isStop">
                    <ElSwitch
                        v-model="saveForm.data.isStop"
                        :active-value="1"
                        :inactive-value="0"
                        :inline-prompt="true"
                        :active-text="t('common.yes')"
                        :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
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
