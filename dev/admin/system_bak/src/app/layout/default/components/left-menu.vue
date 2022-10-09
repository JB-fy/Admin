<template>
    <el-scrollbar>
        <el-menu :default-active="$route.path" :collapse="$store.state.setting.leftMenuFold" :router="false"
            :unique-opened="true" background-color="#545c64" text-color="#fff" active-text-color="#ffd04b"
            @select="menuSelect" @open="menuOpen" @close="menuClose">
            <left-menu-item :tree="$store.state.user.leftMenuTree" />
        </el-menu>
    </el-scrollbar>
</template>

<script>
import leftMenuItem from './left-menu-item'

export default {
    components: {
        'left-menu-item': leftMenuItem  //做成组件才能实现无限递归（组件内部无限递归自身）
    },
    setup: () => {
        /**--------bug处理（组件el-menu设置背景色后，鼠标移动到含有子菜单的菜单上背景色不变） 开始--------**/
        const subMenuTitle = async () => {
            await nextTick()
            let subMenuTitleList = document.getElementsByClassName('el-sub-menu__title')
            for (let i = 0; i < subMenuTitleList.length; i++) {
                subMenuTitleList[i].style.removeProperty('background-color')
            }
        }
        const router = useRouter()
        const menuSelect = (path) => {
            subMenuTitle()    //菜单激活也会导致bug。:router="true"会使这里无效，虽有执行，但不能去除css属性。可能路由跳转后才追加css属性
            router.push(path)
        }
        const menuOpen = () => {
            subMenuTitle()
        }
        const menuClose = () => {
            subMenuTitle()
        }
        /**--------bug处理（组件el-menu设置背景色后，鼠标移动到含有子菜单的菜单上背景色不变） 结束--------**/
        return {
            menuSelect,
            menuOpen,
            menuClose
        }
    }
}
</script>

<style scoped>
.el-menu {
    border-right: none;
}

/* .el-menu ::v-deep .el-sub-menu__title {
    background-color: '' !important;
} */
</style>