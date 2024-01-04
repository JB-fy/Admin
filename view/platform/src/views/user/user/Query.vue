<script setup lang="ts">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
    ...queryCommon.data,
    timeRange: (() => {
        // const date = new Date()
        return [
            // new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
            // new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
        ]
    })(),
    timeRangeStart: computed(() => {
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[0]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),
    timeRangeEnd: computed(() => {
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[1]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),
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
    },
})
</script>

<template>
    <ElForm class="query-form" :ref="(el: any) => (queryForm.ref = el)" :model="queryCommon.data" :inline="true" @keyup.enter="queryForm.submit">
        <ElFormItem prop="id">
            <ElInputNumber v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
        </ElFormItem>
        <ElFormItem prop="phone">
            <ElInput v-model="queryCommon.data.phone" :placeholder="t('user.user.name.phone')" maxlength="30" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="account">
            <ElInput v-model="queryCommon.data.account" :placeholder="t('user.user.name.account')" maxlength="30" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="nickname">
            <ElInput v-model="queryCommon.data.nickname" :placeholder="t('user.user.name.nickname')" maxlength="30" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="gender" style="width: 120px">
            <ElSelectV2 v-model="queryCommon.data.gender" :options="tm('user.user.status.gender')" :placeholder="t('user.user.name.gender')" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="birthday">
            <ElDatePicker v-model="queryCommon.data.birthday" type="date" :placeholder="t('user.user.name.birthday')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
        </ElFormItem>
        <!-- <ElFormItem prop="address">
			<ElInput v-model="queryCommon.data.address" :placeholder="t('user.user.name.address')" maxlength="60" :clearable="true" />
		</ElFormItem> -->
        <ElFormItem prop="idCardName">
            <ElInput v-model="queryCommon.data.idCardName" :placeholder="t('user.user.name.idCardName')" maxlength="30" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="idCardNo">
            <ElInput v-model="queryCommon.data.idCardNo" :placeholder="t('user.user.name.idCardNo')" maxlength="30" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="isStop" style="width: 120px">
            <ElSelectV2 v-model="queryCommon.data.isStop" :options="tm('common.status.whether')" :placeholder="t('user.user.name.isStop')" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="timeRange">
            <ElDatePicker
                v-model="queryCommon.data.timeRange"
                type="datetimerange"
                range-separator="-"
                :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]"
                :start-placeholder="t('common.name.timeRangeStart')"
                :end-placeholder="t('common.name.timeRangeEnd')"
            />
        </ElFormItem>
        <ElFormItem>
            <ElButton type="primary" @click="queryForm.submit" :loading="queryForm.loading"> <AutoiconEpSearch />{{ t('common.query') }} </ElButton>
            <ElButton type="info" @click="queryForm.reset"> <AutoiconEpCircleClose />{{ t('common.reset') }} </ElButton>
        </ElFormItem>
    </ElForm>
</template>
