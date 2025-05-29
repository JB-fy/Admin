<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        vodType: 'vodOfAliyun',
        vodOfAliyun: {},
    } as { [propName: string]: any },
    rules: {
        vodType: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: [`vodOfAliyun`], message: t('validation.select') },
        ],
        'vodOfAliyun.access_key_id': [
            { required: computed((): boolean => (saveForm.data.vodType == `vodOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'vodOfAliyun.access_key_secret': [
            { required: computed((): boolean => (saveForm.data.vodType == `vodOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'vodOfAliyun.endpoint': [
            { required: computed((): boolean => (saveForm.data.vodType == `vodOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'vodOfAliyun.roleArn': [
            { required: computed((): boolean => (saveForm.data.vodType == `vodOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
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
        <el-form-item :label="t('platform.config.plugin.name.vodType')" prop="vodType">
            <el-radio-group v-model="saveForm.data.vodType">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.vodType') as any" :key="index" :value="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.vodType == 'vodOfAliyun'">
            <el-form-item :label="t('platform.config.plugin.name.vodOfAliyun.access_key_id')" prop="vodOfAliyun.access_key_id">
                <el-input v-model="saveForm.data.vodOfAliyun.access_key_id" :placeholder="t('platform.config.plugin.name.vodOfAliyun.access_key_id')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.vodOfAliyun.access_key_secret')" prop="vodOfAliyun.access_key_secret">
                <el-input v-model="saveForm.data.vodOfAliyun.access_key_secret" :placeholder="t('platform.config.plugin.name.vodOfAliyun.access_key_secret')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.vodOfAliyun.endpoint')" prop="vodOfAliyun.endpoint">
                <el-input v-model="saveForm.data.vodOfAliyun.endpoint" :placeholder="t('platform.config.plugin.name.vodOfAliyun.endpoint')" :clearable="true" style="max-width: 500px" />
                <el-alert type="info" :show-icon="true" :closable="false">
                    <template #title>
                        <span v-html="t('platform.config.plugin.tip.vodOfAliyun.endpoint')"></span>
                    </template>
                </el-alert>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.vodOfAliyun.roleArn')" prop="vodOfAliyun.roleArn">
                <el-input v-model="saveForm.data.vodOfAliyun.roleArn" :placeholder="t('platform.config.plugin.name.vodOfAliyun.roleArn')" :clearable="true" style="max-width: 500px" />
                <el-alert :title="t('platform.config.plugin.tip.vodOfAliyun.roleArn')" type="info" :show-icon="true" :closable="false" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isVodSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
