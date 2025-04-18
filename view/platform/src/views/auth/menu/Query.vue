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
        <el-form-item prop="menu_name">
            <el-input v-model="queryCommon.data.menu_name" :placeholder="t('auth.menu.name.menu_name')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="scene_id">
            <my-select v-model="queryCommon.data.scene_id" :placeholder="t('auth.menu.name.scene_id')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list' }" @change="() => (queryCommon.data.pid = undefined)" />
        </el-form-item>
        <el-form-item v-if="queryCommon.data.scene_id" prop="pid">
            <my-cascader
                v-model="queryCommon.data.pid"
                :placeholder="t('auth.menu.name.pid')"
                :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { filter: { scene_id: queryCommon.data.scene_id } } }"
                :options="tm('common.status.pid')"
                :props="{ checkStrictly: true, emitPath: false }"
            />
        </el-form-item>
        <el-form-item prop="is_leaf">
            <el-select-v2 v-model="queryCommon.data.is_leaf" :options="tm('common.status.whether')" :placeholder="t('auth.menu.name.is_leaf')" :clearable="true" style="width: 86px" />
        </el-form-item>
        <el-form-item prop="level">
            <el-input-number v-model="queryCommon.data.level" :placeholder="t('auth.menu.name.level')" :min="1" :max="255" :precision="0" :controls="false" />
        </el-form-item>
        <el-form-item prop="is_stop">
            <el-select-v2 v-model="queryCommon.data.is_stop" :options="tm('common.status.whether')" :placeholder="t('auth.menu.name.is_stop')" :clearable="true" style="width: 86px" />
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
