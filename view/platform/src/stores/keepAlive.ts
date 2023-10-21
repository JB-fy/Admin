import { defineStore } from 'pinia'
import router from '@/router'

/**
 * 路由定义时，组件name设置为path，故appContainerInclude和appContainerExclude内只有存放path才能实现组件缓存和页面刷新
 */
export const useKeepAliveStore = defineStore('keepAlive', {
  state: () => {
    return {
      appContainerExclude: [] as string[], //用于页面刷新
      appContainerMax: 20 as number,  //缓存组件最大数量
    }
  },
  getters: {
    //用于组件缓存。打开的菜单标签才做缓存，好处：菜单标签关闭后，对应的组件缓存就会被删除，重新打开时就不需要刷新页面
    appContainerInclude: (state): string[] => {
      const include: string[] = []
      useAdminStore().menuTabList.forEach((menuTab) => {
        if (menuTab.keepAlive) {
          include.push(menuTab.componentName)
        }
      })
      return include
    },
  },
  actions: {
    /**
     * 删除不允许缓存的组件
     * @param {*} componentName
     */
    removeAppContainerExclude(componentName: string) {
      this.appContainerExclude = this.appContainerExclude.filter((item) => {
        return item !== componentName
      })
    },
    /**
     * 刷新菜单标签
     *      实现流程：
     *          1：AppContainer.vue文件内component标签加上判断是否允许缓存，允许才显示界面（v-if="keepAliveStore.appContainerExclude.indexOf(<any>route.meta.componentName) === -1"）
     *          2：设置路由不允许缓存，不显示页面
     *          3：打开路由，路由后置守卫afterEach中重新设置成允许缓存，显示页面
     * @param {*} componentName
     */
    refreshMenuTab(componentName: string) {
      this.appContainerExclude.push(componentName)
      if (componentName === router.currentRoute.value.meta.componentName) {
        router.push(router.currentRoute.value.fullPath)
      }
    },
  }
})
