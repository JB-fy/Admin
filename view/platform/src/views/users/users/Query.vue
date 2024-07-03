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
            <el-input-number v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :max="4294967295" :controls="false" />
        </el-form-item>
        <el-form-item prop="nickname">
            <el-input v-model="queryCommon.data.nickname" :placeholder="t('users.users.name.nickname')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="gender">
            <el-select-v2 v-model="queryCommon.data.gender" :options="tm('users.users.status.gender')" :placeholder="t('users.users.name.gender')" :clearable="true" style="width: 100px" />
        </el-form-item>
        <el-form-item prop="birthday">
            <el-date-picker v-model="queryCommon.data.birthday" type="date" :placeholder="t('users.users.name.birthday')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 160px" />
        </el-form-item>
        <el-form-item prop="phone">
            <el-input v-model="queryCommon.data.phone" :placeholder="t('users.users.name.phone')" maxlength="20" :clearable="true" />
        </el-form-item>
        <el-form-item prop="email">
            <el-input v-model="queryCommon.data.email" :placeholder="t('users.users.name.email')" maxlength="60" :clearable="true" />
        </el-form-item>
        <el-form-item prop="account">
            <el-input v-model="queryCommon.data.account" :placeholder="t('users.users.name.account')" maxlength="20" :clearable="true" />
        </el-form-item>
        <el-form-item prop="wx_openid">
            <el-input v-model="queryCommon.data.wx_openid" :placeholder="t('users.users.name.wx_openid')" maxlength="128" :clearable="true" />
        </el-form-item>
        <el-form-item prop="wx_unionid">
            <el-input v-model="queryCommon.data.wx_unionid" :placeholder="t('users.users.name.wx_unionid')" maxlength="64" :clearable="true" />
        </el-form-item>
        <el-form-item prop="id_card_no">
            <el-input v-model="queryCommon.data.id_card_no" :placeholder="t('users.users.name.id_card_no')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="id_card_name">
            <el-input v-model="queryCommon.data.id_card_name" :placeholder="t('users.users.name.id_card_name')" maxlength="30" :clearable="true" />
        </el-form-item>
        <el-form-item prop="id_card_gender">
            <el-select-v2 v-model="queryCommon.data.id_card_gender" :options="tm('users.users.status.id_card_gender')" :placeholder="t('users.users.name.id_card_gender')" :clearable="true" style="width: 128px" />
        </el-form-item>
        <el-form-item prop="id_card_birthday">
            <el-date-picker v-model="queryCommon.data.id_card_birthday" type="date" :placeholder="t('users.users.name.id_card_birthday')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 160px" />
        </el-form-item>
        <el-form-item prop="is_stop">
            <el-select-v2 v-model="queryCommon.data.is_stop" :options="tm('common.status.whether')" :placeholder="t('users.users.name.is_stop')" :clearable="true" style="width: 86px" />
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
            <el-button type="primary" @click="queryForm.submit" :loading="queryForm.loading"> <autoicon-ep-search />{{ t('common.query') }}</el-button>
            <el-button type="info" @click="queryForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
