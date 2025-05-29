<script setup lang="tsx">
const Sms = defineAsyncComponent(() => import('./plugin/Sms.vue'))
const Email = defineAsyncComponent(() => import('./plugin/Email.vue'))
const IdCard = defineAsyncComponent(() => import('./plugin/IdCard.vue'))
const OneClick = defineAsyncComponent(() => import('./plugin/OneClick.vue'))
const Push = defineAsyncComponent(() => import('./plugin/Push.vue'))
const Vod = defineAsyncComponent(() => import('./plugin/Vod.vue'))
const Wx = defineAsyncComponent(() => import('./plugin/Wx.vue'))

const { t } = useI18n()
const adminStore = useAdminStore()

const isRead = adminStore.isAction('platformConfigRead')
const isSave = adminStore.isAction('platformConfigSave')
const authAction: { [propName: string]: boolean } = {
    isSmsRead: isRead || adminStore.isAction('platformConfigSmsRead'),
    isSmsSave: isSave || adminStore.isAction('platformConfigSmsSave'),
    isEmailRead: isRead || adminStore.isAction('platformConfigEmailRead'),
    isEmailSave: isSave || adminStore.isAction('platformConfigEmailSave'),
    isIdCardRead: isRead || adminStore.isAction('platformConfigIdCardRead'),
    isIdCardSave: isSave || adminStore.isAction('platformConfigIdCardSave'),
    isOneClickRead: isRead || adminStore.isAction('platformConfigOneClickRead'),
    isOneClickSave: isSave || adminStore.isAction('platformConfigOneClickSave'),
    isPushRead: isRead || adminStore.isAction('platformConfigPushRead'),
    isPushSave: isSave || adminStore.isAction('platformConfigPushSave'),
    isVodRead: isRead || adminStore.isAction('platformConfigVodRead'),
    isVodSave: isSave || adminStore.isAction('platformConfigVodSave'),
    isWxRead: isRead || adminStore.isAction('platformConfigWxRead'),
    isWxSave: isSave || adminStore.isAction('platformConfigWxSave'),
}
provide('authAction', authAction)
const notReadAll = !(authAction.isSmsRead || authAction.isEmailRead || authAction.isIdCardRead || authAction.isOneClickRead || authAction.isPushRead || authAction.isVodRead || authAction.isWxRead)
</script>

<template>
    <div v-if="notReadAll" style="text-align: center; font-size: 60px; color: #f56c6c">{{ t('common.tip.notAuthActionRead') }}</div>
    <el-container v-else class="common-container">
        <el-main>
            <el-tabs type="border-card" tab-position="top">
                <el-tab-pane v-if="authAction.isSmsRead" :label="t('platform.config.plugin.label.sms')" :lazy="true"><sms /></el-tab-pane>
                <el-tab-pane v-if="authAction.isEmailRead" :label="t('platform.config.plugin.label.email')" :lazy="true"><email /></el-tab-pane>
                <el-tab-pane v-if="authAction.isIdCardRead" :label="t('platform.config.plugin.label.id_card')" :lazy="true"><id-card /></el-tab-pane>
                <el-tab-pane v-if="authAction.isOneClickRead" :label="t('platform.config.plugin.label.one_click')" :lazy="true"><one-click /></el-tab-pane>
                <el-tab-pane v-if="authAction.isPushRead" :label="t('platform.config.plugin.label.push')" :lazy="true"><push /></el-tab-pane>
                <el-tab-pane v-if="authAction.isVodRead" :label="t('platform.config.plugin.label.vod')" :lazy="true"><vod /></el-tab-pane>
                <el-tab-pane v-if="authAction.isWxRead" :label="t('platform.config.plugin.label.wx')" :lazy="true"><wx /></el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
