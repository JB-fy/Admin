<!--直接批量导入组件 开始-->
<!-- <script lang="tsx">
const components: { [propName: string]: any } = batchImport(import.meta.glob('./app/*.vue', { eager: true }))
export default {
    components: components,
}
</script> -->
<!--直接批量导入组件 结束-->
<script setup lang="tsx">
//import Common from './app/Common.vue'
//下面方式引入好处：组件会被打包成单独一个文件
const Common = defineAsyncComponent(() => import('./app/Common.vue'))

const { t } = useI18n()
const adminStore = useAdminStore()

const isRead = adminStore.isAction('orgCfgRead')
const isSave = adminStore.isAction('orgCfgSave')
const authAction: { [propName: string]: boolean } = {
    isCommonRead: isRead || adminStore.isAction('orgCfgCommonRead'),
    isCommonSave: isSave || adminStore.isAction('orgCfgCommonSave'),
}
provide('authAction', authAction)
const notReadAll = !authAction.isCommonRead
</script>

<template>
    <div v-if="notReadAll" style="text-align: center; font-size: 60px; color: #f56c6c">{{ t('common.tip.notAuthActionRead') }}</div>
    <el-container v-else class="common-container">
        <el-main>
            <el-tabs type="border-card" tab-position="top">
                <el-tab-pane v-if="authAction.isCommonRead" :label="t('org.config.app.label.common')" :lazy="true"><common /></el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
