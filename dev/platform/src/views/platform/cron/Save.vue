<script setup lang="ts">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
	ref: null as any,
	loading: false,
	data: {
		...saveCommon.data
	} as { [propName: string]: any },
	rules: {
		cronName: [
			{ type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],
		cronCode: [
			{ type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],
		cronPattern: [
			{ type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
		],
		remark: [
			{ type: 'string', min: 1, max: 120, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 120 }) },
		],
		isStop: [
            { type: 'enum', enum: [0, 1], trigger: 'change', message: t('validation.select') }
        ],
	} as any,
	submit: () => {
		saveForm.ref.validate(async (valid: boolean) => {
			if (!valid) {
				return false
			}
			saveForm.loading = true
			const param = removeEmptyOfObj(saveForm.data, false)
			try {
				if (param?.idArr?.length > 0) {
					await request('/platform/cron/update', param, true)
				} else {
					await request('/platform/cron/create', param, true)
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
				<ElFormItem :label="t('platform.cron.name.cronName')" prop="cronName">
					<ElInput v-model="saveForm.data.cronName" :placeholder="t('platform.cron.name.cronName')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('platform.cron.name.cronCode')" prop="cronCode">
					<ElInput v-model="saveForm.data.cronCode" :placeholder="t('platform.cron.name.cronCode')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('platform.cron.name.cronPattern')" prop="cronPattern">
					<ElInput v-model="saveForm.data.cronPattern" :placeholder="t('platform.cron.name.cronPattern')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('common.name.remark')" prop="remark">
					<ElInput v-model="saveForm.data.remark" type="textarea" :autosize="{ minRows: 3 }" />
				</ElFormItem>
				<ElFormItem :label="t('common.name.isStop')" prop="isStop">
                    <ElSwitch v-model="saveForm.data.isStop" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')"
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