<!--直接批量导入组件 开始-->
<script lang="ts">
const components: { [propName: string]: any } = batchImport(import.meta.globEager('./tabPane/*.vue'))
export default {
    components: components,
}
</script>
<!--直接批量导入组件 结束-->
<script setup lang="ts">
//import Demo from './tabPane/Demo.vue'
// const Demo  = defineAsyncComponent(() => import('./tabPane/Demo.vue'))    //好处：该组件会被打包成单独一个文件
// console.log(Demo)

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
            <!-- <ElTabs type="border-card" tab-position="top" v-model="tabs.activeTabName" @tab-change="tabs.change"> -->
            <ElTabs type="border-card" tab-position="top" @tab-change="tabs.change">
                <ElTabPane label="后台" :lazy="true" :key="tabs.tabPaneKey">
                    <Demo />
                </ElTabPane>
            </ElTabs>
        </ElMain>
    </ElContainer>
</template>