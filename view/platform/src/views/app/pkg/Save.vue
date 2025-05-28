<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const extraConfig = saveCommon.data.extra_config ? JSON.parse(saveCommon.data.extra_config) : {}
const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        pkg_type: 0,
        ...saveCommon.data,
        app_id: saveCommon.data.app_id ? saveCommon.data.app_id : undefined,
        extra_config_1: (() => {
            const extraConfig1: { [propName: string]: any } = {}
            if (saveCommon.data.pkg_type == 1) {
                extraConfig1.pkg_source = extraConfig.pkg_source ?? 0
                delete extraConfig.pkg_source
                extraConfig1.market_url = extraConfig.market_url
                delete extraConfig.market_url
                extraConfig1.qyq_h5_url = extraConfig.qyq_h5_url
                delete extraConfig.qyq_h5_url
                extraConfig1.qyq_plist_file = extraConfig.qyq_plist_file
                delete extraConfig.qyq_plist_file
            }
            return extraConfig1
        })(),
        extra_config: Object.keys(extraConfig).length > 0 ? JSON.stringify(extraConfig) : undefined,
    } as { [propName: string]: any },
    rules: {
        app_id: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'change', max: 15, message: t('validation.select') },
        ],
        pkg_type: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: (tm('app.pkg.status.pkg_type') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') },
        ],
        pkg_name: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 60, message: t('validation.max.string', { max: 60 }) },
        ],
        pkg_file: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 200, message: t('validation.max.string', { max: 200 }) },
            { type: 'url', trigger: computed((): string => (saveForm.data.is_input_pkg_file ? 'blur' : 'change')), message: computed((): string => (saveForm.data.is_input_pkg_file ? t('validation.url') : t('validation.upload'))) },
        ],
        ver_no: [
            { required: true, message: t('validation.required') },
            { type: 'integer', trigger: 'change', min: 0, max: 4294967295, message: t('validation.between.number', { min: 0, max: 4294967295 }) },
        ],
        ver_name: [{ type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) }],
        ver_intro: [{ type: 'string', trigger: 'blur', max: 255, message: t('validation.max.string', { max: 255 }) }],
        'extra_config_1.pkg_source': [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: (tm('app.pkg.status.extra_config_1.pkg_source') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') },
        ],
        'extra_config_1.market_url': [
            { required: computed((): boolean => (saveForm.data.pkg_type == 1 && saveForm.data.extra_config_1.pkg_source == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'extra_config_1.qyq_h5_url': [
            { required: computed((): boolean => (saveForm.data.pkg_type == 1 && saveForm.data.extra_config_1.pkg_source == 1 ? true : false)), message: t('validation.required') },
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],
        'extra_config_1.qyq_plist_file': [
            { required: computed((): boolean => (saveForm.data.pkg_type == 1 && saveForm.data.extra_config_1.pkg_source == 1 ? true : false)), message: t('validation.required') },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],
        extra_config: [
            {
                type: 'object',
                trigger: 'blur',
                message: t('validation.json'),
                // fields: { xxxx: [{ required: true, message: 'xxxx' + t('validation.required') }] }, //内部添加规则时，不再需要设置trigger属性
                transform: (value: any) => (value ? jsonDecode(value) : undefined),
            },
        ],
        remark: [{ type: 'string', trigger: 'blur', max: 120, message: t('validation.max.string', { max: 120 }) }],
        is_force_prev: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
        is_stop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            param.app_id === undefined && (param.app_id = 0)
            let extraConfig = param.extra_config ? JSON.parse(param.extra_config) : {}
            extraConfig = { ...extraConfig, ...param['extra_config_' + param.pkg_type] }
            param.extra_config = Object.keys(extraConfig).length > 0 ? JSON.stringify(extraConfig) : ''
            try {
                if (param?.id) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/app/pkg/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/app/pkg/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } finally {
                saveForm.loading = false
            }
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
            }).then(() => done())
        } else {
            done()
        }
    },
    buttonClose: () => saveDrawer.ref.handleClose(), //saveCommon.visible = false //不会触发beforeClose
})
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
                <el-form-item :label="t('app.pkg.name.app_id')" prop="app_id">
                    <my-select v-model="saveForm.data.app_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/app/app/list' }" />
                </el-form-item>
                <el-form-item :label="t('app.pkg.name.pkg_type')" prop="pkg_type">
                    <el-radio-group v-model="saveForm.data.pkg_type">
                        <el-radio v-for="(item, index) in (tm('app.pkg.status.pkg_type') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item :label="t('app.pkg.name.pkg_name')" prop="pkg_name">
                    <el-input v-model="saveForm.data.pkg_name" :placeholder="t('app.pkg.name.pkg_name')" maxlength="60" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('app.pkg.name.pkg_file')" prop="pkg_file">
                    <div style="width: 100%">
                        <el-checkbox v-model="saveForm.data.is_input_pkg_file" :label="t('app.pkg.name.is_input_pkg_file')" />
                    </div>
                    <el-input v-if="saveForm.data.is_input_pkg_file" v-model="saveForm.data.pkg_file" :placeholder="t('app.pkg.name.pkg_file')" maxlength="200" :show-word-limit="true" :clearable="true" />
                    <my-upload v-else v-model="saveForm.data.pkg_file" accept="application/*" />
                </el-form-item>
                <el-form-item :label="t('app.pkg.name.ver_no')" prop="ver_no">
                    <el-input-number v-model="saveForm.data.ver_no" :placeholder="t('app.pkg.name.ver_no')" :min="0" :max="4294967295" :precision="0" :value-on-clear="0" />
                </el-form-item>
                <el-form-item :label="t('app.pkg.name.ver_name')" prop="ver_name">
                    <el-input v-model="saveForm.data.ver_name" :placeholder="t('app.pkg.name.ver_name')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('app.pkg.name.ver_intro')" prop="ver_intro">
                    <el-input v-model="saveForm.data.ver_intro" type="textarea" :autosize="{ minRows: 3 }" maxlength="255" :show-word-limit="true" />
                </el-form-item>
                <template v-if="saveForm.data.pkg_type == 1">
                    <el-form-item :label="t('app.pkg.name.extra_config_1.pkg_source')" prop="extra_config_1.pkg_source">
                        <el-radio-group v-model="saveForm.data.extra_config_1.pkg_source">
                            <el-radio v-for="(item, index) in (tm('app.pkg.status.extra_config_1.pkg_source') as any)" :key="index" :value="item.value">
                                {{ item.label }}
                            </el-radio>
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item v-if="saveForm.data.extra_config_1.pkg_source == 0" :label="t('app.pkg.name.extra_config_1.market_url')" prop="extra_config_1.market_url">
                        <el-input v-model="saveForm.data.extra_config_1.market_url" :placeholder="t('app.pkg.name.extra_config_1.market_url')" :clearable="true" style="max-width: 400px" />
                        <el-alert :title="t('app.pkg.tip.extra_config_1.market_url')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <template v-else-if="saveForm.data.extra_config_1.pkg_source == 1">
                        <el-form-item :label="t('app.pkg.name.extra_config_1.qyq_h5_url')" prop="extra_config_1.qyq_h5_url">
                            <el-input v-model="saveForm.data.extra_config_1.qyq_h5_url" :placeholder="t('app.pkg.name.extra_config_1.qyq_h5_url')" :clearable="true" style="max-width: 400px" />
                            <el-alert :title="t('app.pkg.tip.extra_config_1.qyq_h5_url')" type="info" :show-icon="true" :closable="false" />
                        </el-form-item>
                        <el-form-item :label="t('app.pkg.name.extra_config_1.qyq_plist_file')" prop="extra_config_1.qyq_plist_file">
                            <div style="width: 100%">
                                <el-checkbox v-model="saveForm.data.is_qyq_plist_file" :label="t('app.pkg.name.extra_config_1.is_qyq_plist_file')" />
                            </div>
                            <template v-if="saveForm.data.is_qyq_plist_file">
                                <el-input v-model="saveForm.data.extra_config_1.qyq_plist_file" :placeholder="t('app.pkg.name.extra_config_1.qyq_plist_file')" :clearable="true" style="max-width: 600px" />
                                <el-alert :title="t('app.pkg.tip.extra_config_1.qyq_plist_file')" type="info" :show-icon="true" :closable="false" />
                            </template>
                            <my-upload v-else v-model="saveForm.data.extra_config_1.qyq_plist_file">
                                <template #tip>
                                    <el-alert :title="t('app.pkg.tip.extra_config_1.qyq_plist_file')" type="info" :show-icon="true" :closable="false" />
                                </template>
                            </my-upload>
                        </el-form-item>
                    </template>
                </template>
                <el-form-item :label="t('app.pkg.name.extra_config')" prop="extra_config">
                    <el-alert :title="t('app.pkg.tip.extra_config')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
                    <el-input v-model="saveForm.data.extra_config" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item>
                <el-form-item :label="t('app.pkg.name.remark')" prop="remark">
                    <el-input v-model="saveForm.data.remark" type="textarea" :autosize="{ minRows: 3 }" maxlength="120" :show-word-limit="true" />
                </el-form-item>
                <el-form-item :label="t('app.pkg.name.is_force_prev')" prop="is_force_prev">
                    <el-switch
                        v-model="saveForm.data.is_force_prev"
                        :active-value="(tm('common.status.whether') as any[])[1].value"
                        :inactive-value="(tm('common.status.whether') as any[])[0].value"
                        :active-text="(tm('common.status.whether') as any[])[1].label"
                        :inactive-text="(tm('common.status.whether') as any[])[0].label"
                        :inline-prompt="true"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                    <el-alert :title="t('app.pkg.tip.is_force_prev')" type="info" :show-icon="true" :closable="false" style="margin-left: 10px" />
                </el-form-item>
                <el-form-item :label="t('app.pkg.name.is_stop')" prop="is_stop">
                    <el-switch
                        v-model="saveForm.data.is_stop"
                        :active-value="(tm('common.status.whether') as any[])[1].value"
                        :inactive-value="(tm('common.status.whether') as any[])[0].value"
                        :active-text="(tm('common.status.whether') as any[])[1].label"
                        :inactive-text="(tm('common.status.whether') as any[])[0].label"
                        :inline-prompt="true"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </el-form-item>
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">{{ t('common.save') }}</el-button>
        </template>
    </el-drawer>
</template>
