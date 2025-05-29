<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const payConfig = saveCommon.data.pay_config ? JSON.parse(saveCommon.data.pay_config) : {}
const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        pay_type: 0,
        ...saveCommon.data,
        pay_config_0: (() => {
            const payConfig0: { [propName: string]: any } = {}
            if (saveCommon.data.pay_type == 0) {
                payConfig0.app_id = payConfig.app_id
                delete payConfig.app_id
                payConfig0.private_key = payConfig.private_key
                delete payConfig.private_key
                payConfig0.public_key = payConfig.public_key
                delete payConfig.public_key
                payConfig0.op_app_id = payConfig.op_app_id
                delete payConfig.op_app_id
            }
            return payConfig0
        })(),
        pay_config_1: (() => {
            const payConfig1: { [propName: string]: any } = {}
            if (saveCommon.data.pay_type == 1) {
                payConfig1.app_id = payConfig.app_id
                delete payConfig.app_id
                payConfig1.mch_id = payConfig.mch_id
                delete payConfig.mch_id
                payConfig1.serial_no = payConfig.serial_no
                delete payConfig.serial_no
                payConfig1.api_v3_key = payConfig.api_v3_key
                delete payConfig.api_v3_key
                payConfig1.private_key = payConfig.private_key
                delete payConfig.private_key
            }
            return payConfig1
        })(),
        pay_config: Object.keys(payConfig).length > 0 ? JSON.stringify(payConfig) : undefined,
    } as { [propName: string]: any },
    rules: {
        pay_name: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
        ],
        pay_type: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: (tm('pay.pay.status.pay_type') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') },
        ],
        'pay_config_0.app_id': [
            { required: computed((): boolean => (saveForm.data.pay_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config_0.private_key': [
            { required: computed((): boolean => (saveForm.data.pay_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config_0.public_key': [
            { required: computed((): boolean => (saveForm.data.pay_type == 0 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config_0.op_app_id': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'pay_config_1.app_id': [
            { required: computed((): boolean => (saveForm.data.pay_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config_1.mch_id': [
            { required: computed((): boolean => (saveForm.data.pay_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config_1.serial_no': [
            { required: computed((): boolean => (saveForm.data.pay_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config_1.api_v3_key': [
            { required: computed((): boolean => (saveForm.data.pay_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'pay_config_1.private_key': [
            { required: computed((): boolean => (saveForm.data.pay_type == 1 ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        pay_config: [
            // { required: true, message: t('validation.required') },
            {
                type: 'object',
                trigger: 'blur',
                message: t('validation.json'),
                // fields: { xxxx: [{ required: true, message: 'xxxx' + t('validation.required') }] }, //内部添加规则时，不再需要设置trigger属性
                transform: (value: any) => (value ? jsonDecode(value) : undefined),
            },
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
        remark: [{ type: 'string', trigger: 'blur', max: 120, message: t('validation.max.string', { max: 120 }) }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            let payConfig = param.pay_config ? JSON.parse(param.pay_config) : {}
            payConfig = { ...payConfig, ...param['pay_config_' + param.pay_type] }
            param.pay_config = Object.keys(payConfig).length > 0 ? JSON.stringify(payConfig) : ''
            try {
                if (param?.id) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/pay/pay/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/pay/pay/create', param, true)
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
                <el-form-item :label="t('pay.pay.name.pay_name')" prop="pay_name">
                    <el-input v-model="saveForm.data.pay_name" :placeholder="t('pay.pay.name.pay_name')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('pay.pay.name.pay_type')" prop="pay_type">
                    <el-radio-group v-model="saveForm.data.pay_type">
                        <el-radio v-for="(item, index) in (tm('pay.pay.status.pay_type') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
                <template v-if="saveForm.data.pay_type == 0">
                    <el-form-item :label="t('pay.pay.name.pay_config_0.app_id')" prop="pay_config_0.app_id">
                        <el-input v-model="saveForm.data.pay_config_0.app_id" :placeholder="t('pay.pay.name.pay_config_0.app_id')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_0.private_key')" prop="pay_config_0.private_key">
                        <el-input v-model="saveForm.data.pay_config_0.private_key" type="textarea" :autosize="{ minRows: 5 }" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_0.public_key')" prop="pay_config_0.public_key">
                        <el-input v-model="saveForm.data.pay_config_0.public_key" type="textarea" :autosize="{ minRows: 5 }" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_0.op_app_id')" prop="pay_config_0.op_app_id">
                        <el-input v-model="saveForm.data.pay_config_0.op_app_id" :placeholder="t('pay.pay.name.pay_config_0.op_app_id')" :clearable="true" style="max-width: 250px" />
                        <el-alert :title="t('pay.pay.tip.pay_config_0.op_app_id')" type="info" :show-icon="true" :closable="false" />
                    </el-form-item>
                </template>
                <template v-else-if="saveForm.data.pay_type == 1">
                    <el-form-item :label="t('pay.pay.name.pay_config_1.app_id')" prop="pay_config_1.app_id">
                        <el-input v-model="saveForm.data.pay_config_1.app_id" :placeholder="t('pay.pay.name.pay_config_1.app_id')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_1.mch_id')" prop="pay_config_1.mch_id">
                        <el-input v-model="saveForm.data.pay_config_1.mch_id" :placeholder="t('pay.pay.name.pay_config_1.mch_id')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_1.serial_no')" prop="pay_config_1.serial_no">
                        <el-input v-model="saveForm.data.pay_config_1.serial_no" :placeholder="t('pay.pay.name.pay_config_1.serial_no')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_1.api_v3_key')" prop="pay_config_1.api_v3_key">
                        <el-input v-model="saveForm.data.pay_config_1.api_v3_key" :placeholder="t('pay.pay.name.pay_config_1.api_v3_key')" :clearable="true" />
                    </el-form-item>
                    <el-form-item :label="t('pay.pay.name.pay_config_1.private_key')" prop="pay_config_1.private_key">
                        <el-input v-model="saveForm.data.pay_config_1.private_key" type="textarea" :autosize="{ minRows: 5 }" />
                    </el-form-item>
                </template>
                <el-form-item :label="t('pay.pay.name.pay_config')" prop="pay_config">
                    <el-alert :title="t('pay.pay.tip.pay_config')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
                    <el-input v-model="saveForm.data.pay_config" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item>
                <el-form-item :label="t('pay.pay.name.pay_rate')" prop="pay_rate">
                    <el-input-number v-model="saveForm.data.pay_rate" :placeholder="t('pay.pay.name.pay_rate')" :min="0" :max="0.9999" :precision="4" :controls="false" :value-on-clear="0.0" />
                </el-form-item>
                <!-- <el-form-item :label="t('pay.pay.name.total_amount')" prop="total_amount">
                    <el-input-number v-model="saveForm.data.total_amount" :placeholder="t('pay.pay.name.total_amount')" :min="0" :max="999999999999.99" :precision="2" :controls="false" :value-on-clear="0.0" />
                </el-form-item>
                <el-form-item :label="t('pay.pay.name.balance')" prop="balance">
                    <el-input-number v-model="saveForm.data.balance" :placeholder="t('pay.pay.name.balance')" :min="0" :max="999999999999.999999" :precision="6" :controls="false" :value-on-clear="0.0" />
                </el-form-item> -->
                <el-form-item :label="t('pay.pay.name.remark')" prop="remark">
                    <el-input v-model="saveForm.data.remark" type="textarea" :autosize="{ minRows: 3 }" maxlength="120" :show-word-limit="true" />
                </el-form-item>
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">{{ t('common.save') }}</el-button>
        </template>
    </el-drawer>
</template>
