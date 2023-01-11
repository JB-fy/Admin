<!--直接批量导入组件 开始-->
<!-- <script lang="ts">
const components: { [propName: string]: any } = batchImport(import.meta.globEager('./tabPane/*.vue'))
export default {
    components: components,
}
</script> -->
<!--直接批量导入组件 结束-->
<script setup lang="ts">
//import Demo from './tabPane/Demo.vue'
const Demo = defineAsyncComponent(() => import('./tabPane/Demo.vue'))    //好处：该组件会被打包成单独一个文件

const { t } = useI18n()
</script>

<template>
    <ElContainer class="common-container">
        <ElMain>
            <ElTabs type="border-card" tab-position="top">
                <ElTabPane :label="t('view.platform.config.label.test')" :lazy="true">
                    <Demo />
                </ElTabPane>
            </ElTabs>
        </ElMain>
    </ElContainer>
</template>