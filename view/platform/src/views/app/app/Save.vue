<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        app_type: 0,
        ...saveCommon.data,
        extra_config: saveCommon.data.extra_config ? JSON.parse(saveCommon.data.extra_config) : {},
    } as { [propName: string]: any },
    rules: {
        app_type: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: (tm('app.app.status.app_type') as any).map((item: any) => item.value), message: t('validation.select') },
        ],
        package_name: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 60, message: t('validation.max.string', { max: 60 }) },
        ],
        package_file: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 200, message: t('validation.max.string', { max: 200 }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],
        ver_no: [
            { required: true, message: t('validation.required') },
            { type: 'integer', trigger: 'change', min: 0, max: 4294967295, message: t('validation.between.number', { min: 0, max: 4294967295 }) },
        ],
        ver_name: [{ type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) }],
        ver_intro: [{ type: 'string', trigger: 'blur', max: 255, message: t('validation.max.string', { max: 255 }) }],
        /* extra_config: [
            {
                type: 'object',
                trigger: 'blur',
                message: t('validation.json'),
                // fields: { xxxx: [{ required: true, message: 'xxxx' + t('validation.required') }] }, //内部添加规则时，不再需要设置trigger属性
                transform: (value: any) => (value ? jsonDecode(value) : undefined),
            },
        ], */
        'extra_config.marketUrl': [
            // { required: computed((): boolean => (saveForm.data.app_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'extra_config.plistFile': [
            // { required: computed((): boolean => (saveForm.data.app_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],
        remark: [{ type: 'string', trigger: 'blur', max: 255, message: t('validation.max.string', { max: 255 }) }],
        is_force_prev: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
        is_stop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            try {
                if (param?.id) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/app/app/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/app/app/create', param, true)
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
                <el-form-item :label="t('app.app.name.app_type')" prop="app_type">
                    <el-radio-group v-model="saveForm.data.app_type" @change="() => (saveForm.data.extra_config = {})">
                        <el-radio v-for="(item, index) in (tm('app.app.status.app_type') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item :label="t('app.app.name.package_name')" prop="package_name">
                    <el-input v-model="saveForm.data.package_name" :placeholder="t('app.app.name.package_name')" maxlength="60" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('app.app.name.package_file')" prop="package_file">
                    <my-upload v-model="saveForm.data.package_file" accept="application/*" />
                </el-form-item>
                <el-form-item :label="t('app.app.name.ver_no')" prop="ver_no">
                    <el-input-number v-model="saveForm.data.ver_no" :placeholder="t('app.app.name.ver_no')" :min="0" :max="4294967295" :precision="0" :value-on-clear="0" />
                </el-form-item>
                <el-form-item :label="t('app.app.name.ver_name')" prop="ver_name">
                    <el-input v-model="saveForm.data.ver_name" :placeholder="t('app.app.name.ver_name')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('app.app.name.ver_intro')" prop="ver_intro">
                    <el-input v-model="saveForm.data.ver_intro" type="textarea" :autosize="{ minRows: 3 }" maxlength="255" :show-word-limit="true" />
                </el-form-item>
                <!-- <el-form-item :label="t('app.app.name.extra_config')" prop="extra_config">
                    <el-input v-model="saveForm.data.extra_config" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item> -->
                <template v-if="saveForm.data.app_type == 1">
                    <el-form-item :label="t('app.app.name.extra_config_obj.marketUrl')" prop="extra_config.marketUrl">
                        <el-input v-model="saveForm.data.extra_config.marketUrl" :placeholder="t('app.app.name.extra_config_obj.marketUrl')" :clearable="true" style="max-width: 400px" />
                        <el-alert :title="t('app.app.tip.extra_config_obj.marketUrl')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <el-form-item :label="t('app.app.name.extra_config_obj.plistFile')" prop="extra_config.plistFile">
                        <my-upload v-model="saveForm.data.extra_config.plistFile">
                            <template #tip>
                                <el-alert :title="t('app.app.tip.extra_config_obj.plistFile')" type="info" :show-icon="true" :closable="false" />
                            </template>
                        </my-upload>
                    </el-form-item>
                </template>
                <el-form-item :label="t('app.app.name.remark')" prop="remark">
                    <el-input v-model="saveForm.data.remark" type="textarea" :autosize="{ minRows: 3 }" maxlength="255" :show-word-limit="true" />
                </el-form-item>
                <el-form-item :label="t('app.app.name.is_force_prev')" prop="is_force_prev">
                    <el-switch
                        v-model="saveForm.data.is_force_prev"
                        :active-value="1"
                        :inactive-value="0"
                        :inline-prompt="true"
                        :active-text="t('common.yes')"
                        :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                    <el-alert :title="t('app.app.tip.is_force_prev')" type="info" :show-icon="true" :closable="false" style="margin-left: 10px" />
                </el-form-item>
                <el-form-item :label="t('app.app.name.is_stop')" prop="is_stop">
                    <el-switch
                        v-model="saveForm.data.is_stop"
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
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">{{ t('common.save') }}</el-button>
        </template>
    </el-drawer>
</template>
