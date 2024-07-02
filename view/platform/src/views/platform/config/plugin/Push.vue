<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        pushType: 'pushOfTx',
        pushOfTxHost: '',
        pushOfTxAndroidAccessID: '',
        pushOfTxAndroidSecretKey: '',
        pushOfTxIosAccessID: '',
        pushOfTxIosSecretKey: '',
        pushOfTxMacOSAccessID: '',
        pushOfTxMacOSSecretKey: '',
    } as { [propName: string]: any },
    rules: {
        pushType: [{ type: 'enum', trigger: 'change', enum: [`pushOfTx`], message: t('validation.select') }],
        pushOfTxHost: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        pushOfTxAndroidAccessID: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        pushOfTxAndroidSecretKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        pushOfTxIosAccessID: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        pushOfTxIosSecretKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        pushOfTxMacOSAccessID: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        pushOfTxMacOSSecretKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
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
            } catch (error) {
                /* eslint-disable-next-line no-empty */
            }
            saveForm.loading = false
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
            <el-form-item :label="t('platform.config.plugin.name.pushOfTxHost')" prop="pushOfTxHost">
                <el-input v-model="saveForm.data.pushOfTxHost" :placeholder="t('platform.config.plugin.name.pushOfTxHost')" :clearable="true" style="max-width: 500px" />
                <el-alert type="info" :show-icon="true" :closable="false">
                    <template #title>
                        <span v-html="t('platform.config.plugin.tip.pushOfTxHost')"></span>
                    </template>
                </el-alert>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTxAndroidAccessID')" prop="pushOfTxAndroidAccessID">
                <el-input v-model="saveForm.data.pushOfTxAndroidAccessID" :placeholder="t('platform.config.plugin.name.pushOfTxAndroidAccessID')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTxAndroidSecretKey')" prop="pushOfTxAndroidSecretKey">
                <el-input v-model="saveForm.data.pushOfTxAndroidSecretKey" :placeholder="t('platform.config.plugin.name.pushOfTxAndroidSecretKey')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTxIosAccessID')" prop="pushOfTxIosAccessID">
                <el-input v-model="saveForm.data.pushOfTxIosAccessID" :placeholder="t('platform.config.plugin.name.pushOfTxIosAccessID')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTxIosSecretKey')" prop="pushOfTxIosSecretKey">
                <el-input v-model="saveForm.data.pushOfTxIosSecretKey" :placeholder="t('platform.config.plugin.name.pushOfTxIosSecretKey')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTxMacOSAccessID')" prop="pushOfTxMacOSAccessID">
                <el-input v-model="saveForm.data.pushOfTxMacOSAccessID" :placeholder="t('platform.config.plugin.name.pushOfTxMacOSAccessID')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.pushOfTxMacOSSecretKey')" prop="pushOfTxMacOSSecretKey">
                <el-input v-model="saveForm.data.pushOfTxMacOSSecretKey" :placeholder="t('platform.config.plugin.name.pushOfTxMacOSSecretKey')" :clearable="true" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isPushSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
