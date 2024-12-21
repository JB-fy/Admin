<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        pay_method: 0,
        sort: 100,
        ...saveCommon.data,
        scene_id: saveCommon.data.scene_id ? saveCommon.data.scene_id : undefined,
        pay_id: saveCommon.data.pay_id ? saveCommon.data.pay_id : undefined,
    } as { [propName: string]: any },
    rules: {
        channel_name: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
        ],
        channel_icon: [
            { type: 'string', trigger: 'blur', max: 200, message: t('validation.max.string', { max: 200 }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],
        scene_id: [
            { required: true, message: t('validation.required') },
            { type: 'integer', trigger: 'change', min: 1, max: 4294967295, message: t('validation.select') },
        ],
        pay_id: [
            { required: true, message: t('validation.required') },
            { type: 'integer', trigger: 'change', min: 1, max: 4294967295, message: t('validation.select') },
        ],
        pay_method: [{ type: 'enum', trigger: 'change', enum: (tm('pay.channel.status.pay_method') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
        sort: [{ type: 'integer', trigger: 'change', min: 0, max: 255, message: t('validation.between.number', { min: 0, max: 255 }) }],
        /* total_amount: [
            { type: 'number', trigger: 'change', min: 0, max: 999999999999.99, message: t('validation.between.number', { min: 0, max: 999999999999.99 }) }, // type: 'float'在值为0时验证不能通过
        ], */
        is_stop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            param.scene_id === undefined && (param.scene_id = 0)
            param.pay_id === undefined && (param.pay_id = 0)
            try {
                if (param?.id) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/pay/channel/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/pay/channel/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } finally {
                saveForm.loading = false
            }
        })
    },
})

const saveDrawer = reactive({
    ref: null as any,
    size: useSettingStore().saveDrawer.size,
    beforeClose: (done: Function) => {
        if (useSettingStore().saveDrawer.isTipClose) {
            ElMessageBox.confirm('', {
                type: 'info',
                title: t('common.tip.configExit'),
                center: true,
                showClose: false,
            }).then(() => done())
        } else {
            done()
        }
    },
    buttonClose: () => saveDrawer.ref.handleClose(), //saveCommon.visible = false //不会触发beforeClose
})
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
                <el-form-item :label="t('pay.channel.name.channel_name')" prop="channel_name">
                    <el-input v-model="saveForm.data.channel_name" :placeholder="t('pay.channel.name.channel_name')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('pay.channel.name.channel_icon')" prop="channel_icon">
                    <my-upload v-model="saveForm.data.channel_icon" accept="image/*" />
                </el-form-item>
                <el-form-item :label="t('pay.channel.name.scene_id')" prop="scene_id">
                    <my-select v-model="saveForm.data.scene_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/pay/scene/list' }" />
                </el-form-item>
                <el-form-item :label="t('pay.channel.name.pay_id')" prop="pay_id">
                    <my-select v-model="saveForm.data.pay_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/pay/pay/list' }" />
                </el-form-item>
                <el-form-item :label="t('pay.channel.name.pay_method')" prop="pay_method">
                    <el-radio-group v-model="saveForm.data.pay_method">
                        <el-radio v-for="(item, index) in (tm('pay.channel.status.pay_method') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item :label="t('pay.channel.name.sort')" prop="sort">
                    <el-input-number v-model="saveForm.data.sort" :placeholder="t('pay.channel.name.sort')" :min="0" :max="255" :precision="0" :value-on-clear="100" />
                    <el-alert :title="t('pay.channel.tip.sort')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <!-- <el-form-item :label="t('pay.channel.name.total_amount')" prop="total_amount">
                    <el-input-number v-model="saveForm.data.total_amount" :placeholder="t('pay.channel.name.total_amount')" :min="0" :max="999999999999.99" :precision="2" :controls="false" :value-on-clear="0.0" />
                </el-form-item> -->
                <el-form-item :label="t('pay.channel.name.is_stop')" prop="is_stop">
                    <el-switch
                        v-model="saveForm.data.is_stop"
                        :active-value="(tm('common.status.whether') as any[])[1].value"
                        :inactive-value="(tm('common.status.whether') as any[])[0].value"
                        :active-text="(tm('common.status.whether') as any[])[1].label"
                        :inactive-text="(tm('common.status.whether') as any[])[0].label"
                        :inline-prompt="true"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </el-form-item>
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">{{ t('common.save') }}</el-button>
        </template>
    </el-drawer>
</template>
