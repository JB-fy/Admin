<script setup lang="tsx">
const { t } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        hot_search: [],
    } as { [propName: string]: any },
    rules: {
        hot_search: [
            { type: 'array', trigger: 'change', message: t('validation.array'), defaultField: { type: 'string', message: t('validation.input') } }, // 限制数组数量时用：max: 10, message: t('validation.max.array', { max: 10 })
        ],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    initData: async () => {
        const param = { config_key_arr: Object.keys(saveForm.data) }
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/org/config/get', param)
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
                await request(t('config.VITE_HTTP_API_PREFIX') + '/org/config/save', param, true)
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

const hotSearchHandle = reactive({
    ref: [] as any[],
    add: () => {
        !Array.isArray(saveForm.data.hot_search) && (saveForm.data.hot_search = [])
        saveForm.data.hot_search.push(undefined)
        nextTick(() => hotSearchHandle.ref[hotSearchHandle.ref.length - 1].focus())
    },
    del: (index: number, isBlur: boolean = false) => {
        if (isBlur && saveForm.data.hot_search[index] !== undefined && saveForm.data.hot_search[index] !== null && saveForm.data.hot_search[index] !== '') {
            return
        }
        saveForm.data.hot_search.splice(index, 1)
        hotSearchHandle.ref.splice(index, 1)
    },
})

saveForm.initData()
</script>

<template>
    <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
        <el-form-item :label="t('org.config.app.name.hot_search')" prop="hot_search">
            <template v-for="(_, index) in saveForm.data.hot_search" :key="index">
                <el-tag type="info" :closable="true" @close="hotSearchHandle.del(index)" size="large" style="padding-left: 0; margin: 3px 10px 3px 0">
                    <el-input :ref="(el: any) => hotSearchHandle.ref[index] = el" v-model="saveForm.data.hot_search[index]" @blur="hotSearchHandle.del(index, true)" :placeholder="t('org.config.app.name.hot_search')" style="width: 150px" />
                </el-tag>
            </template>
            <el-button type="primary" @click="hotSearchHandle.add" style="margin: 3px 0"> <autoicon-ep-plus />{{ t('common.add') }} </el-button>
        </el-form-item>
        <el-form-item>
            <el-button v-if="authAction.isCommonSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
