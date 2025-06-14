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
        <el-form-item prop="org_id">
            <my-select v-model="queryCommon.data.org_id" :placeholder="t('org.admin.name.org_id')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/org/org/list' }" />
        </el-form-item>
        <el-form-item prop="nickname">
            <el-input v-model="queryCommon.data.nickname" :placeholder="t('org.admin.name.nickname')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="phone">
            <el-input v-model="queryCommon.data.phone" :placeholder="t('org.admin.name.phone')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="email">
            <el-input v-model="queryCommon.data.email" :placeholder="t('org.admin.name.email')" maxlength="60" :clearable="true" />
        </el-form-item>
        <el-form-item prop="account">
            <el-input v-model="queryCommon.data.account" :placeholder="t('org.admin.name.account')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="is_super">
            <el-select-v2 v-model="queryCommon.data.is_super" :options="tm('common.status.whether')" :placeholder="t('org.admin.name.is_super')" :clearable="true" style="width: 86px" />
        </el-form-item>
        <el-form-item prop="role_id">
            <my-select v-model="queryCommon.data.role_id" :placeholder="t('org.admin.name.role_id_arr')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/role/list', param: { filter: { scene_id: `org` } } }" />
        </el-form-item>
        <el-form-item prop="is_stop">
            <el-select-v2 v-model="queryCommon.data.is_stop" :options="tm('common.status.whether')" :placeholder="t('org.admin.name.is_stop')" :clearable="true" style="width: 86px" />
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
