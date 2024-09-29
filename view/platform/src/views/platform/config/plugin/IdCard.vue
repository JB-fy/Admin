<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        idCardType: 'idCardOfAliyun',
        idCardOfAliyun: {},
    } as { [propName: string]: any },
    rules: {
        idCardType: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: [`idCardOfAliyun`], message: t('validation.select') },
        ],
        'idCardOfAliyun.host': [
            { required: computed((): boolean => (saveForm.data.idCardType == `idCardOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],
        'idCardOfAliyun.path': [
            { required: computed((): boolean => (saveForm.data.idCardType == `idCardOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'idCardOfAliyun.appcode': [
            { required: computed((): boolean => (saveForm.data.idCardType == `idCardOfAliyun` ? true : false)), message: t('validation.required') },
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
        <el-form-item :label="t('platform.config.plugin.name.idCardType')" prop="idCardType">
            <el-radio-group v-model="saveForm.data.idCardType">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.idCardType') as any" :key="index" :value="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.idCardType == 'idCardOfAliyun'">
            <el-form-item :label="t('platform.config.plugin.name.idCardOfAliyun.host')" prop="idCardOfAliyun.host">
                <el-input v-model="saveForm.data.idCardOfAliyun.host" :placeholder="t('platform.config.plugin.name.idCardOfAliyun.host')" :clearable="true" style="max-width: 500px" />
                <el-alert type="info" :show-icon="true" :closable="false">
                    <template #title>
                        <span v-html="t('platform.config.plugin.tip.idCardOfAliyun.host')"></span>
                    </template>
                </el-alert>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.idCardOfAliyun.path')" prop="idCardOfAliyun.path">
                <el-input v-model="saveForm.data.idCardOfAliyun.path" :placeholder="t('platform.config.plugin.name.idCardOfAliyun.path')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.idCardOfAliyun.appcode')" prop="idCardOfAliyun.appcode">
                <el-input v-model="saveForm.data.idCardOfAliyun.appcode" :placeholder="t('platform.config.plugin.name.idCardOfAliyun.appcode')" :clearable="true" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isIdCardSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
