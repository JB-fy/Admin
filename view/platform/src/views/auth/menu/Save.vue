<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        sort: 50,
        ...saveCommon.data,
        sceneId: saveCommon.data.sceneId ? saveCommon.data.sceneId : undefined,
        pid: saveCommon.data.pid ? saveCommon.data.pid : undefined,
    } as { [propName: string]: any },
    rules: {
        menuName: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
        ],
        sceneId: [
            { required: true, message: t('validation.required') },
            { type: 'integer', trigger: 'change', min: 1, message: t('validation.select') },
        ],
        pid: [{ type: 'integer', trigger: 'change', min: 1, message: t('validation.select') }],
        menuIcon: [
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
            // { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],
        menuUrl: [
            { type: 'string', trigger: 'blur', max: 120, message: t('validation.max.string', { max: 120 }) },
            // { trigger: 'blur', type: 'url', message: t('validation.url') },
        ],
        extraData: [
            {
                type: 'object',
                trigger: 'blur',
                message: t('validation.json'),
                /* fields: {
                    xxxx: [
						{ type: 'string', message: 'xxxx' + t('validation.input') },
						// { type: 'integer', min: 1, message: 'xxxx' + t('validation.min.number', { min: 1 }) },
					],
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
            },
        ],
        sort: [{ type: 'integer', trigger: 'change', min: 0, max: 100, message: t('validation.between.number', { min: 0, max: 100 }) }],
        isStop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            param.sceneId === undefined ? (param.sceneId = 0) : null
            param.pid === undefined ? (param.pid = 0) : null
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
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
                <el-form-item :label="t('auth.menu.name.menuName')" prop="menuName">
                    <el-input v-model="saveForm.data.menuName" :placeholder="t('auth.menu.name.menuName')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('auth.menu.name.sceneId')" prop="sceneId">
                    <my-select v-model="saveForm.data.sceneId" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list' }" @change="() => (saveForm.data.pid = 0)" />
                </el-form-item>
                <el-form-item v-if="saveForm.data.sceneId" :label="t('auth.menu.name.pid')" prop="pid">
                    <my-cascader
                        v-model="saveForm.data.pid"
                        :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { filter: { sceneId: saveForm.data.sceneId, excIdArr: saveForm.data.idArr } } }"
                        :props="{ checkStrictly: true, emitPath: false }"
                    />
                </el-form-item>
                <el-form-item :label="t('auth.menu.name.menuIcon')" prop="menuIcon">
                    <el-input v-model="saveForm.data.menuIcon" :placeholder="t('auth.menu.name.menuIcon')" maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('auth.menu.tip.menuIcon')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('auth.menu.name.menuUrl')" prop="menuUrl">
                    <el-input v-model="saveForm.data.menuUrl" :placeholder="t('auth.menu.name.menuUrl')" maxlength="120" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('auth.menu.name.extraData')" prop="extraData">
                    <el-alert :title="t('auth.menu.tip.extraData')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
                    <el-input v-model="saveForm.data.extraData" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item>
                <el-form-item :label="t('auth.menu.name.sort')" prop="sort">
                    <el-input-number v-model="saveForm.data.sort" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="50" />
                    <el-alert :title="t('auth.menu.tip.sort')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('auth.menu.name.isStop')" prop="isStop">
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
