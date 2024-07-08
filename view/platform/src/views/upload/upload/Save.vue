<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        upload_type: 0,
        ...saveCommon.data,
        upload_config: saveCommon.data.upload_config ? JSON.parse(saveCommon.data.upload_config) : {},
    } as { [propName: string]: any },
    rules: {
        upload_type: [{ type: 'enum', trigger: 'change', enum: (tm('upload.upload.status.upload_type') as any).map((item: any) => item.value), message: t('validation.select') }],
        /* upload_config: [
            { required: true, message: t('validation.required') },
            {
                type: 'object',
                trigger: 'blur',
                message: t('validation.json'),
                // fields: { xxxx: [{ required: true, message: 'xxxx' + t('validation.required') }] }, //内部添加规则时，不再需要设置trigger属性
                transform: (value: any) => {
                    if (!value) {
                        return undefined
                    }
                    try {
                        return JSON.parse(value)
                    } catch (error) {
                        return value
                    }
                },
            },
        ], */
        'upload_config.uploadOfLocalUrl': [
            { required: computed((): boolean => (saveForm.data.upload_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],
        'upload_config.uploadOfLocalSignKey': [
            { required: computed((): boolean => (saveForm.data.upload_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'upload_config.uploadOfLocalFileSaveDir': [
            { required: computed((): boolean => (saveForm.data.upload_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'upload_config.uploadOfLocalFileUrlPrefix': [
            { required: computed((): boolean => (saveForm.data.upload_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],
        'upload_config.uploadOfAliyunOssHost': [
            { required: computed((): boolean => (saveForm.data.upload_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],
        'upload_config.uploadOfAliyunOssBucket': [
            { required: computed((): boolean => (saveForm.data.upload_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'upload_config.uploadOfAliyunOssAccessKeyId': [
            { required: computed((): boolean => (saveForm.data.upload_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') },
        ],
        'upload_config.uploadOfAliyunOssAccessKeySecret': [
            { required: computed((): boolean => (saveForm.data.upload_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') },
        ],
        'upload_config.uploadOfAliyunOssEndpoint': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'upload_config.uploadOfAliyunOssRoleArn': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'upload_config.uploadOfAliyunOssIsNotify': [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
        remark: [{ type: 'string', trigger: 'blur', max: 120, message: t('validation.max.string', { max: 120 }) }],
        is_default: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
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
                if (param?.id_arr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/upload/upload/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/upload/upload/create', param, true)
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
                <el-form-item :label="t('upload.upload.name.upload_type')" prop="upload_type">
                    <el-radio-group v-model="saveForm.data.upload_type" @change="() => (saveForm.data.upload_config = {})">
                        <el-radio v-for="(item, index) in (tm('upload.upload.status.upload_type') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
                <!-- <el-form-item :label="t('upload.upload.name.upload_config')" prop="upload_config">
                    <el-alert :title="t('upload.upload.tip.upload_config')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
                    <el-input v-model="saveForm.data.upload_config" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item> -->
                <template v-if="saveForm.data.upload_type == 0">
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfLocalUrl')" prop="upload_config.uploadOfLocalUrl">
                        <el-input v-model="saveForm.data.upload_config.uploadOfLocalUrl" :placeholder="t('upload.upload.name.upload_config_obj.uploadOfLocalUrl')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfLocalSignKey')" prop="upload_config.uploadOfLocalSignKey">
                        <el-input v-model="saveForm.data.upload_config.uploadOfLocalSignKey" :placeholder="t('upload.upload.name.upload_config_obj.uploadOfLocalSignKey')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfLocalFileSaveDir')" prop="upload_config.uploadOfLocalFileSaveDir">
                        <el-input v-model="saveForm.data.upload_config.uploadOfLocalFileSaveDir" :placeholder="t('upload.upload.name.upload_config_obj.uploadOfLocalFileSaveDir')" :clearable="true" style="max-width: 300px" />
                        <el-alert :title="t('upload.upload.tip.upload_config_obj.uploadOfLocalFileSaveDir')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfLocalFileUrlPrefix')" prop="upload_config.uploadOfLocalFileUrlPrefix">
                        <el-input v-model="saveForm.data.upload_config.uploadOfLocalFileUrlPrefix" :placeholder="t('upload.upload.name.upload_config_obj.uploadOfLocalFileUrlPrefix')" :clearable="true" style="max-width: 300px" />
                        <el-alert :title="t('upload.upload.tip.upload_config_obj.uploadOfLocalFileUrlPrefix')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                </template>
                <template v-else-if="saveForm.data.upload_type == 1">
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssHost')" prop="upload_config.uploadOfAliyunOssHost">
                        <el-input v-model="saveForm.data.upload_config.uploadOfAliyunOssHost" :placeholder="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssHost')" :clearable="true" style="max-width: 300px" />
                        <el-alert :title="t('upload.upload.tip.upload_config_obj.uploadOfAliyunOssHost')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssBucket')" prop="upload_config.uploadOfAliyunOssBucket">
                        <el-input v-model="saveForm.data.upload_config.uploadOfAliyunOssBucket" :placeholder="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssBucket')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssAccessKeyId')" prop="upload_config.uploadOfAliyunOssAccessKeyId">
                        <el-input v-model="saveForm.data.upload_config.uploadOfAliyunOssAccessKeyId" :placeholder="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssAccessKeyId')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssAccessKeySecret')" prop="upload_config.uploadOfAliyunOssAccessKeySecret">
                        <el-input v-model="saveForm.data.upload_config.uploadOfAliyunOssAccessKeySecret" :placeholder="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssAccessKeySecret')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssEndpoint')" prop="upload_config.uploadOfAliyunOssEndpoint">
                        <el-input v-model="saveForm.data.upload_config.uploadOfAliyunOssEndpoint" :placeholder="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssEndpoint')" :clearable="true" style="max-width: 300px" />
                        <el-alert type="info" :show-icon="true" :closable="false">
                            <template #title>
                                <span v-html="t('upload.upload.tip.upload_config_obj.uploadOfAliyunOssEndpoint')"></span>
                            </template>
                        </el-alert>
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssRoleArn')" prop="upload_config.uploadOfAliyunOssRoleArn">
                        <el-input v-model="saveForm.data.upload_config.uploadOfAliyunOssRoleArn" :placeholder="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssRoleArn')" :clearable="true" style="max-width: 300px" />
                        <el-alert :title="t('upload.upload.tip.upload_config_obj.uploadOfAliyunOssRoleArn')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_obj.uploadOfAliyunOssIsNotify')" prop="upload_config.uploadOfAliyunOssIsNotify">
                        <el-switch
                            v-model="saveForm.data.upload_config.uploadOfAliyunOssIsNotify"
                            :active-value="1"
                            :inactive-value="0"
                            :inline-prompt="true"
                            :active-text="t('common.yes')"
                            :inactive-text="t('common.no')"
                            style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                        />
                    </el-form-item>
                </template>
                <el-form-item :label="t('upload.upload.name.remark')" prop="remark">
                    <el-input v-model="saveForm.data.remark" type="textarea" :autosize="{ minRows: 3 }" maxlength="120" :show-word-limit="true" />
                </el-form-item>
                <el-form-item :label="t('upload.upload.name.is_default')" prop="is_default">
                    <el-switch
                        v-model="saveForm.data.is_default"
                        :active-value="1"
                        :inactive-value="0"
                        :inline-prompt="true"
                        :active-text="t('common.yes')"
                        :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </el-form-item>
                <el-form-item :label="t('upload.upload.name.is_stop')" prop="is_stop">
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