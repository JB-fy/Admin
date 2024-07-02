<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        vodType: 'vodOfAliyun',
        vodOfAliyunAccessKeyId: '',
        vodOfAliyunAccessKeySecret: '',
        vodOfAliyunEndpoint: '',
        vodOfAliyunRoleArn: '',
    } as { [propName: string]: any },
    rules: {
        vodType: [{ type: 'enum', trigger: 'change', enum: [`vodOfAliyun`], message: t('validation.select') }],
        vodOfAliyunAccessKeyId: [{ type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') }],
        vodOfAliyunAccessKeySecret: [{ type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') }],
        vodOfAliyunEndpoint: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        vodOfAliyunRoleArn: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
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
            <el-form-item :label="t('platform.config.plugin.name.vodOfAliyunAccessKeyId')" prop="vodOfAliyunAccessKeyId">
                <el-input v-model="saveForm.data.vodOfAliyunAccessKeyId" :placeholder="t('platform.config.plugin.name.vodOfAliyunAccessKeyId')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.vodOfAliyunAccessKeySecret')" prop="vodOfAliyunAccessKeySecret">
                <el-input v-model="saveForm.data.vodOfAliyunAccessKeySecret" :placeholder="t('platform.config.plugin.name.vodOfAliyunAccessKeySecret')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.vodOfAliyunEndpoint')" prop="vodOfAliyunEndpoint">
                <el-input v-model="saveForm.data.vodOfAliyunEndpoint" :placeholder="t('platform.config.plugin.name.vodOfAliyunEndpoint')" :clearable="true" style="max-width: 500px" />
                <el-alert type="info" :show-icon="true" :closable="false">
                    <template #title>
                        <span v-html="t('platform.config.plugin.tip.vodOfAliyunEndpoint')"></span>
                    </template>
                </el-alert>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.vodOfAliyunRoleArn')" prop="vodOfAliyunRoleArn">
                <el-input v-model="saveForm.data.vodOfAliyunRoleArn" :placeholder="t('platform.config.plugin.name.vodOfAliyunRoleArn')" :clearable="true" style="max-width: 500px" />
                <el-alert :title="t('platform.config.plugin.tip.vodOfAliyunRoleArn')" type="info" :show-icon="true" :closable="false" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isVodSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
