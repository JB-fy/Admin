<!--直接批量导入组件 开始-->
<!-- <script lang="ts">
const components: { [propName: string]: any } = batchImport(import.meta.glob('./tabPane/*.vue', { eager: true }))
export default {
    components: components,
}
</script> -->
<!--直接批量导入组件 结束-->
<script setup lang="ts">
//import WebsiteConfig from './tabPane/WebsiteConfig.vue'
//下面方式引入好处：组件会被打包成单独一个文件
const WebsiteConfig = defineAsyncComponent(() => import('./tabPane/WebsiteConfig.vue'))
const AppConfig = defineAsyncComponent(() => import('./tabPane/AppConfig.vue'))

const { t } = useI18n()
</script>

<template>
    <ElContainer class="common-container">
        <ElMain>
            <ElTabs type="border-card" tab-position="top">
                <ElTabPane :label="t('platform.config.label.websiteConfig')" :lazy="true">
                    <WebsiteConfig />
                </ElTabPane>
                <ElTabPane :label="t('platform.config.label.appConfig')" :lazy="true">
                    <AppConfig />
                </ElTabPane>
            </ElTabs>
        </ElMain>
    </ElContainer>
</template>
