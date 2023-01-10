<!--直接批量导入组件 开始-->
<!-- <script lang="ts">
const components: { [propName: string]: any } = await batchImport(import.meta.globEager('./tabPane/*.vue'))
export default {
    components: components,
}
</script> -->
<!--直接批量导入组件 结束-->
<script setup lang="ts">
//import AdminConfig from './tabPane/AdminConfig.vue'
const AdminConfig  = defineAsyncComponent(() => import('./tabPane/AdminConfig.vue'))    //好处：该组件会被打包成单独一个文件

// console.log(batchImport(import.meta.globEager('./tabPane/*.vue')))
// console.log(batchImport(import.meta.glob('./tabPane/*.vue')))

const tabs = reactive({
    loading: false,
    //activeTabName: 'first',
    tabPaneKey: 0,  //标签切换后重置tabPane
    change: (/* tabName: string */) => {
        //tabs.loading = true
        //console.log(tabName)
        tabs.tabPaneKey++
    },
})
</script>

<template>
    <ElContainer class="common-container">
        <ElMain>
            <!-- <ElTabs type="border-card" tab-position="top" v-model="tabs.activeTabName" @tab-change="tabs.change">
                <ElTabPane label="后台" name="first" :lazy="true" :key="tabs.tabPaneKey">
                    <AdminConfig />
                </ElTabPane>
                <ElTabPane label="代理" name="second" :lazy="true">
                    <AdminConfig />
                </ElTabPane>
            </ElTabs> -->
            <ElTabs type="border-card" tab-position="top" @tab-change="tabs.change">
                <ElTabPane label="后台" :lazy="true" :key="tabs.tabPaneKey">
                    <AdminConfig />
                </ElTabPane>
            </ElTabs>
        </ElMain>
    </ElContainer>
</template>