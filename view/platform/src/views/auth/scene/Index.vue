<script setup lang="tsx">
import List from './List.vue'
import Query from './Query.vue'
import Save from './Save.vue'

const { t } = useI18n()
const adminStore = useAdminStore()

const authAction: { [propName: string]: boolean } = {
    isRead: adminStore.IsAction('authSceneRead'),
    isCreate: adminStore.IsAction('authSceneCreate'),
    isUpdate: adminStore.IsAction('authSceneUpdate'),
    isDelete: adminStore.IsAction('authSceneDelete'),
}
provide('authAction', authAction)

//搜索
const queryCommon = reactive({
    data: {},
})
provide('queryCommon', queryCommon)

//列表
const listCommon = reactive({
    ref: null as any,
})
provide('listCommon', listCommon)

//保存
const saveCommon = reactive({
    visible: false,
    title: '', //新增|编辑|复制
    data: {},
})
provide('saveCommon', saveCommon)
</script>

<template>
    <div v-if="!authAction.isRead" style="text-align: center; font-size: 60px; color: #f56c6c">{{ t('common.tip.notAuthActionRead') }}</div>
    <template v-else>
        <el-container class="main-table-container">
            <el-header>
                <query />
            </el-header>

            <list :ref="(el: any) => listCommon.ref = el" />

            <!-- 加上v-if每次都重新生成组件。可防止不同操作之间的影响；新增操作数据的默认值也能写在save组件内 -->
            <save v-if="saveCommon.visible" />
        </el-container>
    </template>
</template>
