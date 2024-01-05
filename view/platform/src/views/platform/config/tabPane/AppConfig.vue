<script setup lang="tsx">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        packageUrlOfAndroid: '',
        packageSizeOfAndroid: 0,
        packageNameOfAndroid: '',
        isForceUpdateOfAndroid: 0,
        versionNumberOfAndroid: 0,
        versionNameOfAndroid: '',
        versionIntroOfAndroid: '',

        packageUrlOfIos: '',
        packageSizeOfIos: 0,
        packageNameOfIos: '',
        isForceUpdateOfIos: 0,
        versionNumberOfIos: 0,
        versionNameOfIos: '',
        versionIntroOfIos: '',
        plistUrlOfIos: '',
    } as { [propName: string]: any },
    rules: {
        packageUrlOfAndroid: [{ type: 'url', trigger: 'change', message: t('validation.upload') }],
        packageSizeOfAndroid: [{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) }],
        packageNameOfAndroid: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        isForceUpdateOfAndroid: [{ type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') }],
        versionNumberOfAndroid: [{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) }],
        versionNameOfAndroid: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        versionIntroOfAndroid: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],

        packageUrlOfIos: [{ type: 'url', trigger: 'change', message: t('validation.upload') }],
        packageSizeOfIos: [{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) }],
        packageNameOfIos: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        isForceUpdateOfIos: [{ type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') }],
        versionNumberOfIos: [{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) }],
        versionNameOfIos: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        versionIntroOfIos: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        plistUrlOfIos: [{ type: 'url', trigger: 'change', message: t('validation.upload') }],
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

const handleOfAndroid = reactive({
    disabledOfPackageSize: false,
    afterUpload: () => {
        if (saveForm.data.packageUrlOfAndroid) {
            let packageSize = parseInt(new URL(saveForm.data.packageUrlOfAndroid).searchParams.get('s') as string)
            if (packageSize > 0) {
                saveForm.data.packageSizeOfAndroid = packageSize
                handleOfAndroid.disabledOfPackageSize = true
                return
            }
        }
        saveForm.data.packageSizeOfAndroid = 0
        handleOfAndroid.disabledOfPackageSize = false
    },
})

const handleOfIos = reactive({
    disabledOfPackageSize: false,
    afterUpload: () => {
        if (saveForm.data.packageUrlOfIos) {
            let packageSize = parseInt(new URL(saveForm.data.packageUrlOfIos).searchParams.get('s') as string)
            if (packageSize > 0) {
                saveForm.data.packageSizeOfIos = packageSize
                handleOfIos.disabledOfPackageSize = true
                return
            }
        }
        saveForm.data.packageSizeOfIos = 0
        handleOfIos.disabledOfPackageSize = false
    },
})

saveForm.initData()
</script>

<template>
    <ElForm :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
        <ElTabs tab-position="left">
            <ElTabPane :label="t('platform.config.label.android')" :lazy="true">
                <ElFormItem :label="t('platform.config.name.packageUrlOfAndroid')" prop="packageUrlOfAndroid">
                    <MyUpload v-model="saveForm.data.packageUrlOfAndroid" accept=".apk" :isImage="false" @change="handleOfAndroid.afterUpload" :key="saveForm.data.packageUrlOfAndroid" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.packageSizeOfAndroid')" prop="packageSizeOfAndroid">
                    <ElInputNumber v-model="saveForm.data.packageSizeOfAndroid" :precision="0" :min="0" :step="1" :step-strictly="true" :controls="false" :disabled="handleOfAndroid.disabledOfPackageSize" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.packageNameOfAndroid')" prop="packageNameOfAndroid">
                    <ElInput v-model="saveForm.data.packageNameOfAndroid" :placeholder="t('platform.config.name.packageNameOfAndroid')" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.isForceUpdateOfAndroid')" prop="isForceUpdateOfAndroid">
                    <ElSwitch
                        v-model="saveForm.data.isForceUpdateOfAndroid"
                        :active-value="1"
                        :inactive-value="0"
                        :inline-prompt="true"
                        :active-text="t('common.yes')"
                        :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.versionNumberOfAndroid')" prop="versionNumberOfAndroid">
                    <ElInputNumber v-model="saveForm.data.versionNumberOfAndroid" :precision="0" :min="0" :step="1" :step-strictly="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.versionNameOfAndroid')" prop="versionNameOfAndroid">
                    <ElInput v-model="saveForm.data.versionNameOfAndroid" :placeholder="t('platform.config.name.versionNameOfAndroid')" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.versionIntroOfAndroid')" prop="versionIntroOfAndroid">
                    <ElInput v-model="saveForm.data.versionIntroOfAndroid" type="textarea" :autosize="{ minRows: 3 }" />
                </ElFormItem>
            </ElTabPane>

            <ElTabPane :label="t('platform.config.label.ios')" :lazy="true">
                <ElFormItem :label="t('platform.config.name.packageUrlOfIos')" prop="packageUrlOfIos">
                    <MyUpload v-model="saveForm.data.packageUrlOfIos" accept=".ipa" :isImage="false" @change="handleOfIos.afterUpload" :key="saveForm.data.packageUrlOfIos" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.packageSizeOfIos')" prop="packageSizeOfIos">
                    <ElInputNumber v-model="saveForm.data.packageSizeOfIos" :precision="0" :min="0" :step="1" :step-strictly="true" :controls="false" :disabled="handleOfIos.disabledOfPackageSize" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.plistUrlOfIos')" prop="plistUrlOfIos">
                    <MyUpload v-model="saveForm.data.plistUrlOfIos" accept=".plist" :isImage="false" :key="saveForm.data.plistUrlOfIos" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.packageNameOfIos')" prop="packageNameOfIos">
                    <ElInput v-model="saveForm.data.packageNameOfIos" :placeholder="t('platform.config.name.packageNameOfIos')" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.isForceUpdateOfIos')" prop="isForceUpdateOfIos">
                    <ElSwitch
                        v-model="saveForm.data.isForceUpdateOfIos"
                        :active-value="1"
                        :inactive-value="0"
                        :inline-prompt="true"
                        :active-text="t('common.yes')"
                        :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.versionNumberOfIos')" prop="versionNumberOfIos">
                    <ElInputNumber v-model="saveForm.data.versionNumberOfIos" :precision="0" :min="0" :step="1" :step-strictly="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.versionNameOfIos')" prop="versionNameOfIos">
                    <ElInput v-model="saveForm.data.versionNameOfIos" :placeholder="t('platform.config.name.versionNameOfIos')" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.versionIntroOfIos')" prop="versionIntroOfIos">
                    <ElInput v-model="saveForm.data.versionIntroOfIos" type="textarea" :autosize="{ minRows: 3 }" />
                </ElFormItem>
            </ElTabPane>
        </ElTabs>

        <ElFormItem>
            <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <AutoiconEpCircleCheck />{{ t('common.save') }} </ElButton>
            <ElButton type="info" @click="saveForm.reset"> <AutoiconEpCircleClose />{{ t('common.reset') }} </ElButton>
        </ElFormItem>
    </ElForm>
</template>
