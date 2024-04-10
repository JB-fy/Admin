<script setup lang="tsx">
const Upload = defineAsyncComponent(() => import('./plugin/Upload.vue'))
const Pay = defineAsyncComponent(() => import('./plugin/Pay.vue'))
const Sms = defineAsyncComponent(() => import('./plugin/Sms.vue'))
const IdCard = defineAsyncComponent(() => import('./plugin/IdCard.vue'))
const OneClick = defineAsyncComponent(() => import('./plugin/OneClick.vue'))
const Push = defineAsyncComponent(() => import('./plugin/Push.vue'))
const Vod = defineAsyncComponent(() => import('./plugin/Vod.vue'))

const { t } = useI18n()
</script>

<template>
    <el-container class="common-container">
        <el-main>
            <el-tabs type="border-card" tab-position="top">
                <el-tab-pane :label="t('platform.config.plugin.label.upload')" :lazy="true">
                    <upload />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.plugin.label.pay')" :lazy="true">
                    <pay />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.plugin.label.sms')" :lazy="true">
                    <sms />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.plugin.label.idCard')" :lazy="true">
                    <id-card />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.plugin.label.oneClick')" :lazy="true">
                    <one-click />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.plugin.label.push')" :lazy="true">
                    <push />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.plugin.label.vod')" :lazy="true">
                    <vod />
                </el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
