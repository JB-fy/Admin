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

const isRead = adminStore.isAction('pltCfgRead')
const isSave = adminStore.isAction('pltCfgSave')
const authAction: { [propName: string]: boolean } = {
    isSmsRead: isRead || adminStore.isAction('pltCfgSmsRead'),
    isSmsSave: isSave || adminStore.isAction('pltCfgSmsSave'),
    isEmailRead: isRead || adminStore.isAction('pltCfgEmailRead'),
    isEmailSave: isSave || adminStore.isAction('pltCfgEmailSave'),
    isIdCardRead: isRead || adminStore.isAction('pltCfgIdCardRead'),
    isIdCardSave: isSave || adminStore.isAction('pltCfgIdCardSave'),
    isOneClickRead: isRead || adminStore.isAction('pltCfgOneClickRead'),
    isOneClickSave: isSave || adminStore.isAction('pltCfgOneClickSave'),
    isPushRead: isRead || adminStore.isAction('pltCfgPushRead'),
    isPushSave: isSave || adminStore.isAction('pltCfgPushSave'),
    isVodRead: isRead || adminStore.isAction('pltCfgVodRead'),
    isVodSave: isSave || adminStore.isAction('pltCfgVodSave'),
    isWxRead: isRead || adminStore.isAction('pltCfgWxRead'),
    isWxSave: isSave || adminStore.isAction('pltCfgWxSave'),
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
