import { defineStore } from 'pinia'
import router from '@/router'

export const useKeepAliveStore = defineStore('keepAlive', {
  state: () => {
    return {
      /* appContainerInclude: (() => {
        const include: string[] = []
        router.getRoutes().forEach((item) => {
          if (item.meta.keepAlive) {
            //include.push(item.components.default.name)
            include.push(item.path)
          }
        })
        return include
      })(), */
      appContainerExclude: [] as string[], //不允许缓存的路由路径列表，这里主要用于实现缓存刷新（动态设置页面组件名称name时，用路径命名，故这里面填写路径）
      appContainerMax: 10 as number   //缓存组件最大数量
    }
  },
  getters: {
    appContainerInclude: (state): string[] => {
      const include: string[] = []
      router.getRoutes().forEach((item) => {
        //菜单允许缓存，且打开菜单才做缓存。打开菜单才做缓存是为在菜单关闭时实现自动清理缓存，否则关闭菜单还得清理缓存
        if (item.meta.keepAlive && useAdminStore().menuTabList.findIndex((menuTab) => {
          return menuTab.path === item.path
        }) !== -1) {
          //include.push(item.components.default.name)
          include.push(item.path)
        }
      })
      return include
    },
  },
  actions: {
    /**
     * 删除不允许缓存的组件
     * @param {*} path  路径
     */
    removeAppContainerExclude(path: string) {
      this.appContainerExclude = this.appContainerExclude.filter((item) => {
        return item !== path
      })
    },
    /**
     * 刷新菜单标签
     *      实现流程：
     *          1：AppContainer.vue文件内component标签加上判断是否允许缓存，允许才显示界面（v-if="keepAliveStore.appContainerExclude.indexOf(route.path) === -1"）
     *          2：设置路由不允许缓存，不显示页面
     *          3：打开路由，路由后置守卫afterEach中重新设置成允许缓存，显示页面
     * @param {*} path  菜单标签的路由路径
     */
    refreshMenuTab(path: string) {
      this.appContainerExclude.push(path)
      const currentPath = router.currentRoute.value.path
      if (path === currentPath) {
        router.push(path)
      }
    },
  }
})
