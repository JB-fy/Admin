<script setup lang="ts">
import md5 from 'js-md5'

const { t } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        ...saveCommon.data
    } as { [propName: string]: any },
    rules: {
        account: [
            { type: 'string', required: computed((): boolean => { return saveForm.data.phone ? false : true; }), min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^(?!\d*$)[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.account') }
        ],
        phone: [
            { type: 'string', required: computed((): boolean => { return saveForm.data.account ? false : true; }), min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^1[3-9]\d{9}$/, trigger: 'blur', message: t('validation.phone') }
        ],
        password: [
            { type: 'string', required: computed((): boolean => { return saveForm.data.id ? false : true; }), min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ],
        roleIdArr: [
            { type: 'array', required: true, min: 1, defaultField: { type: 'integer' }, trigger: 'change', message: t('validation.select') }
        ],
        nickname: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        avatar: [
            { type: 'string', min: 1, max: 120, trigger: 'change', message: t('validation.upload') }
        ],
        isStop: [
            { type: 'enum', enum: [0, 1], trigger: 'change', message: t('validation.select') }
        ]
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data, false)
            param.password ? param.password = md5(param.password) : delete param.password
            try {
                if (param?.idArr?.length > 0) {
                    await request('/platform/admin/update', param, true)
                } else {
                    await request('/platform/admin/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } catch (error) { }
            saveForm.loading = false
        })
    }
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
            }).then(() => {
                done()
            }).catch(() => { })
        } else {
            done()
        }
    },
    buttonClose: () => {
        //saveCommon.visible = false
        saveDrawer.ref.handleClose()    //会触发beforeClose
    }
})
</script>

<template>
    <ElDrawer class="save-drawer" :ref="(el: any) => { saveDrawer.ref = el }" v-model="saveCommon.visible"
        :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <ElScrollbar>
            <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules"
                label-width="auto" :status-icon="true" :scroll-to-error="false">
                <ElFormItem :label="t('common.name.account')" prop="account">
                    <ElInput v-model="saveForm.data.account" :placeholder="t('common.name.account')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.phone')" prop="phone">
                    <ElInput v-model="saveForm.data.phone" :placeholder="t('common.name.phone')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.password')" prop="password">
                    <ElInput v-model="saveForm.data.password" :placeholder="t('common.name.password')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" :show-password="true"
                        style="max-width: 250px;" />
                    <label v-if="saveForm.data.id">
                        <ElAlert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem :label="t('common.name.nickname')" prop="nickname">
                    <ElInput v-model="saveForm.data.nickname" :placeholder="t('common.name.nickname')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.avatar')" prop="avatar">
                    <MyUpload v-model="saveForm.data.avatar" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.rel.roleIdArr')" prop="roleIdArr">
                    <MyTransfer v-model="saveForm.data.roleIdArr"
                        :api="{ code: 'auth/role/list', param: { field: ['id', 'roleName'] } }" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.isStop')" prop="isStop">
                    <ElSwitch v-model="saveForm.data.isStop" :active-value="1" :inactive-value="0" :inline-prompt="true"
                        :active-text="t('common.yes')" :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)" />
                </ElFormItem>
            </ElForm>
        </ElScrollbar>
        <template #footer>
            <ElButton @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</ElButton>
            <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                {{ t('common.save') }}
            </ElButton>
        </template>
    </ElDrawer>
</template>