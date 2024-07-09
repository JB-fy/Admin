<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        pushType: 'pushOfTx',
        pushOfTx: {},
    } as { [propName: string]: any },
    rules: {
        pushType: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: [`pushOfTx`], message: t('validation.select') },
        ],
        'pushOfTx.host': [
            { required: computed((): boolean => (saveForm.data.pushType == `pushOfTx` ? true : false)), message: t('validation.required') },
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],
        'pushOfTx.accessIDOfAndroid': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'pushOfTx.secretKeyOfAndroid': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'pushOfTx.accessIDOfIos': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'pushOfTx.secretKeyOfIos': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'pushOfTx.accessIDOfMacOS': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'pushOfTx.secretKeyOfMacOS': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    initData: async () => {
        const param = { config_key_arr: Object.keys(saveForm.data) }
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
        saveForm.data = {
            ...saveForm.data,
            ...res.data.config,
        }
    },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            try {
                await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/save', param, true)
            } finally {
                saveForm.loading = false
            }
        })
    },
    reset: () => {
        saveForm.ref.resetFields()
        saveForm.initData()
    },
})

saveForm.initData()
</script>

<template>
    <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
        <el-form-item :label="t('platform.config.plugin.name.pushType')" prop="pushType">
            <el-radio-group v-model="saveForm.data.pushType">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.pushType') as any" :key="index" :value="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.pushType == 'pushOfTx'">
            <el-form-item :label="t('platform.config.plugin.name.pushOfTx.host')" prop="pushOfTx.host">
                <el-input v-model="saveForm.data.pushOfTx.host" :placeholder="t('platform.config.plugin.name.pushOfTx.host')" :clearable="true" style="max-width: 500px" />
                <el-alert type="info" :show-icon="true" :closable="false">
                    <template #title>
                        <span v-html="t('platform.config.plugin.tip.pushOfTx.host')"></span>
                    </template>
                </el-alert>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTx.accessIDOfAndroid')" prop="pushOfTx.accessIDOfAndroid">
                <el-input v-model="saveForm.data.pushOfTx.accessIDOfAndroid" :placeholder="t('platform.config.plugin.name.pushOfTx.accessIDOfAndroid')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTx.secretKeyOfAndroid')" prop="pushOfTx.secretKeyOfAndroid">
                <el-input v-model="saveForm.data.pushOfTx.secretKeyOfAndroid" :placeholder="t('platform.config.plugin.name.pushOfTx.secretKeyOfAndroid')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTx.accessIDOfIos')" prop="pushOfTx.accessIDOfIos">
                <el-input v-model="saveForm.data.pushOfTx.accessIDOfIos" :placeholder="t('platform.config.plugin.name.pushOfTx.accessIDOfIos')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTx.secretKeyOfIos')" prop="pushOfTx.secretKeyOfIos">
                <el-input v-model="saveForm.data.pushOfTx.secretKeyOfIos" :placeholder="t('platform.config.plugin.name.pushOfTx.secretKeyOfIos')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTx.accessIDOfMacOS')" prop="pushOfTx.accessIDOfMacOS">
                <el-input v-model="saveForm.data.pushOfTx.accessIDOfMacOS" :placeholder="t('platform.config.plugin.name.pushOfTx.accessIDOfMacOS')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTx.secretKeyOfMacOS')" prop="pushOfTx.secretKeyOfMacOS">
                <el-input v-model="saveForm.data.pushOfTx.secretKeyOfMacOS" :placeholder="t('platform.config.plugin.name.pushOfTx.secretKeyOfMacOS')" :clearable="true" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isPushSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
