<script setup lang="tsx">
import md5 from 'js-md5'
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        gender: 0,
        ...saveCommon.data,
    } as { [propName: string]: any },
    rules: {
        phone: [
            { type: 'string', max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) },
            { pattern: /^1[3-9]\d{9}$/, trigger: 'blur', message: t('validation.phone') },
        ],
        account: [
            { type: 'string', max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) },
            { pattern: /^[\p{L}][\p{L}\p{N}_]{3,}$/u, trigger: 'blur', message: t('validation.account') },
        ],
        password: [
            {
                type: 'string',
                required: computed((): boolean => {
                    return saveForm.data.idArr?.length ? false : true
                }),
                min: 6,
                max: 20,
                trigger: 'blur',
                message: t('validation.between.string', { min: 6, max: 20 }),
            },
        ],
        nickname: [{ type: 'string', max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) }],
        avatar: [
            { type: 'string', max: 200, trigger: 'blur', message: t('validation.max.string', { max: 200 }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],
        gender: [{ type: 'enum', enum: (tm('user.user.status.gender') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') }],
        birthday: [{ type: 'string', trigger: 'change', message: t('validation.select') }],
        address: [{ type: 'string', max: 60, trigger: 'blur', message: t('validation.max.string', { max: 60 }) }],
        idCardName: [{ type: 'string', max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) }],
        idCardNo: [{ type: 'string', max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) }],
        isStop: [{ type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') }],
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            param.password ? (param.password = md5(param.password)) : delete param.password
            try {
                if (param?.idArr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/user/user/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/user/user/create', param, true)
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
                <el-form-item :label="t('user.user.name.phone')" prop="phone">
                    <el-input v-model="saveForm.data.phone" :placeholder="t('user.user.name.phone')" maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('user.user.name.account')" prop="account">
                    <el-input v-model="saveForm.data.account" :placeholder="t('user.user.name.account')" maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('user.user.name.password')" prop="password">
                    <el-input v-model="saveForm.data.password" :placeholder="t('user.user.name.password')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert v-if="saveForm.data.idArr?.length" :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('user.user.name.nickname')" prop="nickname">
                    <el-input v-model="saveForm.data.nickname" :placeholder="t('user.user.name.nickname')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('user.user.name.avatar')" prop="avatar">
                    <my-upload v-model="saveForm.data.avatar" accept="image/*" />
                </el-form-item>
                <el-form-item :label="t('user.user.name.gender')" prop="gender">
                    <el-radio-group v-model="saveForm.data.gender">
                        <el-radio v-for="(item, index) in (tm('user.user.status.gender') as any)" :key="index" :label="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item :label="t('user.user.name.birthday')" prop="birthday">
                    <el-date-picker v-model="saveForm.data.birthday" type="date" :placeholder="t('user.user.name.birthday')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 150px" />
                </el-form-item>
                <el-form-item :label="t('user.user.name.address')" prop="address">
                    <el-input v-model="saveForm.data.address" :placeholder="t('user.user.name.address')" maxlength="60" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('user.user.name.idCardName')" prop="idCardName">
                    <el-input v-model="saveForm.data.idCardName" :placeholder="t('user.user.name.idCardName')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('user.user.name.idCardNo')" prop="idCardNo">
                    <el-input v-model="saveForm.data.idCardNo" :placeholder="t('user.user.name.idCardNo')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('user.user.name.isStop')" prop="isStop">
                    <el-switch
                        v-model="saveForm.data.isStop"
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
