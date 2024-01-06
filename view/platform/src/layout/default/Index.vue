<script setup lang="tsx">
import AppContainer from './AppContainer.vue'
import LeftMenu from './LeftMenu.vue'
import RightHeader from './RightHeader.vue'
import RightFooter from './RightFooter.vue'

const { t } = useI18n()
const settingStore = useSettingStore()
</script>

<template>
    <el-container id="layout-container">
        <el-aside :class="{ 'is-fold': settingStore.leftMenuFold }">
            <left-menu />
        </el-aside>
        <el-container>
            <el-header>
                <right-header />
            </el-header>

            <el-main class="main-container">
                <app-container />
                <el-backtop target=".main-container" :right="16" :bottom="60">
                    <span style="font-size: 14px">{{ t('common.upTop') }}</span>
                    <!-- <autoicon-ep-top /> -->
                </el-backtop>
            </el-main>

            <el-footer>
                <right-footer />
            </el-footer>
        </el-container>
    </el-container>
</template>

<style scoped>
#layout-container {
    height: 100vh;
}

#layout-container .el-aside {
    color: #fff;
    background-color: #545c64;
    transition-duration: 0.5s;
    width: 200px;
}

#layout-container .el-aside.is-fold {
    width: 64px;
}

#layout-container .el-header {
    height: 100px;
    padding: 0;
    box-shadow: var(--el-box-shadow-lighter);
    z-index: 1;
    /* background-color: var(--el-color-primary-light-7);
    color: var(--el-text-color-primary); */
}

#layout-container .el-main {
    background-color: #f2f3f5;
    /* padding: 0; */
}

#layout-container .el-main::-webkit-scrollbar {
    width: 6px;
}

#layout-container .el-main::-webkit-scrollbar-thumb {
    border-radius: 4px;
    background-color: #dcdfe6;
}

#layout-container .el-main::-webkit-scrollbar-thumb:hover {
    background-color: #c0c4cc;
}

#layout-container .el-footer {
    height: 50px;
    text-align: center;
    box-shadow: var(--el-box-shadow-lighter);
}
</style>
