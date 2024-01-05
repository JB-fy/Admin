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
    <ElForm class="query-form" :ref="(el: any) => (queryForm.ref = el)" :model="queryCommon.data" :inline="true" @keyup.enter="queryForm.submit">
        <ElFormItem prop="id">
            <ElInputNumber v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
        </ElFormItem>
        <ElFormItem prop="phone">
            <ElInput v-model="queryCommon.data.phone" :placeholder="t('platform.admin.name.phone')" maxlength="30" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="account">
            <ElInput v-model="queryCommon.data.account" :placeholder="t('platform.admin.name.account')" maxlength="30" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="nickname">
            <ElInput v-model="queryCommon.data.nickname" :placeholder="t('platform.admin.name.nickname')" maxlength="30" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="roleId">
            <MySelect v-model="queryCommon.data.roleId" :placeholder="t('platform.admin.name.roleId')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/role/list' }" />
        </ElFormItem>
        <ElFormItem prop="isStop" style="width: 120px">
            <ElSelectV2 v-model="queryCommon.data.isStop" :options="tm('common.status.whether')" :placeholder="t('platform.admin.name.isStop')" :clearable="true" />
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
