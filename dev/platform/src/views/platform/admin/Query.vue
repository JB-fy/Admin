<script setup lang="ts">
const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
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
        <ElFormItem prop="account">
            <ElInput v-model="queryCommon.data.account" :placeholder="t('common.name.account')" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="phone">
            <ElInput v-model="queryCommon.data.phone" :placeholder="t('common.name.phone')" :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="roleId">
            <MySelect v-model="queryCommon.data.roleId" :placeholder="t('platform.admin.name.roleId')"
                :api="{ code: '/auth/role/list' }" />
        </ElFormItem>
        <ElFormItem prop="isStop" style="width: 100px;">
            <ElSelectV2 v-model="queryCommon.data.isStop" :options="tm('common.status.whether')"
                :placeholder="t('common.name.isStop')" :clearable="true" />
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