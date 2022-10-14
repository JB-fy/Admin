import { defineStore } from 'pinia'
import md5 from 'js-md5'
import { getInfo, getEncryptStr, getMenuTree, login } from '@/app/api/login'
import router from '@/router'

export const useUserStore = defineStore('user', {
  state: () => {
    return {
      info: {}, //用户信息。格式：{nickname: 昵称, avatar: 头像, rawInfo: 原始信息（后台传过来的原始数据）}
      leftMenuTree: [],   //左侧菜单树。单个菜单格式：{title: 标题, path: 路径, icon: 图标, children: [子集]}
      menuTabList: [], //菜单标签列表（打开标签即是允许缓存的组件）
    }
  },
  getters: {
    infoIsExist: (state) => {
      return Object.keys(state.info).length ? true : false
    },
    menuTabListLength: (state) => {
      return state.menuTabList.length
    }
  },
  actions: {
    /**
     * 推入菜单标签列表
     * @param {*} routeTo  将要打开的路由
     */
    pushMenuTabList(routeTo) {
      let result = this.menuTabList.findIndex((item) => {
        return item.path === routeTo.path
      })
      if (result === -1) {
        this.menuTabList.push({
          title: routeTo.meta.title,
          path: routeTo.path,
          icon: routeTo.meta.icon,
          closable: routeTo.closable === false ? false : true,
        })
      }
    },
    /**
     * 关闭自身菜单标签
     * @param {*} path  菜单标签的路由路径
     */
    closeSelfMenuTab(path) {
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
    closeOtherMenuTab(path) {
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
    closeLeftMenuTab(path) {
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
    closeRightMenuTab(path) {
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
        setAccessToken(res.data.token)
        /**--------初始化数据（可有效清理上一个登录用户的脏数据） 开始--------**/
        //在logout退出登录操作中也可以清理，但在登录操作这里处理，应变能力更好。不用考虑有多少种情况需及时清理脏数据，如：accessToken失效、切换用户等
        //this.info = {}; //清空用户信息
        this.setInfo(); //设置用户信息（可选，路由前置守卫有执行，此处执行，路由可减少一次跳转）

        //this.leftMenuTree = []  //清空用户左侧菜单
        this.setLeftMenuTree()   //设置左侧菜单树（可选，路由前置守卫有执行，此处执行，路由可减少一次跳转）

        this.menuTabList = [] //清空菜单标签列表
        //不用清空缓存组件，登录后切换页面过程中，layout布局组件已经重新生成，其内部所有缓存组件已经重置
        /**--------初始化数据（可有效清理上一个登录用户的脏数据） 结束--------**/
        return true
      } catch (err) {
        await errorHandle(err)
        return false
      }
    },
    /**
     * 设置登录用户信息
     */
    async setInfo() {
      try {
        const res = await getInfo()
        this.info = {
          nickname: res.data.info.nickname ? res.data.info.nickname : res.data.info.account,
          avatar: res.data.info.avatar,
          rawInfo: res.data.info,
        }
        return true
      } catch (err) {
        await errorHandle(err)
        return false
      }
    },
    /**
     * 设置左侧菜单（包含注册动态路由）
     */
    async setLeftMenuTree() {
      try {
        const res = await getMenuTree()
        /**--------注册动态路由 开始--------**/
        const handleMenuTree = (menuTree, pMenuList = []) => {
          const leftMenuTree = []
          let tmpExtendData = {};
          for (let i = 0; i < menuTree.length; i++) {
            tmpExtendData = JSON.parse(menuTree[i].extendData);
            leftMenuTree[i] = {
              title: tmpExtendData.title,
              path: tmpExtendData.url,
              icon: tmpExtendData.icon,
              children: [],
            }
            if (menuTree[i].children.length) {
              pMenuList.push({
                title: tmpExtendData.title,
                path: tmpExtendData.url,
                icon: tmpExtendData.icon,
              })
              leftMenuTree[i].children = handleMenuTree(menuTree[i].children, Object.assign({}, pMenuList))
              pMenuList.pop()
            }
          }
          return leftMenuTree
        }
        const leftMenuTree = handleMenuTree(res.data.tree)
        this.leftMenuTree = leftMenuTree
        /**--------注册动态路由 结束--------**/
        return true
      } catch (err) {
        await errorHandle(err)
        return false
      }
    },
    /**
     * 退出登录
     * @param {*} toPath  跳转路径
     */
    async logout(toPath = '/login') {
      await removeAccessToken()
      await router.push(toPath)
      /* if (toPath === '/login') {
        await router.push(toPath)
      } else {
        await router.push('/login?redirect=' + toPath)
      } */
    }
  },
})
