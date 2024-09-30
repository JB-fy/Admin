<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

let sceneConfig = saveCommon.data.scene_config ? JSON.parse(saveCommon.data.scene_config) : {}
let tokenConfig = sceneConfig?.token_config ?? {}
delete sceneConfig.token_config
const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        ...saveCommon.data,
        token_config: {
            token_type: tokenConfig.token_type ?? -1,
            expire_time: tokenConfig.expire_time ?? 14400,
            active_time: tokenConfig.active_time ?? 0,
            is_unique: tokenConfig.is_unique ?? 0,
        },
        token_config_0: {
            sign_type: tokenConfig.token_type === 0 ? tokenConfig.sign_type : 'HS256',
            private_key: tokenConfig.token_type === 0 ? tokenConfig.private_key : undefined,
            public_key: tokenConfig.token_type === 0 ? tokenConfig.public_key : undefined,
        },
        scene_config: Object.keys(sceneConfig).length > 0 ? JSON.stringify(sceneConfig) : undefined,
    } as { [propName: string]: any },
    rules: {
        scene_name: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
        ],
        scene_code: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
            { type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') },
        ],
        'token_config.token_type': [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: (tm('auth.scene.status.token_config.token_type') as any).map((item: any) => item.value), message: t('validation.select') },
        ],
        'token_config_0.sign_type': [
            { required: computed((): boolean => (saveForm.data.token_config.token_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: (tm('auth.scene.status.token_config_0.sign_type') as any).map((item: any) => item.value), message: t('validation.select') },
        ],
        'token_config_0.private_key': [
            { required: computed((): boolean => (saveForm.data.token_config.token_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'token_config_0.public_key': [
            {
                required: computed((): boolean => (saveForm.data.token_config.token_type == 0 && !['HS256', 'HS384', 'HS512'].includes(saveForm.data.token_config_0.sign_type) ? true : false)),
                message: t('validation.required'),
            },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'token_config.expire_time': [
            { required: true, message: t('validation.required') },
            { type: 'integer', trigger: 'change', min: 0, message: t('validation.min.number', { min: 0 }) },
        ],
        'token_config.active_time': [{ type: 'integer', trigger: 'change', min: 0, message: t('validation.min.number', { min: 0 }) }],
        'token_config.is_unique': [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
        scene_config: [
            {
                type: 'object',
                trigger: 'blur',
                message: t('validation.json'),
                // fields: { xxxx: [{ required: true, message: 'xxxx' + t('validation.required') }] }, //内部添加规则时，不再需要设置trigger属性
                transform: (value: any) => (value ? jsonDecode(value) : undefined),
            },
        ],
        remark: [{ type: 'string', trigger: 'blur', max: 120, message: t('validation.max.string', { max: 120 }) }],
        is_stop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            let sceneConfig = param.scene_config ? JSON.parse(param.scene_config) : {}
            if (param.token_config.token_type > -1) {
                sceneConfig.token_config = { ...param.token_config, ...param['token_config_' + param.token_config.token_type] }
                if (sceneConfig.token_config.token_type == 0 && ['HS256', 'HS384', 'HS512'].includes(sceneConfig.token_config.sign_type)) {
                    delete sceneConfig.token_config.public_key
                }
            }
            param.scene_config = Object.keys(sceneConfig).length > 0 ? JSON.stringify(sceneConfig) : ''
            try {
                if (param?.id > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/create', param, true)
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
                <el-form-item :label="t('auth.scene.name.scene_name')" prop="scene_name">
                    <el-input v-model="saveForm.data.scene_name" :placeholder="t('auth.scene.name.scene_name')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('auth.scene.name.scene_code')" prop="scene_code">
                    <el-input v-model="saveForm.data.scene_code" :placeholder="t('auth.scene.name.scene_code')" maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('auth.scene.name.token_config.token_type')" prop="token_config.token_type">
                    <el-radio-group v-model="saveForm.data.token_config.token_type">
                        <el-radio v-for="(item, index) in (tm('auth.scene.status.token_config.token_type') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
                <template v-if="saveForm.data.token_config.token_type == 0">
                    <el-form-item :label="t('auth.scene.name.token_config_0.sign_type')" prop="token_config_0.sign_type">
                        <el-select-v2 v-model="saveForm.data.token_config_0.sign_type" :options="tm('auth.scene.status.token_config_0.sign_type')" :placeholder="t('auth.scene.name.token_config_0.sign_type')" :clearable="true" style="width: 160px" />
                    </el-form-item>
                    <el-form-item :label="t('auth.scene.name.token_config_0.' + (['HS256', 'HS384', 'HS512'].includes(saveForm.data.token_config_0.sign_type) ? 'private_key_HS' : 'private_key'))" prop="token_config_0.private_key">
                        <el-input v-model="saveForm.data.token_config_0.private_key" type="textarea" :autosize="{ minRows: 3 }" />
                    </el-form-item>
                    <el-form-item v-if="!['HS256', 'HS384', 'HS512'].includes(saveForm.data.token_config_0.sign_type)" :label="t('auth.scene.name.token_config_0.public_key')" prop="token_config_0.public_key">
                        <el-input v-model="saveForm.data.token_config_0.public_key" type="textarea" :autosize="{ minRows: 3 }" />
                    </el-form-item>
                </template>
                <template v-if="saveForm.data.token_config.token_type !== -1">
                    <el-form-item :label="t('auth.scene.name.token_config.expire_time')" prop="token_config.expire_time">
                        <el-input-number v-model="saveForm.data.token_config.expire_time" :placeholder="t('auth.scene.name.token_config.expire_time')" :min="0" :precision="0" :controls="false" />
                        <el-alert :title="t('auth.scene.tip.token_config.expire_time')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <el-form-item :label="t('auth.scene.name.token_config.active_time')" prop="token_config.active_time">
                        <el-input-number v-model="saveForm.data.token_config.active_time" :placeholder="t('auth.scene.name.token_config.active_time')" :min="0" :precision="0" :controls="false" />
                        <el-alert :title="t('auth.scene.tip.token_config.active_time')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                    <el-form-item :label="t('auth.scene.name.token_config.is_unique')" prop="token_config.is_unique">
                        <el-switch
                            v-model="saveForm.data.token_config.is_unique"
                            :active-value="1"
                            :inactive-value="0"
                            :inline-prompt="true"
                            :active-text="t('common.yes')"
                            :inactive-text="t('common.no')"
                            style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success); margin-right: 10px"
                        />
                        <el-alert :title="t('auth.scene.tip.token_config.is_unique')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                </template>
                <el-form-item :label="t('auth.scene.name.scene_config')" prop="scene_config">
                    <el-alert :title="t('auth.scene.tip.scene_config')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
                    <el-input v-model="saveForm.data.scene_config" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item>
                <el-form-item :label="t('auth.scene.name.remark')" prop="remark">
                    <el-input v-model="saveForm.data.remark" type="textarea" :autosize="{ minRows: 3 }" maxlength="120" :show-word-limit="true" />
                </el-form-item>
                <el-form-item :label="t('auth.scene.name.is_stop')" prop="is_stop">
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
