import { defineStore } from 'pinia'
import md5 from 'js-md5'
import router from '@/router'

export const useAdminStore = defineStore('admin', {
  state: () => {
    return {
      info: {} as { account: string, phone: string, nickname: string, avatar: string, [propName: string]: any }, //用户信息
      menuTree: [] as { menuName: string, title: { [propName: string]: any }, url: string, icon: string, children: { [propName: string]: any }[] }[],   //菜单树。单个菜单格式：{menuName: 菜单名称（title不存在时默认该字段作标题）, title: {"i18n语言标识":"语言标题",...}, url: 地址, icon: 图标, children: [子集]}
      menuList: [] as { menuName: string, title: { [propName: string]: any }, url: string, icon: string, menuChain: { [propName: string]: any }[] }[],   //菜单列表。单个菜单格式：{menuName: 菜单名称（title不存在时默认该字段作标题）, title: {"i18n语言标识":"语言标题",...}, url: 地址, icon: 图标, menuChain: [菜单链（包含自身）]}
      menuTabList: [] as { keepAlive: boolean, componentName: string, url: string, menuName: string, title: { [propName: string]: string }, icon: string, closable: boolean }[], //菜单标签列表
      //开发工具菜单。只在开发模式显示（即import.meta.env.DEV为true）
      menuTreeOfDev: {
        "menuName": "开发工具",
        "title": {
          "en": "Dev Tool",
          "zh-cn": "开发工具"
        },
        "url": "",
        "icon": "AutoiconEpHelpFilled",
        "children": [
          /* {
            "menuName": "说明文档",
            "title": {
              "en": "Document",
              "zh-cn": "说明文档"
            },
            //"url": "https://www.baidu.com/",  //新窗口打开
            "url": "/thirdSite?url=https://www.baidu.com/", //标签页打开
            "icon": "AutoiconEpDocument",
            "children": []
          }, */
          {
            "menuName": "Hyperf",
            "title": {
              "en": "Hyperf",
              "zh-cn": "Hyperf"
            },
            "url": "https://www.hyperf.io/",
            "icon": "AutoiconEpChromeFilled",
            "children": []
          },
          {
            "menuName": "ElementPlus",
            "title": {
              "en": "Element Plus",
              "zh-cn": "Element Plus"
            },
            "url": "https://element-plus.gitee.io/zh-CN/",
            "icon": "AutoiconEpElementPlus",
            "children": []
          },
          {
            "menuName": "Vant4",
            "title": {
              "en": "Vant 4",
              "zh-cn": "Vant 4"
            },
            "url": "https://vant-contrib.gitee.io/vant/#/zh-CN",
            "icon": "Vant-wechat-moments",
            "children": []
          },
          {
            "menuName": "Vue",
            "title": {
              "en": "Vue",
              "zh-cn": "Vue"
            },
            "url": "https://cn.vuejs.org/api/",
            "icon": "AutoiconEpChromeFilled",
            "children": []
          }
        ]
      } as { menuName: string, title: { [propName: string]: any }, url: string, icon: string, children: { [propName: string]: any }[] },
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
      const menuTabList = state.menuTabList.map((item) => {
        return {
          componentName: item.componentName,
          url: item.url,
          title: useLanguageStore().getMenuTitle(item),
          icon: item.icon,
          closable: item.closable,
        }
      })
      /*--------增加首页的菜单标签并置顶 开始--------*/
      const routeOfIndex = (<any>router).getRoutes().find((item: any) => {
        return item.path == '/'
      })
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
        menuTabOfIndex.title = useLanguageStore().getMenuTitle(menuTabOfIndex)
        menuTabList.unshift({ ...menuTabOfIndex })
      }
      /*--------增加首页的菜单标签并置顶 结束--------*/
      return menuTabList
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
     * 设置左侧菜单树（包含更新路由meta数据）
     */
    async setMenuTree() {
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
            //设置菜单列表
            this.menuList.push({
              ...menu,
              menuChain: [...menuChain, menu]
            })
          }
        }
        return menuTreeTmp
      }
      const res = await request('login/menuTree')
      const tree = import.meta.env.DEV ? [...res.data.tree, this.menuTreeOfDev] : res.data.tree
      this.menuTree = handleMenuTree(tree)
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
