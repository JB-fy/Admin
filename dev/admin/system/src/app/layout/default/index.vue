<template>
    <el-container id="layout-container">
        <el-aside :class="{ 'is-fold': $store.state.setting.leftMenuFold }">
            <left-menu />
        </el-aside>
        <el-container>
            <el-header>
                <right-header />
            </el-header>

            <el-main class="main-container">
                <app-container />
                <el-backtop target=".main-container" :right="16" :bottom="60">
                    <span style="font-size: 14px;">置顶</span>
                    <!-- <autoicon-ep-top /> -->
                </el-backtop>
            </el-main>

            <el-footer>
                <right-footer />
            </el-footer>
        </el-container>
    </el-container>
</template>

<script>
export default {
    components: {
        ...getImportComponents(import.meta.globEager('@/app/layout/default/components/*.vue'))
    }
}
</script>


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