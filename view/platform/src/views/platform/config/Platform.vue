<!--直接批量导入组件 开始-->
<!-- <script lang="tsx">
const components: { [propName: string]: any } = batchImport(import.meta.glob('./tabPane/*.vue', { eager: true }))
export default {
    components: components,
}
</script> -->
<!--直接批量导入组件 结束-->
<script setup lang="tsx">
//import WebsiteConfig from './tabPane/WebsiteConfig.vue'
//下面方式引入好处：组件会被打包成单独一个文件
const WebsiteConfig = defineAsyncComponent(() => import('./tabPane/WebsiteConfig.vue'))
const AppConfig = defineAsyncComponent(() => import('./tabPane/AppConfig.vue'))

const { t } = useI18n()
</script>

<template>
    <el-container class="common-container">
        <el-main>
            <el-tabs type="border-card" tab-position="top">
                <el-tab-pane :label="t('platform.config.label.websiteConfig')" :lazy="true">
                    <website-config />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.label.appConfig')" :lazy="true">
                    <app-config />
                </el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
