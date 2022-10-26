import { defineStore } from 'pinia'
import md5 from 'js-md5'
import router from '@/router'

export const useAdminStore = defineStore('admin', {
  state: () => {
    return {
      info: {} as { nickname: string, avatar: string, [propName: string]: any }, //用户信息。格式：{nickname: 昵称, avatar: 头像,...}
      menuTree: [] as { title: string, url: string, icon: string, children: { [propName: string]: any }[] }[],   //菜单树。单个菜单格式：{title: 标题, url: 地址, icon: 图标, children: [子集]}
      menuList: [] as { title: string, url: string, icon: string, menuChain: { title: string, url: string, icon: string }[] }[],   //菜单列表。单个菜单格式：{title: 标题, url: 地址, icon: 图标, menuChain: [菜单链]}
      menuTabList: (() => {
        const indexRoute = router.getRoutes().find((item) => {
          return item.path == '/'
        })
        /* router.getRoutes().forEach((item) => {
          item.meta.icon = 'autoicon-ep-lock'
        }) */
        return [{
          title: (<any>indexRoute).meta.title ?? '',
          path: (<any>indexRoute).path,
          icon: (<any>indexRoute).meta.icon ?? '',
          closable: false,
        }]
      })(), //菜单标签列表
    }
  },
  getters: {
    infoIsExist: (state) => {
      return Object.keys(state.info).length ? true : false
    },
    //获取当前菜单的菜单链
    menuChain: (state) => {
      const path = router.currentRoute.value.path
      const menu = state.menuList.find((item) => {
        return item.url == path
      })
      return menu?.menuChain ?? []
    }
  },
  actions: {
    /**
     * 推入菜单标签列表
     * @param menuTab 
     */
    pushMenuTabList(menuTab: { title: string, path: string, icon: string }) {
      let result = this.menuTabList.findIndex((item) => {
        return item.path === menuTab.path
      })
      if (result !== -1) {
        return
      }
      /*--------当前路径在菜单列表中时，以菜单列表中的数据为准 开始--------*/
      const menu = this.menuList.find((item) => {
        return item.url == menuTab.path
      })
      if (menu) {
        menuTab.title = menu.title
        menuTab.icon = menu.icon
      }
      /*--------当前路径在菜单列表中时，以菜单列表中的数据为准 开始--------*/
      this.menuTabList.push({
        closable: true,
        ...menuTab
      })
    },
    /**
     * 关闭自身菜单标签
     * 
     * @param path 
     */
    closeSelfMenuTab(path: string) {
      this.menuTabList = this.menuTabList.filter((item) => {
        return !item.closable || item.path !== path
      })
      const currentPath = router.currentRoute.value.path
      if (path === currentPath) {
        router.push(this.menuTabList[this.menuTabList.length - 1].path)
      }
    },
    /**
     * 关闭其他菜单标签
     * @param {*} path  菜单标签的路由路径
     */
    closeOtherMenuTab(path: string) {
      this.menuTabList = this.menuTabList.filter((item) => {
        return !item.closable || item.path === path
      })
      const currentPath = router.currentRoute.value.path
      if (path !== currentPath) {
        router.push(path)
      }
    },
    /**
     * 关闭左侧菜单标签
     * @param {*} path  菜单标签的路由路径
     */
    closeLeftMenuTab(path: string) {
      const leftIndex = this.menuTabList.findIndex((item) => {
        return item.path === path
      })
      this.menuTabList = this.menuTabList.filter((item, index) => {
        return !item.closable || index >= leftIndex
      })
      const currentPath = router.currentRoute.value.path
      if (path !== currentPath) {
        const currentLeftIndex = this.menuTabList.findIndex((item) => {
          return item.path === currentPath
        })
        if (currentLeftIndex === -1) {
          router.push(path)
        }
      }
    },
    /**
     * 关闭右侧菜单标签
     * @param {*} path  菜单标签的路由路径
     */
    closeRightMenuTab(path: string) {
      const rightIndex = this.menuTabList.findIndex((item) => {
        return item.path === path
      })
      this.menuTabList = this.menuTabList.filter((item, index) => {
        return !item.closable || index <= rightIndex
      })
      const currentPath = router.currentRoute.value.path
      if (path !== currentPath) {
        const currentRightIndex = this.menuTabList.findIndex((item) => {
          return item.path === currentPath
        })
        if (currentRightIndex === -1) {
          router.push(path)
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
      //router.push('/')
      router.push(this.menuTabList[this.menuTabList.length - 1].path)
    },
    /**
     * 登录
     * @param {*} account   账号
     * @param {*} password  密码
     * @returns 
     */
    async login(account: string, password: string) {
      let res = await request('login.encryptStr', {
        account: account
      }, false)
      res = await request('login.login', {
        account: account,
        password: md5(md5(password) + res.data.encryptStr)
      }, false)
      this.$reset() //重置状态（可有效清理上一个登录用户的脏数据）
      //不用清空缓存组件，登录后切换页面过程中，layout布局组件已经重新生成，其内部所有缓存组件已经重置
      //useKeepAliveStore().$reset()

      setAccessToken(res.data.token)
      this.setInfo(); //设置用户信息（可选，路由前置守卫有执行，此处执行，路由可减少一次跳转）
      this.setMenuTree()   //设置左侧菜单树（可选，路由前置守卫有执行，此处执行，路由可减少一次跳转）
    },
    /**
     * 设置登录用户信息
     */
    async setInfo() {
      const res = await request('login.info', {}, false)
      this.info = res.data.info
    },
    /**
     * 设置左侧菜单（包含注册动态路由）
     */
    async setMenuTree() {
      const res = await request('login.menuTree', {}, false)
      /**--------注册动态路由 开始--------**/
      const handleMenuTree = (menuTree: any, menuChain: any = []) => {
        const menuTreeTmp: any = []
        for (let i = 0; i < menuTree.length; i++) {
          menuTreeTmp[i] = {
            title: menuTree[i].title,
            url: menuTree[i].url,
            icon: menuTree[i].icon,
            children: [],
          }
          if (menuTree[i].children.length) {
            menuChain.push({
              title: menuTree[i].title,
              url: menuTree[i].url,
              icon: menuTree[i].icon,
            })
            menuTreeTmp[i].children = handleMenuTree(menuTree[i].children, [...menuChain])
            menuChain.pop()
          } else {
            const menu = {
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
