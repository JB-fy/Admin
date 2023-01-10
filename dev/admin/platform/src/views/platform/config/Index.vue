<script setup lang="ts">
import AdminConfig from './AdminConfig.vue'

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
                <ElTabPane label="代理" :lazy="true">
                    <AdminConfig />
                </ElTabPane>
            </ElTabs>
        </ElMain>
    </ElContainer>
</template>