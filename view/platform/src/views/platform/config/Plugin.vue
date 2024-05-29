<script setup lang="tsx">
const Upload = defineAsyncComponent(() => import('./plugin/Upload.vue'))
const Pay = defineAsyncComponent(() => import('./plugin/Pay.vue'))
const Sms = defineAsyncComponent(() => import('./plugin/Sms.vue'))
const IdCard = defineAsyncComponent(() => import('./plugin/IdCard.vue'))
const OneClick = defineAsyncComponent(() => import('./plugin/OneClick.vue'))
const Push = defineAsyncComponent(() => import('./plugin/Push.vue'))
const Vod = defineAsyncComponent(() => import('./plugin/Vod.vue'))
const Wx = defineAsyncComponent(() => import('./plugin/Wx.vue'))

const { t } = useI18n()
const adminStore = useAdminStore()

const authAction: { [propName: string]: boolean } = {
    isRead: adminStore.IsAction('platformConfigRead'),
    isSave: adminStore.IsAction('platformConfigSave'),
}
provide('authAction', authAction)
</script>

<template>
    <div v-if="!authAction.isRead" style="text-align: center; font-size: 60px; color: #f56c6c">{{ t('common.tip.notAuthActionRead') }}</div>
    <template v-else>
        <el-container class="common-container">
            <el-main>
                <el-tabs type="border-card" tab-position="top">
                    <el-tab-pane :label="t('platform.config.plugin.label.upload')" :lazy="true"><upload /></el-tab-pane>
                    <el-tab-pane :label="t('platform.config.plugin.label.pay')" :lazy="true"><pay /></el-tab-pane>
                    <el-tab-pane :label="t('platform.config.plugin.label.sms')" :lazy="true"><sms /></el-tab-pane>
                    <el-tab-pane :label="t('platform.config.plugin.label.idCard')" :lazy="true"><id-card /></el-tab-pane>
                    <el-tab-pane :label="t('platform.config.plugin.label.oneClick')" :lazy="true"><one-click /></el-tab-pane>
                    <el-tab-pane :label="t('platform.config.plugin.label.push')" :lazy="true"><push /></el-tab-pane>
                    <el-tab-pane :label="t('platform.config.plugin.label.vod')" :lazy="true"><vod /></el-tab-pane>
                    <el-tab-pane :label="t('platform.config.plugin.label.wx')" :lazy="true"><wx /></el-tab-pane>
                </el-tabs>
            </el-main>
        </el-container>
    </template>
</template>
