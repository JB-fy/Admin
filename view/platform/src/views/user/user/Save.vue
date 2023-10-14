<script setup lang="ts">
import md5 from 'js-md5'

const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
	ref: null as any,
	loading: false,
	data: {
		gender: 0,
		...saveCommon.data
	} as { [propName: string]: any },
	rules: {
		phone: [
			{ type: 'string', required: computed((): boolean => { return saveForm.data.account ? false : true; }), min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^1[3-9]\d{9}$/, trigger: 'blur', message: t('validation.phone') },
		],
		account: [
			{ type: 'string', required: computed((): boolean => { return saveForm.data.phone ? false : true; }), min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^(?!\d*$)[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.account') },
		],
		password: [
			{ type: 'string', required: computed((): boolean => { return saveForm.data.idArr?.length ? false : true; }), min: 1, max: 32, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 32 }) }
		],
		nickname: [
			{ type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
		],
		avatar: [
			{ type: 'url', trigger: 'change', message: t('validation.upload') },
			{ type: 'string', min: 1, max: 120, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 120 }) }
		],
		gender: [
			{ type: 'enum', enum: [0, 1, 2], trigger: 'change', message: t('validation.select') }
		],
		birthday: [
			{ type: 'string', trigger: 'change', message: t('validation.select') }
		],
		address: [
			{ type: 'string', min: 1, max: 60, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 60 }) },
		],
		idCardName: [
			{ type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		idCardNo: [
			{ type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
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
			param.password ? param.password = md5(param.password) : delete param.password
			try {
				if (param?.idArr?.length > 0) {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/user/user/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/user/user/create', param, true)
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
				<ElFormItem :label="t('user.user.name.phone')" prop="phone">
					<ElInput v-model="saveForm.data.phone" :placeholder="t('user.user.name.phone')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px;" />
					<label>
						<ElAlert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
					</label>
				</ElFormItem>
				<ElFormItem :label="t('user.user.name.account')" prop="account">
					<ElInput v-model="saveForm.data.account" :placeholder="t('user.user.name.account')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px;" />
					<label>
						<ElAlert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
					</label>
				</ElFormItem>
				<ElFormItem :label="t('user.user.name.password')" prop="password">
					<ElInput v-model="saveForm.data.password" :placeholder="t('user.user.name.password')" minlength="1" maxlength="32" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px;" />
					<label v-if="saveForm.data.idArr?.length">
						<ElAlert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
					</label>
				</ElFormItem>
				<ElFormItem :label="t('user.user.name.nickname')" prop="nickname">
					<ElInput v-model="saveForm.data.nickname" :placeholder="t('user.user.name.nickname')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('user.user.name.avatar')" prop="avatar">
					<MyUpload v-model="saveForm.data.avatar" accept="image/*" />
				</ElFormItem>
				<ElFormItem :label="t('user.user.name.gender')" prop="gender">
					<ElRadioGroup v-model="saveForm.data.gender">
						<ElRadio v-for="(item, index) in (tm('user.user.status.gender') as any)" :key="index" :label="item.value">
							{{ item.label }}
						</ElRadio>
					</ElRadioGroup>
				</ElFormItem>
				<ElFormItem :label="t('user.user.name.birthday')" prop="birthday">
					<ElDatePicker v-model="saveForm.data.birthday" type="date" :placeholder="t('user.user.name.birthday')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
				</ElFormItem>
				<ElFormItem :label="t('user.user.name.address')" prop="address">
					<ElInput v-model="saveForm.data.address" :placeholder="t('user.user.name.address')" minlength="1" maxlength="60" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('user.user.name.idCardName')" prop="idCardName">
					<ElInput v-model="saveForm.data.idCardName" :placeholder="t('user.user.name.idCardName')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('user.user.name.idCardNo')" prop="idCardNo">
					<ElInput v-model="saveForm.data.idCardNo" :placeholder="t('user.user.name.idCardNo')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('user.user.name.isStop')" prop="isStop">
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