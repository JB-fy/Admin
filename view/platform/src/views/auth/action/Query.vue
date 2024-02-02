<script setup lang="tsx">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
    ...queryCommon.data,
    timeRange: (() => {
        return undefined
        /* const date = new Date()
        return [
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
        ] */
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
    <el-form class="query-form" :ref="(el: any) => queryForm.ref = el" :model="queryCommon.data" :inline="true" @keyup.enter="queryForm.submit">
        <el-form-item prop="id">
            <el-input-number v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
        </el-form-item>
        <el-form-item prop="actionName">
            <el-input v-model="queryCommon.data.actionName" :placeholder="t('auth.action.name.actionName')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="actionCode">
            <el-input v-model="queryCommon.data.actionCode" :placeholder="t('auth.action.name.actionCode')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="sceneId">
            <my-select v-model="queryCommon.data.sceneId" :placeholder="t('auth.action.name.sceneId')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list' }" />
        </el-form-item>
        <el-form-item prop="isStop" style="width: 120px">
            <el-select-v2 v-model="queryCommon.data.isStop" :options="tm('common.status.whether')" :placeholder="t('auth.action.name.isStop')" :clearable="true" />
        </el-form-item>
        <el-form-item prop="timeRange">
            <el-date-picker
                v-model="queryCommon.data.timeRange"
                type="datetimerange"
                range-separator="-"
                :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]"
                :start-placeholder="t('common.name.timeRangeStart')"
                :end-placeholder="t('common.name.timeRangeEnd')"
            />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="queryForm.submit" :loading="queryForm.loading"> <autoicon-ep-search />{{ t('common.query') }} </el-button>
            <el-button type="info" @click="queryForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
