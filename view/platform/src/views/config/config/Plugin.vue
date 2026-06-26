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

const isRead = adminStore.isAction('configRead')
const isSave = adminStore.isAction('configSave')
const authAction: { [propName: string]: boolean } = {
    isSmsRead: isRead || adminStore.isAction('configSmsRead'),
    isSmsSave: isSave || adminStore.isAction('configSmsSave'),
    isEmailRead: isRead || adminStore.isAction('configEmailRead'),
    isEmailSave: isSave || adminStore.isAction('configEmailSave'),
    isIdCardRead: isRead || adminStore.isAction('configIdCardRead'),
    isIdCardSave: isSave || adminStore.isAction('configIdCardSave'),
    isOneClickRead: isRead || adminStore.isAction('configOneClickRead'),
    isOneClickSave: isSave || adminStore.isAction('configOneClickSave'),
    isPushRead: isRead || adminStore.isAction('configPushRead'),
    isPushSave: isSave || adminStore.isAction('configPushSave'),
    isVodRead: isRead || adminStore.isAction('configVodRead'),
    isVodSave: isSave || adminStore.isAction('configVodSave'),
    isWxRead: isRead || adminStore.isAction('configWxRead'),
    isWxSave: isSave || adminStore.isAction('configWxSave'),
}
provide('authAction', authAction)
const notReadAll = !(authAction.isSmsRead || authAction.isEmailRead || authAction.isIdCardRead || authAction.isOneClickRead || authAction.isPushRead || authAction.isVodRead || authAction.isWxRead)
</script>

<template>
    <div v-if="notReadAll" style="text-align: center; font-size: 60px; color: #f56c6c">{{ t('common.tip.notAuthActionRead') }}</div>
    <el-container v-else class="common-container">
        <el-main>
            <el-tabs type="border-card" tab-position="top">
                <el-tab-pane v-if="authAction.isSmsRead" :label="t('config.config.plugin.label.sms')" :lazy="true"><sms /></el-tab-pane>
                <el-tab-pane v-if="authAction.isEmailRead" :label="t('config.config.plugin.label.email')" :lazy="true"><email /></el-tab-pane>
                <el-tab-pane v-if="authAction.isIdCardRead" :label="t('config.config.plugin.label.id_card')" :lazy="true"><id-card /></el-tab-pane>
                <el-tab-pane v-if="authAction.isOneClickRead" :label="t('config.config.plugin.label.one_click')" :lazy="true"><one-click /></el-tab-pane>
                <el-tab-pane v-if="authAction.isPushRead" :label="t('config.config.plugin.label.push')" :lazy="true"><push /></el-tab-pane>
                <el-tab-pane v-if="authAction.isVodRead" :label="t('config.config.plugin.label.vod')" :lazy="true"><vod /></el-tab-pane>
                <el-tab-pane v-if="authAction.isWxRead" :label="t('config.config.plugin.label.wx')" :lazy="true"><wx /></el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
