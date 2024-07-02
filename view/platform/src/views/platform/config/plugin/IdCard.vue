<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        idCardType: 'idCardOfAliyun',
        idCardOfAliyunHost: '',
        idCardOfAliyunPath: '',
        idCardOfAliyunAppcode: '',
    } as { [propName: string]: any },
    rules: {
        idCardType: [{ type: 'enum', trigger: 'change', enum: [`idCardOfAliyun`], message: t('validation.select') }],
        idCardOfAliyunHost: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        idCardOfAliyunPath: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        idCardOfAliyunAppcode: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    initData: async () => {
        const param = { config_key_arr: Object.keys(saveForm.data) }
        try {
            const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
            saveForm.data = {
                ...saveForm.data,
                ...res.data.config,
            }
        } catch (error) {
            /* eslint-disable-next-line no-empty */
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
        <el-form-item :label="t('platform.config.plugin.name.idCardType')" prop="idCardType">
            <el-radio-group v-model="saveForm.data.idCardType">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.idCardType') as any" :key="index" :value="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.idCardType == 'idCardOfAliyun'">
            <el-form-item :label="t('platform.config.plugin.name.idCardOfAliyunHost')" prop="idCardOfAliyunHost">
                <el-input v-model="saveForm.data.idCardOfAliyunHost" :placeholder="t('platform.config.plugin.name.idCardOfAliyunHost')" :clearable="true" style="max-width: 500px" />
                <el-alert type="info" :show-icon="true" :closable="false">
                    <template #title>
                        <span v-html="t('platform.config.plugin.tip.idCardOfAliyunHost')"></span>
                    </template>
                </el-alert>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.idCardOfAliyunPath')" prop="idCardOfAliyunPath">
                <el-input v-model="saveForm.data.idCardOfAliyunPath" :placeholder="t('platform.config.plugin.name.idCardOfAliyunPath')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.idCardOfAliyunAppcode')" prop="idCardOfAliyunAppcode">
                <el-input v-model="saveForm.data.idCardOfAliyunAppcode" :placeholder="t('platform.config.plugin.name.idCardOfAliyunAppcode')" :clearable="true" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isIdCardSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
