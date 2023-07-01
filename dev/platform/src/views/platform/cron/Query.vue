<script setup lang="ts">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
	...queryCommon.data,
	timeRange: (() => {
		const date = new Date()
		return [
			new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
			new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
		]
	})(),
	startTime: computed(() => {
		if (queryCommon.data.timeRange?.length) {
			return dayjs(queryCommon.data.timeRange[0]).format('YYYY-MM-DD HH:mm:ss')
		}
		return ''
	}),
	endTime: computed(() => {
		if (queryCommon.data.timeRange?.length) {
			return dayjs(queryCommon.data.timeRange[1]).format('YYYY-MM-DD HH:mm:ss')
		}
		return ''
	})
}
const listCommon = inject('listCommon') as { ref: any }
const queryForm = reactive({
	ref: null as any,
	loading: false,
	submit: () => {
		queryForm.loading = true
		listCommon.ref.getList(true).finally(() => {
			queryForm.loading = false
		})
	},
	reset: () => {
		queryForm.ref.resetFields()
		//queryForm.submit()
	}
})
</script>

<template>
	<ElForm class="query-form" :ref="(el: any) => { queryForm.ref = el }" :model="queryCommon.data" :inline="true"
		@keyup.enter="queryForm.submit">
		<ElFormItem prop="id">
			<ElInputNumber v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
		</ElFormItem>
		<ElFormItem prop="cronName">
			<ElInput v-model="queryCommon.data.cronName" :placeholder="t('platform.cron.name.cronName')" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="cronCode">
			<ElInput v-model="queryCommon.data.cronCode" :placeholder="t('platform.cron.name.cronCode')" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="cronPattern">
			<ElInput v-model="queryCommon.data.cronPattern" :placeholder="t('platform.cron.name.cronPattern')" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="isStop" style="width: 100px;">
			<ElSelectV2 v-model="queryCommon.data.isStop" :options="tm('common.status.whether')" :placeholder="t('common.name.isStop')" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="timeRange">
			<ElDatePicker v-model="queryCommon.data.timeRange" type="datetimerange" range-separator="-"
				:default-time="queryCommon.data.timeRange" :start-placeholder="t('common.name.startTime')"
				:end-placeholder="t('common.name.endTime')">
			</ElDatePicker>
		</ElFormItem>
		<ElFormItem>
			<ElButton type="primary" @click="queryForm.submit" :loading="queryForm.loading">
				<AutoiconEpSearch />{{ t('common.query') }}
			</ElButton>
			<ElButton type="info" @click="queryForm.reset">
				<AutoiconEpCircleClose />{{ t('common.reset') }}
			</ElButton>
		</ElFormItem>
	</ElForm>
</template>