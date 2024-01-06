<script setup lang="tsx">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        idCardType: 'idCardOfAliyun',
        idCardOfAliyunHost: '',
        idCardOfAliyunPath: '',
        idCardOfAliyunAppcode: '',
    } as { [propName: string]: any },
    rules: {
        idCardType: [{ type: 'enum', enum: [`idCardOfAliyun`], trigger: 'change', message: t('validation.select') }],
        idCardOfAliyunHost: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        idCardOfAliyunPath: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        idCardOfAliyunAppcode: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
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
        <el-form-item :label="t('platform.config.name.idCardType')" prop="idCardType">
            <el-radio-group v-model="saveForm.data.idCardType">
                <el-radio v-for="(item, index) in tm('platform.config.status.idCardType') as any" :key="index" :label="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.idCardType == 'idCardOfAliyun'">
            <el-form-item :label="t('platform.config.name.idCardOfAliyunHost')" prop="idCardOfAliyunHost">
                <el-input v-model="saveForm.data.idCardOfAliyunHost" :placeholder="t('platform.config.name.idCardOfAliyunHost')" :clearable="true" style="max-width: 500px" />
                <label>
                    <el-alert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.tip.idCardOfAliyunHost')"></span>
                        </template>
                    </el-alert>
                </label>
            </el-form-item>
            <el-form-item :label="t('platform.config.name.idCardOfAliyunPath')" prop="idCardOfAliyunPath">
                <el-input v-model="saveForm.data.idCardOfAliyunPath" :placeholder="t('platform.config.name.idCardOfAliyunPath')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.name.idCardOfAliyunAppcode')" prop="idCardOfAliyunAppcode">
                <el-input v-model="saveForm.data.idCardOfAliyunAppcode" :placeholder="t('platform.config.name.idCardOfAliyunAppcode')" :clearable="true" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
