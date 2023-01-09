import { defineStore } from 'pinia'
import md5 from 'js-md5'
import router from '@/router'

export const useAdminStore = defineStore('admin', {
  state: () => {
    return {
      info: {} as { nickname: string, avatar: string, [propName: string]: any }, //用户信息。格式：{nickname: 昵称, avatar: 头像,...}
      menuTree: [] as { menuName: string, title: { [propName: string]: any }, url: string, icon: string, children: { [propName: string]: any }[] }[],   //菜单树。单个菜单格式：{menuName: 菜单名称（title不存在时默认该字段作标题）, title: {"i18n语言标识":"语言标题",...}, url: 地址, icon: 图标, children: [子集]}
      menuList: [] as { menuName: string, title: { [propName: string]: any }, url: string, icon: string, menuChain: { [propName: string]: any }[] }[],   //菜单列表。单个菜单格式：{menuName: 菜单名称（title不存在时默认该字段作标题）, title: {"i18n语言标识":"语言标题",...}, url: 地址, icon: 图标, menuChain: [菜单链（包含自身）]}
      menuTabList: [] as { keepAlive: boolean, componentName: string, url: string, menuName: string, title: { [propName: string]: string }, icon: string, closable: boolean }[], //菜单标签列表
    }
  },
  getters: {
    infoIsExist: (state) => {
      return Object.keys(state.info).length ? true : false
    },
    //获取当前菜单的菜单链
    getCurrentMenuChain: (state) => {
      const menu = state.menuList.find((item) => {
        return item.url == router.currentRoute.value.fullPath
      })
      //菜单中没有，就直接返回路由中的数据。例如：个人中心页面
      if (!menu?.menuChain) {
        return [{
          title: useLanguageStore().getMenuTitle(router.currentRoute.value.meta?.menu),
          url: router.currentRoute.value.fullPath,
          icon: (router.currentRoute.value.meta?.menu as any)?.icon,
        }]
      }
      return menu.menuChain.map((item) => {
        return {
          title: useLanguageStore().getMenuTitle(item),
          url: item.url,
          icon: item.icon,
        }
      })
    },
    //获取菜单标签列表
    getMenuTabList: (state) => {
      const routeOfIndex = (<any>router).getRoutes().find((item: any) => {
        return item.path == '/'
      })
      let menuTabList
      if (routeOfIndex) {
        const menuTabOfIndex = {
          componentName: routeOfIndex.meta.componentName,
          url: routeOfIndex.path,
          closable: false,  //首页的菜单标签不能关闭
          menuName: routeOfIndex.meta?.menu?.menuName,
          title: routeOfIndex.meta?.menu?.title,
          icon: routeOfIndex.meta?.menu?.icon,
        }
        const menuOfIndex = state.menuList.find((item) => {
          return item.url == routeOfIndex.path
        })
        if (menuOfIndex) {
          menuTabOfIndex.menuName = menuOfIndex.menuName
          menuTabOfIndex.title = menuOfIndex.title
          menuTabOfIndex.icon = menuOfIndex.icon
        }
        menuTabList = [{ ...menuTabOfIndex }, ...state.menuTabList]
      } else {
        menuTabList = [...state.menuTabList]
      }
      return menuTabList.map((item) => {
        return {
          componentName: item.componentName,
          url: item.url,
          title: useLanguageStore().getMenuTitle(item),
          icon: item.icon,
          closable: item.closable,
        }
      })
    },
  },
  actions: {
    /**
     * 推入菜单标签列表
     * @param menuTab 
     */
    pushMenuTabList(menuTab: { keepAlive: boolean, componentName: string, url: string, menuName: string, title: { [propName: string]: string }, icon: string }) {
      if (menuTab.url == '/') {
        return
      }
      let result = this.menuTabList.findIndex((item) => {
        return item.url === menuTab.url
      })
      if (result !== -1) {
        return
      }
      /*--------当前路由在菜单列表中时，以菜单列表中的数据为准 开始--------*/
      const menu = this.menuList.find((item) => {
        return item.url == menuTab.url
      })
      if (menu) {
        menuTab.menuName = menu.menuName
        menuTab.title = menu.title
        menuTab.icon = menu.icon
      }
      /*--------当前路由在菜单列表中时，以菜单列表中的数据为准 开始--------*/
      this.menuTabList.push({
        ...menuTab,
        closable: true,
      })
    },
    /**
     * 关闭自身菜单标签
     * @param fullPath 
     */
    closeSelfMenuTab(fullPath: string) {
      this.menuTabList = this.menuTabList.filter((item) => {
        return !item.closable || item.url !== fullPath
      })
      if (fullPath === router.currentRoute.value.fullPath) {
        router.push(this.menuTabList?.[this.menuTabList.length - 1]?.url ?? '/')
      }
    },
    /**
     * 关闭其他菜单标签
     * @param {*} fullPath  菜单标签的路由路径
     */
    closeOtherMenuTab(fullPath: string) {
      this.menuTabList = this.menuTabList.filter((item) => {
        return !item.closable || item.url === fullPath
      })
      if (fullPath !== router.currentRoute.value.fullPath) {
        router.push(fullPath)
      }
    },
    /**
     * 关闭左侧菜单标签
     * @param {*} fullPath  菜单标签的路由路径
     */
    closeLeftMenuTab(fullPath: string) {
      const leftIndex = this.menuTabList.findIndex((item) => {
        return item.url === fullPath
      })
      this.menuTabList = this.menuTabList.filter((item, index) => {
        return !item.closable || index >= leftIndex
      })
      if (fullPath !== router.currentRoute.value.fullPath) {
        const currentLeftIndex = this.menuTabList.findIndex((item) => {
          return item.url === router.currentRoute.value.fullPath
        })
        if (currentLeftIndex === -1) {
          router.push(fullPath)
        }
      }
    },
    /**
     * 关闭右侧菜单标签
     * @param {*} fullPath  菜单标签的路由路径
     */
    closeRightMenuTab(fullPath: string) {
      const rightIndex = this.menuTabList.findIndex((item) => {
        return item.url === fullPath
      })
      this.menuTabList = this.menuTabList.filter((item, index) => {
        return !item.closable || index <= rightIndex
      })
      if (fullPath !== router.currentRoute.value.fullPath) {
        const currentRightIndex = this.menuTabList.findIndex((item) => {
          return item.url === router.currentRoute.value.fullPath
        })
        if (currentRightIndex === -1) {
          router.push(fullPath)
        }
      }
    },
    /**
     * 关闭全部菜单标签
     */
    closeAllMenuTab() {
      this.menuTabList = this.menuTabList.filter((item) => {
        return !item.closable
      })
      router.push(this.menuTabList?.[this.menuTabList.length - 1]?.url ?? '/')
    },
    /**
     * 登录
     * @param {*} account   账号
     * @param {*} password  密码
     * @returns 
     */
    async login(account: string, password: string) {
      let res = await request('login/encryptStr', {
        account: account
      })
      res = await request('login', {
        account: account,
        password: md5(md5(password) + res.data.encryptStr)
      })
      this.$reset() //重置状态（可有效清理上一个登录用户的脏数据）
      //不用清空缓存组件，登录后切换页面过程中，layout布局组件已经重新生成，其内部所有缓存组件已经重置
      //useKeepAliveStore().$reset()
      setAccessToken(res.data.token)
    },
    /**
     * 设置登录用户信息
     */
    async setInfo() {
      const res = await request('login/info', {})
      this.info = res.data.info
    },
    /**
     * 设置左侧菜单（包含更新路由meta数据）
     */
    async setMenuTree() {
      const res = await request('login/menuTree', {})
      /**--------注册动态路由 开始--------**/
      const handleMenuTree = (menuTree: any, menuChain: any = []) => {
        const menuTreeTmp: any = []
        for (let i = 0; i < menuTree.length; i++) {
          menuTreeTmp[i] = {
            menuName: menuTree[i].menuName,
            title: menuTree[i].title,
            url: menuTree[i].url,
            icon: menuTree[i].icon,
            children: [],
          }
          if (menuTree[i].children.length) {
            menuChain.push({
              menuName: menuTree[i].menuName,
              title: menuTree[i].title,
              url: menuTree[i].url,
              icon: menuTree[i].icon,
            })
            menuTreeTmp[i].children = handleMenuTree(menuTree[i].children, [...menuChain])
            menuChain.pop()
          } else {
            const menu = {
              menuName: menuTree[i].menuName,
              title: menuTree[i].title,
              url: menuTree[i].url,
              icon: menuTree[i].icon
            }
            this.menuList.push({
              ...menu,
              menuChain: [...menuChain, menu]
            })
          }
        }
        return menuTreeTmp
      }
      this.menuTree = handleMenuTree(res.data.tree)
      /**--------注册动态路由 结束--------**/
    },
    /**
     * 退出登录
     * @param {*} toPath  跳转路径
     */
    logout(toPath: string = '/login') {
      removeAccessToken()
      //router.push(toPath)
      if (toPath === '/login') {
        router.push(toPath)
      } else {
        router.push('/login?redirect=' + toPath)
      }
    }
  },
})
