import { getInfo, getLoginToken, getMenuTree, login } from '@/app/api/login'
import md5 from 'js-md5'
import router from '@/app/basic/router'

export default {
    namespaced: true,
    state: {
        info: {}, //用户信息。格式：{nickname: 昵称, avatar: 头像, rawInfo: 原始信息（后台传过来的原始数据）}
        leftMenuTree: [],   //左侧菜单树。单个菜单格式：{title: 标题, path: 路径, icon: 图标, children: [子集]}
        constRoutePathList: [], //固定路由路径列表，即注册动态路由前存在的固定路由。用于删除动态路由
        menuTabList: [], //菜单标签列表（打开标签即是允许缓存的组件）
        cacheRoute: {
            //include: [],    //可以通过getters/cacheRouteInclude计算得出，且菜单标签变动会自动删除对应缓存组件
            exclude: [], //不允许缓存的路由路径列表，这里主要用于实现缓存刷新（动态设置页面组件名称name时，用路径命名，故这里面填写路径）
            max: config('app.router.cacheRoute.max')   //缓存组件最大数量
        }
    },
    getters: {
        infoIsExist: (state, getters) => {
           return Object.keys(state.info).length ? true : false
        },
        menuTabListLength: (state) => {
            return state.menuTabList.length
        },
        cacheRouteInclude: (state) => {
            const include = []
            const cacheRouteConstExclude = config('app.router.cacheRoute.constExclude')
            for (let i = 0; i < state.menuTabList.length; i++) {
                if (cacheRouteConstExclude.indexOf(state.menuTabList[i].path) === -1) {
                    include.push(state.menuTabList[i].path)
                }
            }
            return include
        }
    },
    mutations: {
        /**
         * 设置登录用户信息
         * @param {*} info 用户信息
         */
        setInfo: (state, info) => {
            state.info = info
        },
        /**
         * 设置左侧菜单树
         * @param {*} leftMenuTree 左侧菜单树
         */
        setLeftMenuTree: (state, leftMenuTree) => {
            state.leftMenuTree = leftMenuTree
        },
        /**
         * 设置固定路由路径列表（必须在注册动态路由前调用，建议直接在main.js中调用）
         */
        setConstRoutePathList: (state) => {
            const constRoutePathList = []
            const constRoutes = router.getRoutes()
            for (let i = 0; i < constRoutes.length; i++) {
                constRoutePathList.push(constRoutes[i].path)
            }
            state.constRoutePathList = constRoutePathList
        },
        /**
         * 重置路由
         */
        resetRoute: (state) => {
            const allRoutes = router.getRoutes()
            for (let i = 0; i < allRoutes.length; i++) {
                if (state.constRoutePathList.indexOf(allRoutes[i].path) === -1) {
                    router.removeRoute(allRoutes[i].path)
                }
            }
        },
        /**
         * 删除不允许缓存的路由组件
         * @param {*} path  路由路径
         */
        removeCacheRouteExclude: (state, path) => {
            state.cacheRoute.exclude = state.cacheRoute.exclude.filter((item) => {
                return item !== path
            })
        },
        /**
         * 清空菜单标签列表
         */
        clearMenuTabList: (state) => {
            state.menuTabList = []
        },
        /**
         * 推入菜单标签列表
         * @param {*} routeTo  将要打开的路由
         */
        pushMenuTabList: (state, routeTo) => {
            let result = state.menuTabList.findIndex((item) => {
                return item.path === routeTo.path
            })
            if (result === -1) {
                state.menuTabList.push({
                    title: routeTo.meta.title,
                    path: routeTo.path,
                    icon: routeTo.meta.icon,
                    closable: routeTo.closable === false ? false : true,
                })
            }
        },
        /**
         * 刷新菜单标签
         *      实现流程：
         *          1：app-container.vue文件内component标签加上判断是否允许缓存，允许才显示界面（v-if="$store.state.user.cacheRoute.exclude.indexOf(route.path) === -1"）
         *          2：设置路由不允许缓存，不显示页面
         *          3：打开路由，路由后置守卫afterEach中重新设置成允许缓存，显示页面
         * @param {*} path  菜单标签的路由路径
         */
        refreshMenuTab: (state, path) => {
            state.cacheRoute.exclude.push(path)
            const currentPath = getCurrentPath()
            if (path === currentPath) {
                router.push(path)
            }
        },
        /**
         * 关闭自身菜单标签
         * @param {*} path  菜单标签的路由路径
         */
        closeSelfMenuTab: (state, path) => {
            state.menuTabList = state.menuTabList.filter((item) => {
                return !item.closable || item.path !== path
            })
            const currentPath = getCurrentPath()
            if (path === currentPath) {
                router.push(state.menuTabList[state.menuTabList.length - 1].path)
            }
        },
        /**
         * 关闭其他菜单标签
         * @param {*} path  菜单标签的路由路径
         */
        closeOtherMenuTab: (state, path) => {
            state.menuTabList = state.menuTabList.filter((item) => {
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
        closeLeftMenuTab: (state, path) => {
            const leftIndex = state.menuTabList.findIndex((item) => {
                return item.path === path
            })
            state.menuTabList = state.menuTabList.filter((item, index) => {
                return !item.closable || index >= leftIndex
            })
            const currentPath = getCurrentPath()
            if (path !== currentPath) {
                const currentLeftIndex = state.menuTabList.findIndex((item) => {
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
        closeRightMenuTab: (state, path) => {
            const rightIndex = state.menuTabList.findIndex((item) => {
                return item.path === path
            })
            state.menuTabList = state.menuTabList.filter((item, index) => {
                return !item.closable || index <= rightIndex
            })
            const currentPath = getCurrentPath()
            if (path !== currentPath) {
                const currentRightIndex = state.menuTabList.findIndex((item) => {
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
        closeAllMenuTab: (state) => {
            state.menuTabList = state.menuTabList.filter((item) => {
                return !item.closable
            })
            router.push(state.menuTabList[state.menuTabList.length - 1].path)
        }
    },
    actions: {
        /**
         * 登录
         * @param {*} context  包含{ state, getter, commit, dispatch }
         * @param {*} payload
         * @param {*} payload.account   账号
         * @param {*} payload.password  密码
         * @returns 
         */
        login: async ({ commit, dispatch }, payload) => {
            try {
                const { account, password } = payload
                let res = await getLoginToken({
                    account: account
                })
                res = await login({
                    account: account,
                    password: md5(md5(password) + res.data.loginToken)
                })
                setAccessToken(res.data.PlatformToken)
                /**--------初始化数据（可有效清理上一个登录用户的脏数据） 开始--------**/
                //在logout退出登录操作中也可以清理，但在登录操作这里处理，应变能力更好。不用考虑有多少种情况需及时清理脏数据，如：accessToken失效、切换用户等
                //commit('setInfo', {})   //清空用户信息
                await dispatch('setInfo')   //设置用户信息（可选，路由前置守卫有执行，此处执行，路由可减少一次跳转）
                commit('resetRoute')    //重置路由
                //commit('setLeftMenuTree', [])   //清空用户左侧菜单
                await dispatch('setLeftMenuTree')   //设置左侧菜单树（可选，路由前置守卫有执行，此处执行，路由可减少一次跳转）
                commit('clearMenuTabList')    //清空菜单标签列表
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
        setInfo: async ({ commit }) => {
            try {
                const res = await getInfo()
                commit('setInfo', {
                    nickname: res.data.info.nickname ? res.data.info.nickname : res.data.info.account,
                    avatar: res.data.info.avatar,
                    rawInfo: res.data.info,
                })
                return true
            } catch (err) {
                await errorHandle(err)
                return false
            }
        },
        /**
         * 设置左侧菜单（包含注册动态路由）
         */
        setLeftMenuTree: async ({ commit }) => {
            try {
                const res = await getMenuTree()
                /**--------注册动态路由 开始--------**/
                const layoutName = config('app.router.layoutName')
                const handleMenuTree = (menuTree, pMenuList = []) => {
                    const leftMenuTree = []
                    for (let i = 0; i < menuTree.length; i++) {
                        leftMenuTree[i] = {
                            title: menuTree[i].menuName,
                            path: menuTree[i].menuUrl,
                            icon: menuTree[i].menuIcon,
                            children: [],
                        }
                        if (menuTree[i].children.length) {
                            pMenuList.push({
                                title: menuTree[i].menuName,
                                path: menuTree[i].menuUrl,
                                icon: menuTree[i].menuIcon,
                            })
                            leftMenuTree[i].children = handleMenuTree(menuTree[i].children, Object.assign({}, pMenuList))
                            pMenuList.pop()
                        } else {
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
                        }
                    }
                    return leftMenuTree
                }
                const leftMenuTree = handleMenuTree(res.data.tree)
                commit('setLeftMenuTree', leftMenuTree)
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
        logout: async ({}, toPath = '/login') => {
            await removeAccessToken()
            const whiteList = config('app.router.whiteList')
            if (whiteList.indexOf(toPath) === -1) {
                await router.push('/login?redirect=' + toPath)
            } else {
                await router.push(toPath)
            }
        }
    }
}
