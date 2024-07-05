<!--直接批量导入组件 开始-->
<!-- <script lang="tsx">
const components: { [propName: string]: any } = batchImport(import.meta.glob('./platform/*.vue', { eager: true }))
export default {
    components: components,
}
</script> -->
<!--直接批量导入组件 结束-->
<script setup lang="tsx">
//import WebsiteConfig from './app/WebsiteConfig.vue'
//下面方式引入好处：组件会被打包成单独一个文件
const Website = defineAsyncComponent(() => import('./app/Website.vue'))

const { t } = useI18n()
const adminStore = useAdminStore()

const isRead = adminStore.IsAction('platformConfigRead')
const isSave = adminStore.IsAction('platformConfigSave')
const authAction: { [propName: string]: boolean } = {
    isWebsiteRead: isRead || adminStore.IsAction('platformConfigWebsiteRead'),
    isWebsiteSave: isSave || adminStore.IsAction('platformConfigWebsiteSave'),
}
provide('authAction', authAction)
const notReadAll = !(authAction.isWebsiteRead || authAction.isAppRead)
</script>

<template>
    <div v-if="notReadAll" style="text-align: center; font-size: 60px; color: #f56c6c">{{ t('common.tip.notAuthActionRead') }}</div>
    <el-container v-else class="common-container">
        <el-main>
            <el-tabs type="border-card" tab-position="top">
                <el-tab-pane v-if="authAction.isWebsiteRead" :label="t('platform.config.platform.label.website')" :lazy="true"><website /></el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
