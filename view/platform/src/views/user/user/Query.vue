<script setup lang="tsx">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
    ...queryCommon.data,
    time_range: (() => {
        return undefined
        /* const date = new Date()
        return [
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
        ] */
    })(),
    time_range_start: computed(() => {
        if (queryCommon.data.time_range?.length) {
            return dayjs(queryCommon.data.time_range[0]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),
    time_range_end: computed(() => {
        if (queryCommon.data.time_range?.length) {
            return dayjs(queryCommon.data.time_range[1]).format('YYYY-MM-DD HH:mm:ss')
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
            <el-input-number v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :max="4294967295" :controls="false" />
        </el-form-item>
        <el-form-item prop="phone">
            <el-input v-model="queryCommon.data.phone" :placeholder="t('user.user.name.phone')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="account">
            <el-input v-model="queryCommon.data.account" :placeholder="t('user.user.name.account')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="nickname">
            <el-input v-model="queryCommon.data.nickname" :placeholder="t('user.user.name.nickname')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="gender">
            <el-select-v2 v-model="queryCommon.data.gender" :options="tm('user.user.status.gender')" :placeholder="t('user.user.name.gender')" :clearable="true" style="width: 100px" />
        </el-form-item>
        <el-form-item prop="birthday">
            <el-date-picker v-model="queryCommon.data.birthday" type="date" :placeholder="t('user.user.name.birthday')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 160px" />
        </el-form-item>
        <el-form-item prop="open_id_of_wx">
            <el-input v-model="queryCommon.data.open_id_of_wx" :placeholder="t('user.user.name.open_id_of_wx')" maxlength="128" :clearable="true" />
        </el-form-item>
        <el-form-item prop="union_id_of_wx">
            <el-input v-model="queryCommon.data.union_id_of_wx" :placeholder="t('user.user.name.union_id_of_wx')" maxlength="64" :clearable="true" />
        </el-form-item>
        <el-form-item prop="id_card_name">
            <el-input v-model="queryCommon.data.id_card_name" :placeholder="t('user.user.name.id_card_name')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="id_card_no">
            <el-input v-model="queryCommon.data.id_card_no" :placeholder="t('user.user.name.id_card_no')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="is_stop">
            <el-select-v2 v-model="queryCommon.data.is_stop" :options="tm('common.status.whether')" :placeholder="t('user.user.name.is_stop')" :clearable="true" style="width: 86px" />
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
            <el-button type="primary" @click="queryForm.submit" :loading="queryForm.loading"> <autoicon-ep-search />{{ t('common.query') }} </el-button>
            <el-button type="info" @click="queryForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
