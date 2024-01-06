<script setup lang="tsx">
const UploadConfig = defineAsyncComponent(() => import('./tabPane/UploadConfig.vue'))
const PayConfig = defineAsyncComponent(() => import('./tabPane/PayConfig.vue'))
const SmsConfig = defineAsyncComponent(() => import('./tabPane/SmsConfig.vue'))
const IdCardConfig = defineAsyncComponent(() => import('./tabPane/IdCardConfig.vue'))
const PushConfig = defineAsyncComponent(() => import('./tabPane/PushConfig.vue'))
const VodConfig = defineAsyncComponent(() => import('./tabPane/VodConfig.vue'))

const { t } = useI18n()
</script>

<template>
    <el-container class="common-container">
        <el-main>
            <el-tabs type="border-card" tab-position="top">
                <el-tab-pane :label="t('platform.config.label.uploadConfig')" :lazy="true">
                    <upload-config />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.label.payConfig')" :lazy="true">
                    <pay-config />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.label.smsConfig')" :lazy="true">
                    <sms-config />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.label.idCardConfig')" :lazy="true">
                    <id-card-config />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.label.pushConfig')" :lazy="true">
                    <push-config />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.label.vodConfig')" :lazy="true">
                    <vod-config />
                </el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
