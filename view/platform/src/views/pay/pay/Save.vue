<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        pay_type: 0,
        sort: 100,
        ...saveCommon.data,
        pay_config: saveCommon.data.pay_config ? JSON.parse(saveCommon.data.pay_config) : {},
    } as { [propName: string]: any },
    rules: {
        pay_name: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
        ],
        pay_icon: [
            { type: 'string', trigger: 'blur', max: 200, message: t('validation.max.string', { max: 200 }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],
        pay_type: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: (tm('pay.pay.status.pay_type') as any).map((item: any) => item.value), message: t('validation.select') },
        ],
        /* pay_config: [
            { required: true, message: t('validation.required') },
            {
                type: 'object',
                trigger: 'blur',
                message: t('validation.json'),
                transform: (value: any) => {
                    if (value === '' || value === null || value === undefined) {
                        return undefined
                    }
                    try {
                        return JSON.parse(value)
                    } catch (e) {
                        return value
                    }
                },
            },
        ], */
        'pay_config.payOfAliAppId': [
            { required: computed((): boolean => (saveForm.data.pay_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config.payOfAliPrivateKey': [
            { required: computed((): boolean => (saveForm.data.pay_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config.payOfAliPublicKey': [
            { required: computed((): boolean => (saveForm.data.pay_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config.payOfAliOpAppId': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'pay_config.payOfWxAppId': [
            { required: computed((): boolean => (saveForm.data.pay_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config.payOfWxMchid': [
            { required: computed((): boolean => (saveForm.data.pay_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config.payOfWxSerialNo': [
            { required: computed((): boolean => (saveForm.data.pay_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config.payOfWxApiV3Key': [
            { required: computed((): boolean => (saveForm.data.pay_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config.payOfWxPrivateKey': [
            { required: computed((): boolean => (saveForm.data.pay_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        pay_rate: [
            { type: 'number', trigger: 'change', min: 0, max: 0.9999, message: t('validation.between.number', { min: 0, max: 0.9999 }) }, // type: 'float'在值为0时验证不能通过
        ],
        /* total_amount: [
            { type: 'number', trigger: 'change', min: 0, max: 999999999999.99, message: t('validation.between.number', { min: 0, max: 999999999999.99 }) }, // type: 'float'在值为0时验证不能通过
        ],
        balance: [
            { type: 'number', trigger: 'change', min: 0, max: 999999999999.999999, message: t('validation.between.number', { min: 0, max: 999999999999.999999 }) }, // type: 'float'在值为0时验证不能通过
        ], */
        sort: [{ type: 'integer', trigger: 'change', min: 0, max: 255, message: t('validation.between.number', { min: 0, max: 255 }) }],
        remark: [{ type: 'string', trigger: 'blur', max: 120, message: t('validation.max.string', { max: 120 }) }],
        pay_scene_arr: [
            { required: true, message: t('validation.required') },
            { type: 'array', trigger: 'change', message: t('validation.select'), defaultField: { type: 'enum', enum: (tm('pay.pay.status.pay_scene_arr') as any).map((item: any) => item.value), message: t('validation.select') } }, // 限制数组数量时用：max: 10, message: t('validation.max.select', { max: 10 })
        ],
        is_stop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            try {
                if (param?.id_arr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/pay/pay/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/pay/pay/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } catch (error) {}
            saveForm.loading = false
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
            })
                .then(() => {
                    done()
                })
                .catch(() => {})
        } else {
            done()
        }
    },
    buttonClose: () => {
        //saveCommon.visible = false
        saveDrawer.ref.handleClose() //会触发beforeClose
    },
})
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
                <el-form-item :label="t('pay.pay.name.pay_name')" prop="pay_name">
                    <el-input v-model="saveForm.data.pay_name" :placeholder="t('pay.pay.name.pay_name')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('pay.pay.name.pay_icon')" prop="pay_icon">
                    <my-upload v-model="saveForm.data.pay_icon" accept="image/*" />
                </el-form-item>
                <el-form-item :label="t('pay.pay.name.pay_type')" prop="pay_type">
                    <el-radio-group v-model="saveForm.data.pay_type" @change="() => (saveForm.data.pay_config = {})">
                        <el-radio v-for="(item, index) in (tm('pay.pay.status.pay_type') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
                <!-- <el-form-item :label="t('pay.pay.name.pay_config')" prop="pay_config">
                    <el-alert :title="t('pay.pay.tip.pay_config')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
                    <el-input v-model="saveForm.data.pay_config" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item> -->
                <template v-if="saveForm.data.pay_type == 0">
                    <el-form-item :label="t('pay.pay.name.pay_config_obj.payOfAliAppId')" prop="pay_config.payOfAliAppId">
                        <el-input v-model="saveForm.data.pay_config.payOfAliAppId" :placeholder="t('pay.pay.name.pay_config_obj.payOfAliAppId')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_obj.payOfAliPrivateKey')" prop="pay_config.payOfAliPrivateKey">
                        <el-input v-model="saveForm.data.pay_config.payOfAliPrivateKey" type="textarea" :autosize="{ minRows: 5 }" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_obj.payOfAliPublicKey')" prop="pay_config.payOfAliPublicKey">
                        <el-input v-model="saveForm.data.pay_config.payOfAliPublicKey" type="textarea" :autosize="{ minRows: 5 }" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_obj.payOfAliOpAppId')" prop="pay_config.payOfAliOpAppId">
                        <el-input v-model="saveForm.data.pay_config.payOfAliOpAppId" :placeholder="t('pay.pay.name.pay_config_obj.payOfAliOpAppId')" :clearable="true" style="max-width: 250px" />
                        <el-alert :title="t('pay.pay.tip.pay_config_obj.payOfAliOpAppId')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                </template>
                <template v-else-if="saveForm.data.pay_type == 1">
                    <el-form-item :label="t('pay.pay.name.pay_config_obj.payOfWxAppId')" prop="pay_config.payOfWxAppId">
                        <el-input v-model="saveForm.data.pay_config.payOfWxAppId" :placeholder="t('pay.pay.name.pay_config_obj.payOfWxAppId')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_obj.payOfWxMchid')" prop="pay_config.payOfWxMchid">
                        <el-input v-model="saveForm.data.pay_config.payOfWxMchid" :placeholder="t('pay.pay.name.pay_config_obj.payOfWxMchid')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_obj.payOfWxSerialNo')" prop="pay_config.payOfWxSerialNo">
                        <el-input v-model="saveForm.data.pay_config.payOfWxSerialNo" :placeholder="t('pay.pay.name.pay_config_obj.payOfWxSerialNo')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_obj.payOfWxApiV3Key')" prop="pay_config.payOfWxApiV3Key">
                        <el-input v-model="saveForm.data.pay_config.payOfWxApiV3Key" :placeholder="t('pay.pay.name.pay_config_obj.payOfWxApiV3Key')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_obj.payOfWxPrivateKey')" prop="pay_config.payOfWxPrivateKey">
                        <el-input v-model="saveForm.data.pay_config.payOfWxPrivateKey" type="textarea" :autosize="{ minRows: 5 }" />
                    </el-form-item>
                </template>
                <el-form-item :label="t('pay.pay.name.pay_rate')" prop="pay_rate">
                    <el-input-number v-model="saveForm.data.pay_rate" :placeholder="t('pay.pay.name.pay_rate')" :min="0" :max="0.9999" :precision="4" :controls="false" :value-on-clear="0.0" />
                </el-form-item>
                <!-- <el-form-item :label="t('pay.pay.name.total_amount')" prop="total_amount">
                    <el-input-number v-model="saveForm.data.total_amount" :placeholder="t('pay.pay.name.total_amount')" :min="0" :max="999999999999.99" :precision="2" :controls="false" :value-on-clear="0.0" />
                </el-form-item>
                <el-form-item :label="t('pay.pay.name.balance')" prop="balance">
                    <el-input-number v-model="saveForm.data.balance" :placeholder="t('pay.pay.name.balance')" :min="0" :max="999999999999.999999" :precision="6" :controls="false" :value-on-clear="0.0" />
                </el-form-item> -->
                <el-form-item :label="t('pay.pay.name.sort')" prop="sort">
                    <el-input-number v-model="saveForm.data.sort" :precision="0" :min="0" :max="255" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="100" />
                    <el-alert :title="t('pay.pay.tip.sort')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('pay.pay.name.remark')" prop="remark">
                    <el-input v-model="saveForm.data.remark" type="textarea" :autosize="{ minRows: 3 }" maxlength="120" :show-word-limit="true" />
                </el-form-item>
                <el-form-item :label="t('pay.pay.name.pay_scene_arr')" prop="pay_scene_arr">
                    <!-- 根据个人喜好选择组件<el-transfer>或<el-select-v2> -->
                    <el-transfer v-model="saveForm.data.pay_scene_arr" :data="tm('pay.pay.status.pay_scene_arr')" :props="{ key: 'value', label: 'label' }" />
                    <!-- <el-select-v2 v-model="saveForm.data.pay_scene_arr" :options="tm('pay.pay.status.pay_scene_arr')" :placeholder="t('pay.pay.name.pay_scene_arr')" :multiple="true" :collapse-tags="true" :collapse-tags-tooltip="true" style="width: 212px" /> -->
                </el-form-item>
                <el-form-item :label="t('pay.pay.name.is_stop')" prop="is_stop">
                    <el-switch
                        v-model="saveForm.data.is_stop"
                        :active-value="1"
                        :inactive-value="0"
                        :inline-prompt="true"
                        :active-text="t('common.yes')"
                        :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </el-form-item>
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                {{ t('common.save') }}
            </el-button>
        </template>
    </el-drawer>
</template>
