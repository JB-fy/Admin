<script setup lang="ts">
const { t } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        sort: 50,
        ...saveCommon.data
    } as { [propName: string]: any },
    rules: {
        roleName: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        sceneId: [
            { type: 'integer', required: true, min: 1, trigger: 'change', message: t('validation.select') }
        ],
        menuIdArr: [
            { type: 'array', trigger: 'change', message: t('validation.select') }
        ],
        actionIdArr: [
            { type: 'array', defaultField: { type: 'integer' }, trigger: 'change', message: t('validation.select') }
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
            let menuIdArr: any = []
            param.menuIdArr.forEach((item: any) => {
                menuIdArr = menuIdArr.concat(item)
            })
            //param.menuIdArr = [...new Set(menuIdArr)]
            param.menuIdArr = menuIdArr.filter((item: any, index: any) => {
                return menuIdArr.indexOf(item) === index
            })
            try {
                if (param?.idArr?.length > 0) {
                    await request('/auth/role/update', param, true)
                } else {
                    await request('/auth/role/create', param, true)
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
                label-width="auto" :status-icon="true" :scroll-to-error="true">
                <ElFormItem :label="t('common.name.auth.role.roleName')" prop="roleName">
                    <ElInput v-model="saveForm.data.roleName" :placeholder="t('common.name.auth.role.roleName')"
                        minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.rel.sceneId')" prop="sceneId">
                    <MySelect v-model="saveForm.data.sceneId" :api="{ code: 'auth/scene/list' }"
                        @change="() => { saveForm.data.menuIdArr = []; saveForm.data.actionIdArr = [] }" />
                </ElFormItem>
                <ElFormItem v-if="saveForm.data.sceneId" :label="t('common.name.rel.menuIdArr')" prop="menuIdArr">
                    <MyCascader v-model="saveForm.data.menuIdArr"
                        :api="{ code: 'auth/menu/tree', param: { filter: { sceneId: saveForm.data.sceneId } } }"
                        :isPanel="true" :props="{ multiple: true, checkStrictly: false, emitPath: true }" />
                </ElFormItem>
                <ElFormItem v-if="saveForm.data.sceneId" :label="t('common.name.rel.actionIdArr')" prop="actionIdArr">
                    <MyTransfer v-model="saveForm.data.actionIdArr"
                        :api="{ code: 'auth/action/list', param: { filter: { sceneId: saveForm.data.sceneId } } }" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.isStop')" prop="isStop">
                    <ElSwitch v-model="saveForm.data.isStop" :active-value="1" :inactive-value="0" :inline-prompt="true"
                        :active-text="t('common.yes')" :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
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