<script setup lang="tsx">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
    ...queryCommon.data,
    time_range: undefined, //[new Date().setHours(0, 0, 0), new Date().setHours(23, 59, 59)]
    time_range_start: computed(() => (queryCommon.data.time_range?.length ? dayjs(queryCommon.data.time_range[0]).format('YYYY-MM-DD HH:mm:ss') : undefined)),
    time_range_end: computed(() => (queryCommon.data.time_range?.length ? dayjs(queryCommon.data.time_range[1]).format('YYYY-MM-DD HH:mm:ss') : undefined)),
}
const listCommon = inject('listCommon') as { ref: any }
const queryForm = reactive({
    ref: null as any,
    loading: false,
    submit: () => {
        queryForm.loading = true
        listCommon.ref.getList(true).finally(() => (queryForm.loading = false))
    },
    reset: () => queryForm.ref.resetFields(),
})
</script>

<template>
    <el-form class="query-form" :ref="(el: any) => queryForm.ref = el" :model="queryCommon.data" :inline="true" @keyup.enter="queryForm.submit">
        <el-form-item prop="id">
            <el-input-number v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :max="4294967295" :precision="0" :controls="false" />
        </el-form-item>
        <el-form-item prop="app_id">
            <my-select v-model="queryCommon.data.app_id" :placeholder="t('app.pkg.name.app_id')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/app/app/list' }" />
        </el-form-item>
        <el-form-item prop="pkg_type">
            <el-select-v2 v-model="queryCommon.data.pkg_type" :options="tm('app.pkg.status.pkg_type')" :placeholder="t('app.pkg.name.pkg_type')" :clearable="true" style="width: 86px" />
        </el-form-item>
        <el-form-item prop="ver_no">
            <el-input-number v-model="queryCommon.data.ver_no" :placeholder="t('app.pkg.name.ver_no')" :min="0" :max="4294967295" :precision="0" :controls="false" />
        </el-form-item>
        <el-form-item prop="ver_name">
            <el-input v-model="queryCommon.data.ver_name" :placeholder="t('app.pkg.name.ver_name')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="is_force_prev">
            <el-select-v2 v-model="queryCommon.data.is_force_prev" :options="tm('common.status.whether')" :placeholder="t('app.pkg.name.is_force_prev')" :clearable="true" style="width: 114px" />
        </el-form-item>
        <el-form-item prop="is_stop">
            <el-select-v2 v-model="queryCommon.data.is_stop" :options="tm('common.status.whether')" :placeholder="t('app.pkg.name.is_stop')" :clearable="true" style="width: 86px" />
        </el-form-item>
        <el-form-item prop="time_range">
            <el-date-picker
                v-model="queryCommon.data.time_range"
                type="datetimerange"
                range-separator="-"
                :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]"
                :start-placeholder="t('common.name.timeRangeStart')"
                :end-placeholder="t('common.name.timeRangeEnd')"
            />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="queryForm.submit" :loading="queryForm.loading"><autoicon-ep-search />{{ t('common.query') }}</el-button>
            <el-button type="info" @click="queryForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
