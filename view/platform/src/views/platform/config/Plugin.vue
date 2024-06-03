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

const isRead = adminStore.IsAction('platformConfigRead')
const isSave = adminStore.IsAction('platformConfigSave')
const authAction: { [propName: string]: boolean } = {
    isUploadRead: isRead || adminStore.IsAction('platformConfigUploadRead'),
    isUploadSave: isSave || adminStore.IsAction('platformConfigUploadSave'),
    isPayRead: isRead || adminStore.IsAction('platformConfigPayRead'),
    isPaySave: isSave || adminStore.IsAction('platformConfigPaySave'),
    isSmsRead: isRead || adminStore.IsAction('platformConfigSmsRead'),
    isSmsSave: isSave || adminStore.IsAction('platformConfigSmsSave'),
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
</script>

<template>
    <div v-if="!(authAction.isUploadRead || authAction.isPayRead || authAction.isSmsRead || authAction.isIdCardRead || authAction.isOneClickRead || authAction.isPushRead || authAction.isVodRead || authAction.isWxRead)"
        style="text-align: center; font-size: 60px; color: #f56c6c">
        {{ t('common.tip.notAuthActionRead') }}
    </div>
    <template v-else>
        <el-container class="common-container">
            <el-main>
                <el-tabs type="border-card" tab-position="top">
                    <el-tab-pane v-if="authAction.isUploadRead" :label="t('platform.config.plugin.label.upload')"
                        :lazy="true">
                        <upload />
                    </el-tab-pane>
                    <el-tab-pane v-if="authAction.isPayRead" :label="t('platform.config.plugin.label.pay')"
                        :lazy="true">
                        <pay />
                    </el-tab-pane>
                    <el-tab-pane v-if="authAction.isSmsRead" :label="t('platform.config.plugin.label.sms')"
                        :lazy="true">
                        <sms />
                    </el-tab-pane>
                    <el-tab-pane v-if="authAction.isIdCardRead" :label="t('platform.config.plugin.label.idCard')"
                        :lazy="true"><id-card /></el-tab-pane>
                    <el-tab-pane v-if="authAction.isOneClickRead" :label="t('platform.config.plugin.label.oneClick')"
                        :lazy="true"><one-click /></el-tab-pane>
                    <el-tab-pane v-if="authAction.isPushRead" :label="t('platform.config.plugin.label.push')"
                        :lazy="true">
                        <push />
                    </el-tab-pane>
                    <el-tab-pane v-if="authAction.isVodRead" :label="t('platform.config.plugin.label.vod')"
                        :lazy="true">
                        <vod />
                    </el-tab-pane>
                    <el-tab-pane v-if="authAction.isWxRead" :label="t('platform.config.plugin.label.wx')" :lazy="true">
                        <wx />
                    </el-tab-pane>
                </el-tabs>
            </el-main>
        </el-container>
    </template>
</template>
