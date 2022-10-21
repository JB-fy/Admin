import { defineStore } from 'pinia'
import md5 from 'js-md5'
import { getInfo, getEncryptStr, getMenuTree, login } from '@/api/login'
import router from '@/router'

export const useUserStore = defineStore('user', {
  state: () => {
    return {
      info: {} as { nickname: string, avatar: string, [propName: string]: any }, //用户信息。格式：{nickname: 昵称, avatar: 头像,...}
      menuTree: [] as { title: string, url: string, icon: string, children: {}[] }[],   //左侧菜单树。单个菜单格式：{title: 标题, url: 地址, icon: 图标, children: [子集]}
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
      if (result === -1) {
        this.menuTabList.push({
          closable: true,
          ...menuTab
        })
      }
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
      const currentPath = getCurrentPath()
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
      const currentPath = getCurrentPath()
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
      const currentPath = getCurrentPath()
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
      const currentPath = getCurrentPath()
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
      //this.$reset() //重置状态（可有效清理上一个登录用户的脏数据）
      try {
        let res = await getEncryptStr({
          account: account
        })
        res = await login({
          account: account,
          password: md5(md5(password) + res.data.encryptStr)
        })
        this.$reset() //重置状态（可有效清理上一个登录用户的脏数据）
        //不用清空缓存组件，登录后切换页面过程中，layout布局组件已经重新生成，其内部所有缓存组件已经重置
        //useKeepAliveStore().$reset()

        setAccessToken(res.data.token)
        this.setInfo(); //设置用户信息（可选，路由前置守卫有执行，此处执行，路由可减少一次跳转）
        this.setMenuTree()   //设置左侧菜单树（可选，路由前置守卫有执行，此处执行，路由可减少一次跳转）
        return true
      } catch (err) {
        throw err 
        //await errorHandle(err)
        return false
      }
    },
    /**
     * 设置登录用户信息
     */
    async setInfo() {
      try {
        const res = await getInfo()
        this.info = res.data.info
        return true
      } catch (err) {
        throw err 
        return false
      }
    },
    /**
     * 设置左侧菜单（包含注册动态路由）
     */
    async setMenuTree() {
      try {
        const res = await getMenuTree()
        /**--------注册动态路由 开始--------**/
        const handleMenuTree = (menuTree: any, pMenuList: any = []) => {
          const menuTreeTmp: any = []
          let tmpExtendData: any = {};
          for (let i = 0; i < menuTree.length; i++) {
            tmpExtendData = JSON.parse(menuTree[i].extendData);
            menuTreeTmp[i] = {
              title: tmpExtendData.title,
              url: tmpExtendData.url,
              icon: tmpExtendData.icon,
              children: [],
            }
            if (menuTree[i].children.length) {
              pMenuList.push({
                title: tmpExtendData.title,
                url: tmpExtendData.url,
                icon: tmpExtendData.icon,
              })
              menuTreeTmp[i].children = handleMenuTree(menuTree[i].children, Object.assign({}, pMenuList))
              pMenuList.pop()
            }/*  else {
              router.addRoute(layoutName, {
                path: menuTree[i].menuUrl,
                name: menuTree[i].menuUrl,  //命名路由，用户退出登录用于删除路由。要保证唯一，故直接用menuUrl即可
                //component: () => import('@/views' + menuTree[i].menuUrl),
                component: async () => {
                  //let component = await import('@/views' + menuTree[i].menuUrl + '.vue'),
                  let component = await import('@/views' + menuTree[i].menuUrl)
                  component.default.name = menuTree[i].menuUrl    //动态设置页面组件名称，方便清理缓存
                  return component
                },
                meta: {
                  title: menuTree[i].menuName,
                  icon: menuTree[i].menuIcon,
                  pMenuList: Object.assign({}, pMenuList) //面包屑需要
                }
              })
            } */
          }
          return menuTreeTmp
        }
        this.menuTree = handleMenuTree(res.data.tree)
        /**--------注册动态路由 结束--------**/
        return true
      } catch (err) {
        throw err 
        return false
      }
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
