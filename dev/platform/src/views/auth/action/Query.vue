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
        <ElFormItem prop="actionName">
            <ElInput v-model="queryCommon.data.actionName" :placeholder="t('view.auth.action.name.actionName')"
                :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="actionCode">
            <ElInput v-model="queryCommon.data.actionCode" :placeholder="t('view.auth.action.name.actionCode')"
                :clearable="true" />
        </ElFormItem>
        <ElFormItem prop="sceneId">
            <MySelect v-model="queryCommon.data.sceneId" :placeholder="t('common.name.rel.sceneId')"
                :api="{ code: '/auth/scene/list' }" />
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