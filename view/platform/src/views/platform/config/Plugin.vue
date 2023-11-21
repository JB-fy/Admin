<script setup lang="ts">
const UploadConfig = defineAsyncComponent(() => import('./tabPane/UploadConfig.vue'))
const SmsConfig = defineAsyncComponent(() => import('./tabPane/SmsConfig.vue'))
const IdCardConfig = defineAsyncComponent(() => import('./tabPane/IdCardConfig.vue'))
const PushConfig = defineAsyncComponent(() => import('./tabPane/PushConfig.vue'))

const { t } = useI18n()
</script>

<template>
    <ElContainer class="common-container">
        <ElMain>
            <ElTabs type="border-card" tab-position="top">
                <ElTabPane :label="t('platform.config.label.uploadConfig')" :lazy="true">
                    <UploadConfig />
                </ElTabPane>
                <ElTabPane :label="t('platform.config.label.smsConfig')" :lazy="true">
                    <SmsConfig />
                </ElTabPane>
                <ElTabPane :label="t('platform.config.label.idCardConfig')" :lazy="true">
                    <IdCardConfig />
                </ElTabPane>
                <ElTabPane :label="t('platform.config.label.pushConfig')" :lazy="true">
                    <PushConfig />
                </ElTabPane>
            </ElTabs>
        </ElMain>
    </ElContainer>
</template>