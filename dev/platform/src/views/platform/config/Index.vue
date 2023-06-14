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
const UploadConfig = defineAsyncComponent(() => import('./tabPane/UploadConfig.vue'))   //好处：该组件会被打包成单独一个文件

const { t } = useI18n()
</script>

<template>
    <ElContainer class="common-container">
        <ElMain>
            <ElTabs type="border-card" tab-position="top">
                <!-- <ElTabPane :label="t('view.platform.config.label.testConfig')" :lazy="true">
                    <UploadConfig />
                </ElTabPane> -->
                <ElTabPane :label="t('view.platform.config.label.uploadConfig')" :lazy="true">
                    <UploadConfig />
                </ElTabPane>
            </ElTabs>
        </ElMain>
    </ElContainer>
</template>