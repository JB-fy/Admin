<script setup lang="ts">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
	ref: null as any,
	loading: false,
	data: {
		isStop: 0,
		...saveCommon.data
	} as { [propName: string]: any },
	rules: {
		roleName: [
			{ type: 'string', required: true, max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		sceneId: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		],
		/* tableId: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		], */
		isStop: [
			{ type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
		],
		menuIdArr: [
			{ type: 'array', trigger: 'change', message: t('validation.select') },
		],
		actionIdArr: [
			{ type: 'array', defaultField: { type: 'integer' }, trigger: 'change', message: t('validation.select') },
		],
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
					await request(t('config.VITE_HTTP_API_PREFIX') + '/auth/role/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/auth/role/create', param, true)
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
		saveDrawer.ref.handleClose()	//会触发beforeClose
	}
})
</script>

<template>
	<ElDrawer class="save-drawer" :ref="(el: any) => { saveDrawer.ref = el }" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
		<ElScrollbar>
			<ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
				<ElFormItem :label="t('auth.role.name.roleName')" prop="roleName">
					<ElInput v-model="saveForm.data.roleName" :placeholder="t('auth.role.name.roleName')" maxlength="30" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('auth.role.name.sceneId')" prop="sceneId">
					<MySelect v-model="saveForm.data.sceneId" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list' }" @change="() => { saveForm.data.menuIdArr = []; saveForm.data.actionIdArr = [] }" />
				</ElFormItem>
				<!-- <ElFormItem :label="t('auth.role.name.tableId')" prop="tableId">
					<MySelect v-model="saveForm.data.tableId" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/table/list' }" />
				</ElFormItem> -->
				<ElFormItem v-if="saveForm.data.sceneId" :label="t('auth.role.name.menuId')" prop="menuIdArr">
					<MyCascader v-model="saveForm.data.menuIdArr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { filter: { sceneId: saveForm.data.sceneId } } }" :isPanel="true" :props="{ multiple: true }" />
				</ElFormItem>
				<ElFormItem v-if="saveForm.data.sceneId" :label="t('auth.role.name.actionId')" prop="actionIdArr">
					<MyTransfer v-model="saveForm.data.actionIdArr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/action/list', param: { filter: { sceneId: saveForm.data.sceneId } } }" />
				</ElFormItem>
				<ElFormItem :label="t('auth.role.name.isStop')" prop="isStop">
					<ElSwitch v-model="saveForm.data.isStop" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
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