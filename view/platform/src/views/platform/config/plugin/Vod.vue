<script setup lang="tsx">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        vodType: 'vodOfAliyun',
        vodOfAliyunAccessKeyId: '',
        vodOfAliyunAccessKeySecret: '',
        vodOfAliyunEndpoint: '',
        vodOfAliyunRoleArn: '',
    } as { [propName: string]: any },
    rules: {
        vodType: [{ type: 'enum', enum: [`vodOfAliyun`], trigger: 'change', message: t('validation.select') }],
        vodOfAliyunAccessKeyId: [{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }],
        vodOfAliyunAccessKeySecret: [{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }],
        vodOfAliyunEndpoint: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        vodOfAliyunRoleArn: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
    } as any,
    initData: async () => {
        const param = { configKeyArr: Object.keys(saveForm.data) }
        try {
            const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
            saveForm.data = {
                ...saveForm.data,
                ...res.data.config,
            }
        } catch (error) {}
    },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data, false)
            try {
                await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/save', param, true)
            } catch (error) {}
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
        <el-form-item :label="t('platform.config.plugin.name.vodType')" prop="vodType">
            <el-radio-group v-model="saveForm.data.vodType">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.vodType') as any" :key="index" :label="item.value">
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
                <label>
                    <el-alert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.plugin.tip.vodOfAliyunEndpoint')"></span>
                        </template>
                    </el-alert>
                </label>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.vodOfAliyunRoleArn')" prop="vodOfAliyunRoleArn">
                <el-input v-model="saveForm.data.vodOfAliyunRoleArn" :placeholder="t('platform.config.plugin.name.vodOfAliyunRoleArn')" :clearable="true" style="max-width: 500px" />
                <label>
                    <el-alert :title="t('platform.config.plugin.tip.vodOfAliyunRoleArn')" type="info" :show-icon="true" :closable="false" />
                </label>
            </el-form-item>
        </template>

        <el-form-item>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
