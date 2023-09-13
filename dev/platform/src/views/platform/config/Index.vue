<!--直接批量导入组件 开始-->
<!-- <script lang="ts">
const components: { [propName: string]: any } = batchImport(import.meta.globEager('./tabPane/*.vue'))
export default {
    components: components,
}
</script> -->
<!--直接批量导入组件 结束-->
<script setup lang="ts">
//import UploadConfig from './tabPane/UploadConfig.vue'
//下面方式引入好处：该组件会被打包成单独一个文件
const WebsiteConfig = defineAsyncComponent(() => import('./tabPane/WebsiteConfig.vue'))
const UploadConfig = defineAsyncComponent(() => import('./tabPane/UploadConfig.vue'))
const SmsConfig = defineAsyncComponent(() => import('./tabPane/SmsConfig.vue'))
const IdCardConfig = defineAsyncComponent(() => import('./tabPane/IdCardConfig.vue'))

const { t } = useI18n()
</script>

<template>
    <ElContainer class="common-container">
        <ElMain>
            <ElTabs type="border-card" tab-position="top">
                <ElTabPane :label="t('platform.config.label.websiteConfig')" :lazy="true">
                    <WebsiteConfig />
                </ElTabPane>
                <ElTabPane :label="t('platform.config.label.uploadConfig')" :lazy="true">
                    <UploadConfig />
                </ElTabPane>
                <ElTabPane :label="t('platform.config.label.smsConfig')" :lazy="true">
                    <SmsConfig />
                </ElTabPane>
                <ElTabPane :label="t('platform.config.label.idCardConfig')" :lazy="true">
                    <IdCardConfig />
                </ElTabPane>
            </ElTabs>
        </ElMain>
    </ElContainer>
</template>