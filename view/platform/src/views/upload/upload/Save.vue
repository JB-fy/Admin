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
        upload_config_0: (() => {
            const uploadConfig = saveCommon.data.upload_type == 0 && saveCommon.data.upload_config ? JSON.parse(saveCommon.data.upload_config) : {}
            uploadConfig.serverList || (uploadConfig.serverList = [])
            return uploadConfig
        })(),
        upload_config_1: saveCommon.data.upload_type == 1 && saveCommon.data.upload_config ? JSON.parse(saveCommon.data.upload_config) : {},
    } as { [propName: string]: any },
    rules: {
        upload_type: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: (tm('upload.upload.status.upload_type') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') },
        ],
        /* upload_config: [
            { required: true, message: t('validation.required') },
            {
                type: 'object',
                trigger: 'blur',
                message: t('validation.json'),
                // fields: { xxxx: [{ required: true, message: 'xxxx' + t('validation.required') }] }, //内部添加规则时，不再需要设置trigger属性
                transform: (value: any) => (value ? jsonDecode(value) : undefined),
            },
        ], */
        'upload_config_0.signKey': [
            { required: computed((): boolean => (saveForm.data.upload_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'upload_config_0.url': [{ type: 'url', trigger: 'blur', message: t('validation.input') }],
        'upload_config_0.fileSaveDir': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'upload_config_0.isCluster': [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
        'upload_config_0.serverList': [
            {
                type: 'array',
                trigger: 'change',
                defaultField: [
                    {
                        type: 'object',
                        fields: {
                            ip: [
                                { required: true, message: t('upload.upload.name.upload_config_0.serverList.ip') + t('validation.required') },
                                {
                                    type: 'string',
                                    trigger: 'blur',
                                    pattern:
                                        /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(?:[A-Fa-f0-9]{1,4}:){7}[A-Fa-f0-9]{1,4}$|^((?:[A-Fa-f0-9]{1,4}:){0,6}[A-Fa-f0-9]{1,4})?::((?:[A-Fa-f0-9]{1,4}:){0,6}[A-Fa-f0-9]{1,4})?$/,
                                    message: t('upload.upload.name.upload_config_0.serverList.ip') + t('validation.ip'),
                                },
                            ],
                            host: [
                                { required: true, message: t('upload.upload.name.upload_config_0.serverList.host') + t('validation.required') },
                                { type: 'string', trigger: 'blur', pattern: /^([0-9a-zA-Z][0-9a-zA-Z\-]{0,62}\.)+([a-zA-Z]{0,62})$/, message: t('upload.upload.name.upload_config_0.serverList.host') + t('validation.url') },
                            ],
                        },
                    },
                ],
            },
        ],
        'upload_config_0.isSameServer': [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
        'upload_config_1.host': [
            { required: computed((): boolean => (saveForm.data.upload_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],
        'upload_config_1.bucket': [
            { required: computed((): boolean => (saveForm.data.upload_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'upload_config_1.accessKeyId': [
            { required: computed((): boolean => (saveForm.data.upload_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') },
        ],
        'upload_config_1.accessKeySecret': [
            { required: computed((): boolean => (saveForm.data.upload_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') },
        ],
        'upload_config_1.endpoint': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'upload_config_1.roleArn': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'upload_config_1.isNotify': [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
        remark: [{ type: 'string', trigger: 'blur', max: 120, message: t('validation.max.string', { max: 120 }) }],
        is_default: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            param.upload_config = param['upload_config_' + param.upload_type]
            try {
                if (param?.id) {
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
                    <el-radio-group v-model="saveForm.data.upload_type">
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
                    <el-form-item :label="t('upload.upload.name.upload_config_0.signKey')" prop="upload_config_0.signKey">
                        <el-input v-model="saveForm.data.upload_config_0.signKey" :placeholder="t('upload.upload.name.upload_config_0.signKey')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_0.url')" prop="upload_config_0.url">
                        <el-input v-model="saveForm.data.upload_config_0.url" :placeholder="t('upload.upload.name.upload_config_0.url')" :clearable="true" style="max-width: 300px" />
                        <el-alert :title="t('upload.upload.tip.upload_config_0.url')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_0.fileSaveDir')" prop="upload_config_0.fileSaveDir">
                        <el-input v-model="saveForm.data.upload_config_0.fileSaveDir" :placeholder="t('upload.upload.name.upload_config_0.fileSaveDir')" :clearable="true" style="max-width: 300px" />
                        <el-alert :title="t('upload.upload.tip.upload_config_0.fileSaveDir')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_0.isCluster')" prop="upload_config_0.isCluster">
                        <el-switch
                            v-model="saveForm.data.upload_config_0.isCluster"
                            :active-value="(tm('common.status.whether') as any[])[1].value"
                            :inactive-value="(tm('common.status.whether') as any[])[0].value"
                            :active-text="(tm('common.status.whether') as any[])[1].label"
                            :inactive-text="(tm('common.status.whether') as any[])[0].label"
                            :inline-prompt="true"
                            style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                        />
                    </el-form-item>
                    <template v-if="saveForm.data.upload_config_0.isCluster">
                        <el-form-item :label="t('upload.upload.name.upload_config_0.serverList_label')" prop="upload_config_0.serverList" style="min-height: 60px">
                            <template #label>
                                <span style="text-align: right">
                                    <div>{{ t('upload.upload.name.upload_config_0.serverList_label') }}</div>
                                    <el-button type="primary" size="small" @click="() => saveForm.data.upload_config_0.serverList.push({})"><autoicon-ep-plus />{{ t('common.add') }}</el-button>
                                </span>
                            </template>
                            <el-alert type="info" :show-icon="true" :closable="false" style="width: 100%">
                                <template #title>
                                    <span v-html="t('upload.upload.tip.upload_config_0.serverList')"></span>
                                </template>
                            </el-alert>
                            <template v-for="(_, index) in saveForm.data.upload_config_0.serverList" :key="index">
                                <div style="width: 100%; margin: 3px 0; display: flex; align-items: center; gap: 10px">
                                    <el-button type="danger" size="small" @click="() => saveForm.data.upload_config_0.serverList.splice(index, 1)"><autoicon-ep-close />{{ t('common.delete') }}</el-button>
                                    <el-input v-model="saveForm.data.upload_config_0.serverList[index].ip" :placeholder="t('upload.upload.name.upload_config_0.serverList.ip')" :clearable="true" style="max-width: 200px" />
                                    <el-input v-model="saveForm.data.upload_config_0.serverList[index].host" :placeholder="t('upload.upload.name.upload_config_0.serverList.host')" :clearable="true" />
                                </div>
                            </template>
                        </el-form-item>
                        <el-form-item :label="t('upload.upload.name.upload_config_0.isSameServer')" prop="upload_config_0.isSameServer">
                            <el-switch
                                v-model="saveForm.data.upload_config_0.isSameServer"
                                :active-value="(tm('common.status.whether') as any[])[1].value"
                                :inactive-value="(tm('common.status.whether') as any[])[0].value"
                                :active-text="(tm('common.status.whether') as any[])[1].label"
                                :inactive-text="(tm('common.status.whether') as any[])[0].label"
                                :inline-prompt="true"
                                style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                            />
                            <el-alert :title="t('upload.upload.tip.upload_config_0.isSameServer')" type="info" :show-icon="true" :closable="false" />
                        </el-form-item>
                    </template>
                </template>
                <template v-else-if="saveForm.data.upload_type == 1">
                    <el-form-item :label="t('upload.upload.name.upload_config_1.host')" prop="upload_config_1.host">
                        <el-input v-model="saveForm.data.upload_config_1.host" :placeholder="t('upload.upload.name.upload_config_1.host')" :clearable="true" style="max-width: 300px" />
                        <el-alert :title="t('upload.upload.tip.upload_config_1.host')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_1.bucket')" prop="upload_config_1.bucket">
                        <el-input v-model="saveForm.data.upload_config_1.bucket" :placeholder="t('upload.upload.name.upload_config_1.bucket')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_1.accessKeyId')" prop="upload_config_1.accessKeyId">
                        <el-input v-model="saveForm.data.upload_config_1.accessKeyId" :placeholder="t('upload.upload.name.upload_config_1.accessKeyId')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_1.accessKeySecret')" prop="upload_config_1.accessKeySecret">
                        <el-input v-model="saveForm.data.upload_config_1.accessKeySecret" :placeholder="t('upload.upload.name.upload_config_1.accessKeySecret')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_1.endpoint')" prop="upload_config_1.endpoint">
                        <el-input v-model="saveForm.data.upload_config_1.endpoint" :placeholder="t('upload.upload.name.upload_config_1.endpoint')" :clearable="true" style="max-width: 300px" />
                        <el-alert type="info" :show-icon="true" :closable="false">
                            <template #title>
                                <span v-html="t('upload.upload.tip.upload_config_1.endpoint')"></span>
                            </template>
                        </el-alert>
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_1.roleArn')" prop="upload_config_1.roleArn">
                        <el-input v-model="saveForm.data.upload_config_1.roleArn" :placeholder="t('upload.upload.name.upload_config_1.roleArn')" :clearable="true" style="max-width: 300px" />
                        <el-alert :title="t('upload.upload.tip.upload_config_1.roleArn')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <el-form-item :label="t('upload.upload.name.upload_config_1.isNotify')" prop="upload_config_1.isNotify">
                        <el-switch
                            v-model="saveForm.data.upload_config_1.isNotify"
                            :active-value="(tm('common.status.whether') as any[])[1].value"
                            :inactive-value="(tm('common.status.whether') as any[])[0].value"
                            :active-text="(tm('common.status.whether') as any[])[1].label"
                            :inactive-text="(tm('common.status.whether') as any[])[0].label"
                            :inline-prompt="true"
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
