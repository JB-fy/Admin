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
const Platform = defineAsyncComponent(() => import('./app/Platform.vue'))
const Org = defineAsyncComponent(() => import('./app/Org.vue'))

const { t } = useI18n()
const adminStore = useAdminStore()

const isRead = adminStore.isAction('pltCfgRead')
const isSave = adminStore.isAction('pltCfgSave')
const authAction: { [propName: string]: boolean } = {
    isCommonRead: isRead || adminStore.isAction('pltCfgCommonRead'),
    isCommonSave: isSave || adminStore.isAction('pltCfgCommonSave'),
    isPlatformRead: isRead || adminStore.isAction('pltCfgPlatformRead'),
    isPlatformSave: isSave || adminStore.isAction('pltCfgPlatformSave'),
    isOrgRead: isRead || adminStore.isAction('pltCfgOrgRead'),
    isOrgSave: isSave || adminStore.isAction('pltCfgOrgSave'),
}
provide('authAction', authAction)
const notReadAll = !(authAction.isCommonRead || authAction.isPlatformRead || authAction.isOrgRead)
</script>

<template>
    <div v-if="notReadAll" style="text-align: center; font-size: 60px; color: #f56c6c">{{ t('common.tip.notAuthActionRead') }}</div>
    <el-container v-else class="common-container">
        <el-main>
            <el-tabs type="border-card" tab-position="top">
                <el-tab-pane v-if="authAction.isCommonRead" :label="t('platform.config.app.label.common')" :lazy="true"><common /></el-tab-pane>
                <el-tab-pane v-if="authAction.isPlatformRead" :label="t('platform.config.app.label.platform')" :lazy="true"><platform /></el-tab-pane>
                <el-tab-pane v-if="authAction.isOrgRead" :label="t('platform.config.app.label.org')" :lazy="true"><org /></el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
