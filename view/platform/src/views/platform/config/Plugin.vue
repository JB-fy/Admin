<script setup lang="tsx">
const Upload = defineAsyncComponent(() => import('./plugin/Upload.vue'))
const Sms = defineAsyncComponent(() => import('./plugin/Sms.vue'))
const Email = defineAsyncComponent(() => import('./plugin/Email.vue'))
const IdCard = defineAsyncComponent(() => import('./plugin/IdCard.vue'))
const OneClick = defineAsyncComponent(() => import('./plugin/OneClick.vue'))
const Push = defineAsyncComponent(() => import('./plugin/Push.vue'))
const Vod = defineAsyncComponent(() => import('./plugin/Vod.vue'))
const Wx = defineAsyncComponent(() => import('./plugin/Wx.vue'))

const { t } = useI18n()
const adminStore = useAdminStore()

const isRead = adminStore.IsAction('platformConfigRead')
const isSave = adminStore.IsAction('platformConfigSave')
const authAction: { [propName: string]: boolean } = {
    isUploadRead: isRead || adminStore.IsAction('platformConfigUploadRead'),
    isUploadSave: isSave || adminStore.IsAction('platformConfigUploadSave'),
    isSmsRead: isRead || adminStore.IsAction('platformConfigSmsRead'),
    isSmsSave: isSave || adminStore.IsAction('platformConfigSmsSave'),
    isEmailRead: isRead || adminStore.IsAction('platformConfigEmailRead'),
    isEmailSave: isSave || adminStore.IsAction('platformConfigEmailSave'),
    isIdCardRead: isRead || adminStore.IsAction('platformConfigIdCardRead'),
    isIdCardSave: isSave || adminStore.IsAction('platformConfigIdCardSave'),
    isOneClickRead: isRead || adminStore.IsAction('platformConfigOneClickRead'),
    isOneClickSave: isSave || adminStore.IsAction('platformConfigOneClickSave'),
    isPushRead: isRead || adminStore.IsAction('platformConfigPushRead'),
    isPushSave: isSave || adminStore.IsAction('platformConfigPushSave'),
    isVodRead: isRead || adminStore.IsAction('platformConfigVodRead'),
    isVodSave: isSave || adminStore.IsAction('platformConfigVodSave'),
    isWxRead: isRead || adminStore.IsAction('platformConfigWxRead'),
    isWxSave: isSave || adminStore.IsAction('platformConfigWxSave'),
}
provide('authAction', authAction)
const notReadAll = !(authAction.isUploadRead || authAction.isSmsRead || authAction.isEmailRead || authAction.isIdCardRead || authAction.isOneClickRead || authAction.isPushRead || authAction.isVodRead || authAction.isWxRead)
</script>

<template>
    <div v-if="notReadAll" style="text-align: center; font-size: 60px; color: #f56c6c">{{ t('common.tip.notAuthActionRead') }}</div>
    <el-container v-else class="common-container">
        <el-main>
            <el-tabs type="border-card" tab-position="top">
                <el-tab-pane v-if="authAction.isUploadRead" :label="t('platform.config.plugin.label.upload')" :lazy="true"><upload /></el-tab-pane>
                <el-tab-pane v-if="authAction.isSmsRead" :label="t('platform.config.plugin.label.sms')" :lazy="true"><sms /></el-tab-pane>
                <el-tab-pane v-if="authAction.isEmailRead" :label="t('platform.config.plugin.label.email')" :lazy="true"><email /></el-tab-pane>
                <el-tab-pane v-if="authAction.isIdCardRead" :label="t('platform.config.plugin.label.idCard')" :lazy="true"><id-card /></el-tab-pane>
                <el-tab-pane v-if="authAction.isOneClickRead" :label="t('platform.config.plugin.label.oneClick')" :lazy="true"><one-click /></el-tab-pane>
                <el-tab-pane v-if="authAction.isPushRead" :label="t('platform.config.plugin.label.push')" :lazy="true"><push /></el-tab-pane>
                <el-tab-pane v-if="authAction.isVodRead" :label="t('platform.config.plugin.label.vod')" :lazy="true"><vod /></el-tab-pane>
                <el-tab-pane v-if="authAction.isWxRead" :label="t('platform.config.plugin.label.wx')" :lazy="true"><wx /></el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
