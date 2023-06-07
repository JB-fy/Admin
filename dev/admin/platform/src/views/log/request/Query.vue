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
        //return queryCommon.data.timeRange?.length ? queryCommon.data.timeRange[0] : '' //如果接口接受任何日期格式，不需要转换
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[0]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),
    endTime: computed(() => {
        //return queryCommon.data.timeRange?.length ? queryCommon.data.timeRange[1] : '' //如果接口接受任何日期格式，不需要转换
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
            <ElInputNumber v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1"
                :controls="false" />
        </ElFormItem>
        <ElFormItem prop="requestUrl">
            <ElInput v-model="queryCommon.data.requestUrl" :placeholder="t('common.name.log.request.requestUrl')"
                :clearable="true" />
        </ElFormItem>
        <ElFormItem style="width: 196px;" prop="minRunTime">
            <ElInput v-model="queryCommon.data.minRunTime" :placeholder="t('common.name.min')">
                <template #prepend>{{ t('common.name.log.request.runTime') }}</template>
            </ElInput>
        </ElFormItem>
        <ElFormItem>-</ElFormItem>
        <ElFormItem style="width: 100px;" prop="maxRunTime">
            <ElInput v-model="queryCommon.data.maxRunTime" :placeholder="t('common.name.max')" />
        </ElFormItem>
        <ElFormItem prop="timeRange">
            <ElDatePicker v-model="queryCommon.data.timeRange" type="datetimerange" range-separator="-"
                :start-placeholder="t('common.name.startTime')" :end-placeholder="t('common.name.endTime')">
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