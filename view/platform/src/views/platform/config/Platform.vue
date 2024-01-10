<!--直接批量导入组件 开始-->
<!-- <script lang="tsx">
const components: { [propName: string]: any } = batchImport(import.meta.glob('./platform/*.vue', { eager: true }))
export default {
    components: components,
}
</script> -->
<!--直接批量导入组件 结束-->
<script setup lang="tsx">
//import WebsiteConfig from './platform/WebsiteConfig.vue'
//下面方式引入好处：组件会被打包成单独一个文件
const Website = defineAsyncComponent(() => import('./platform/Website.vue'))
const App = defineAsyncComponent(() => import('./platform/App.vue'))

const { t } = useI18n()
</script>

<template>
    <el-container class="common-container">
        <el-main>
            <el-tabs type="border-card" tab-position="top">
                <el-tab-pane :label="t('platform.config.platform.label.website')" :lazy="true">
                    <website />
                </el-tab-pane>
                <el-tab-pane :label="t('platform.config.platform.label.app')" :lazy="true">
                    <app />
                </el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>
