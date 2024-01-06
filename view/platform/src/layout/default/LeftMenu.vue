<script setup lang="tsx">
import LeftMenuItem from './LeftMenuItem.vue' //做成组件才能实现无限递归（组件内部无限递归自身）

const settingStore = useSettingStore()
const adminStore = useAdminStore()

const router = useRouter()
const route = useRoute()
const menuSelect = (fullPath: string) => {
    if (fullPath.indexOf('http') === 0) {
        window.open(fullPath, '_blank')
    } else {
        router.push(fullPath)
    }
}
</script>

<template>
    <ElScrollbar>
        <ElMenu :default-active="route.fullPath" :collapse="settingStore.leftMenuFold" :router="false" :unique-opened="true" background-color="#545c64" text-color="#fff" active-text-color="#ffd04b" @select="menuSelect">
            <LeftMenuItem :tree="adminStore.menuTree" />
        </ElMenu>
    </ElScrollbar>
</template>

<style scoped>
.el-menu {
    border-right: none;
}

/* .el-menu :deep(.el-sub-menu__title) {
    background-color: '' !important;
} */
</style>
