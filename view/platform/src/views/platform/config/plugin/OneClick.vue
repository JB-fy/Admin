<script setup lang="tsx">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        oneClickType: 'oneClickOfYidun',
        oneClickOfYidunSecretId: '',
        oneClickOfYidunSecretKey: '',
        oneClickOfYidunBusinessId: '',
    } as { [propName: string]: any },
    rules: {
        oneClickType: [{ type: 'enum', trigger: 'change', enum: [`oneClickOfYidun`], message: t('validation.select') }],
        oneClickOfYidunSecretId: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        oneClickOfYidunSecretKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        oneClickOfYidunBusinessId: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    initData: async () => {
        const param = { config_key_arr: Object.keys(saveForm.data) }
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
            const param = removeEmptyOfObj(saveForm.data)
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
        <el-form-item :label="t('platform.config.plugin.name.oneClickType')" prop="oneClickType">
            <el-radio-group v-model="saveForm.data.oneClickType">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.oneClickType') as any" :key="index" :value="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.oneClickType == 'oneClickOfYidun'">
            <el-form-item :label="t('platform.config.plugin.name.oneClickOfYidunSecretId')" prop="oneClickOfYidunSecretId">
                <el-input v-model="saveForm.data.oneClickOfYidunSecretId" :placeholder="t('platform.config.plugin.name.oneClickOfYidunSecretId')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.oneClickOfYidunSecretKey')" prop="oneClickOfYidunSecretKey">
                <el-input v-model="saveForm.data.oneClickOfYidunSecretKey" :placeholder="t('platform.config.plugin.name.oneClickOfYidunSecretKey')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.oneClickOfYidunBusinessId')" prop="oneClickOfYidunBusinessId">
                <el-input v-model="saveForm.data.oneClickOfYidunBusinessId" :placeholder="t('platform.config.plugin.name.oneClickOfYidunBusinessId')" :clearable="true" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
