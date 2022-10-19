<script setup lang="ts">
import LeftMenuItem from './LeftMenuItem.vue';   //做成组件才能实现无限递归（组件内部无限递归自身）
import { useSettingStore } from '@/stores/setting';
import { useUserStore } from '@/stores/user';

const settingStore = useSettingStore()
const userStore = useUserStore()

/**--------bug处理（组件ElMenu设置背景色后，鼠标移动到含有子菜单的菜单上背景色不变） 开始--------**/
const subMenuTitle = async () => {
    await nextTick()
    let subMenuTitleList = document.getElementsByClassName('el-sub-menu__title')
    for (let i = 0; i < subMenuTitleList.length; i++) {
        (<any>subMenuTitleList[i]).style.removeProperty('background-color')
    }
}
const router = useRouter()
const route = useRoute()
const menuSelect = (path: string) => {
    subMenuTitle()    //菜单激活也会导致bug。:router="true"会使这里无效，虽有执行，但不能去除css属性。可能路由跳转后才追加css属性
    if (path.indexOf('http') === 0) {
        window.open(path, '_blank')
    } else {
        router.push(path)
    }
}
const menuOpen = () => {
    subMenuTitle()
}
const menuClose = () => {
    subMenuTitle()
}
/**--------bug处理（组件ElMenu设置背景色后，鼠标移动到含有子菜单的菜单上背景色不变） 结束--------**/
</script>

<template>
    <ElScrollbar>
        <ElMenu :default-active="route.path" :collapse="settingStore.leftMenuFold" :router="false" :unique-opened="true"
            background-color="#545c64" text-color="#fff" active-text-color="#ffd04b" @select="menuSelect"
            @open="menuOpen" @close="menuClose">
            <LeftMenuItem :tree="userStore.menuTree" />
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