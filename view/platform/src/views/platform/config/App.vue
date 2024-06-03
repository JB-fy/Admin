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
const App = defineAsyncComponent(() => import('./app/App.vue'))

const { t } = useI18n()
const adminStore = useAdminStore()

const isRead = adminStore.IsAction('platformConfigRead')
const isSave = adminStore.IsAction('platformConfigSave')
const authAction: { [propName: string]: boolean } = {
    isWebsiteRead: isRead || adminStore.IsAction('platformConfigWebsiteRead'),
    isWebsiteSave: isSave || adminStore.IsAction('platformConfigWebsiteSave'),
    isAppRead: isRead || adminStore.IsAction('platformConfigAppRead'),
    isAppSave: isSave || adminStore.IsAction('platformConfigAppSave'),
}
provide('authAction', authAction)
</script>

<template>
    <div v-if="!(authAction.isWebsiteRead || authAction.isAppRead)"
        style="text-align: center; font-size: 60px; color: #f56c6c">{{ t('common.tip.notAuthActionRead') }}</div>
    <template v-else>
        <el-container class="common-container">
            <el-main>
                <el-tabs type="border-card" tab-position="top">
                    <el-tab-pane v-if="authAction.isWebsiteRead" :label="t('platform.config.platform.label.website')"
                        :lazy="true">
                        <website />
                    </el-tab-pane>
                    <el-tab-pane v-if="authAction.isAppRead" :label="t('platform.config.platform.label.app')"
                        :lazy="true">
                        <app />
                    </el-tab-pane>
                </el-tabs>
            </el-main>
        </el-container>
    </template>
</template>
