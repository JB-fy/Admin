/*--------使用方式 开始--------*/
/* const route = useRoute()
route.path
route.query
route.params.lotteryId
route.fulPath

const router = useRouter()
router.push(url) 
router.currentRoute.value  //当前页面路由 */
/*--------使用方式 结束--------*/
import { createRouter, createWebHistory } from 'vue-router'
import Layout from '@/layout/default/Index.vue'

/**
 * meta说明：（当路由在后端数据库菜单表中未记录时，menu必须设置，反之不用设置。例如：个人中心不在用户菜单中，则需要设置）
 *      isAuth: true,   //是否需要权限验证
 *      keepAlive: true,    //是否可以缓存
 *      componentName: string,    //组件名称
 *      menu: {
 *          i18n: {    //标题，多语言时设置，未设置以name为准
 *              title: {
 *                  'en': 'homepage',
 *                  'zh-cn': '主页',...
 *              },
 *          },
 *          icon: '图标',
 *      }
 */
const initRouteList = [
    {
        path: '/layout', //必须设置，否则默认为'/'。在多个路由没有设置该参数时，则首页会以最后一个路由为准，会出现首页错误问题
        component: Layout,
        redirect: '/',
        replace: true,
        children: [
            {
                path: '/',
                component: async () => {
                    /**
                     * import说明
                     *   参数为静态（无变量）时，可以使用任意方式导入。即能使用@路径
                     *   参数为动态（有变量）时，必须是相对路径或绝对路径方式，其它方式不允许。（查看文档绝对路径也不允许，但实际可以使用）
                     */
                    //let componentPath='../views/Index.vue'
                    //const component = await import(componentPath)
                    const component = await import('@/views/Index.vue')
                    component.default.name = '/' //meta.keepAlive为true时，要实现组件缓存和页面刷新，必须设置组件name和meta.componentName，且必须相同
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/' },
            },
            {
                path: '/auth/action',
                component: async () => {
                    const component = await import('@/views/auth/action/Index.vue')
                    component.default.name = '/auth/action'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/auth/action' },
            },
            {
                path: '/auth/menu',
                component: async () => {
                    const component = await import('@/views/auth/menu/Index.vue')
                    component.default.name = '/auth/menu'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/auth/menu' },
            },
            {
                path: '/auth/role',
                component: async () => {
                    const component = await import('@/views/auth/role/Index.vue')
                    component.default.name = '/auth/role'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/auth/role' },
            },
            {
                path: '/auth/scene',
                component: async () => {
                    const component = await import('@/views/auth/scene/Index.vue')
                    component.default.name = '/auth/scene'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/auth/scene' },
            },
            {
                path: '/platform/admin',
                component: async () => {
                    const component = await import('@/views/platform/admin/Index.vue')
                    component.default.name = '/platform/admin'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/platform/admin' },
            },
            {
                path: '/platform/config/app',
                component: async () => {
                    const component = await import('@/views/platform/config/App.vue')
                    component.default.name = '/platform/config/app'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/platform/config/platform' },
            },
            {
                path: '/platform/config/plugin',
                component: async () => {
                    const component = await import('@/views/platform/config/Plugin.vue')
                    component.default.name = '/platform/config/plugin'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/platform/config/plugin' },
            },
            {
                path: '/upload/upload',
                component: async () => {
                    const component = await import('@/views/upload/upload/Index.vue')
                    component.default.name = '/upload/upload'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/upload/upload' },
            },
            {
                path: '/pay/pay',
                component: async () => {
                    const component = await import('@/views/pay/pay/Index.vue')
                    component.default.name = '/pay/pay'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/pay/pay' },
            },
            {
                path: '/pay/scene',
                component: async () => {
                    const component = await import('@/views/pay/scene/Index.vue')
                    component.default.name = '/pay/scene'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/pay/scene' },
            },
            {
                path: '/app/app',
                component: async () => {
                    const component = await import('@/views/app/app/Index.vue')
                    component.default.name = '/app/app'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/app/app' },
            },
            {
                path: '/users/users',
                component: async () => {
                    const component = await import('@/views/users/users/Index.vue')
                    component.default.name = '/users/users'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/users/users' },
            },
            {
                path: '/org/org',
                component: async () => {
                    const component = await import('@/views/org/org/Index.vue')
                    component.default.name = '/org/org'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/org/org' },
            },
            {
                path: '/org/admin',
                component: async () => {
                    const component = await import('@/views/org/admin/Index.vue')
                    component.default.name = '/org/admin'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/org/admin' },
            },
            /*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
            {
                path: '/profile',
                component: async () => {
                    const component = await import('@/views/Profile.vue')
                    component.default.name = '/profile'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/profile', menu: { i18n: { title: { en: 'Profile', 'zh-cn': '个人中心' } }, icon: 'autoicon-ep-user-filled' } },
            },
            {
                path: '/thirdSite', //必须带query.url参数。示例：/thirdSite?url=https://element-plus.gitee.io/zh-CN/
                component: () => import('@/views/ThirdSite.vue'),
                meta: { isAuth: true, keepAlive: false, menu: { i18n: { title: { en: 'Third Site', 'zh-cn': '第三方站点' } }, icon: 'autoicon-ep-chrome-filled' } },
            },
        ],
    },
    {
        path: '/login',
        component: () => import('@/views/Login.vue'),
        meta: { isAuth: false, keepAlive: false, menu: { i18n: { title: { en: 'Login', 'zh-cn': '登录' } } } },
    },
    {
        path: '/:pathMatch(.*)*',
        component: () => import('@/views/404.vue'),
        meta: { isAuth: false, keepAlive: false, menu: { i18n: { title: { en: '404', 'zh-cn': '404' } } } },
    },
]

const router = createRouter({
    //history: createWebHistory(import.meta.env.VITE_BASE_PATH),
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: initRouteList,
})

router.beforeEach(async (to: any) => {
    const adminStore = useAdminStore()
    /**--------判断登录状态 开始--------**/
    const accessToken = getAccessToken()
    if (!accessToken) {
        if (to.meta.isAuth) {
            /* //不需要做这步，清理工作换到登录操作中执行，应变能力更好
            adminStore.logout(to.path)
            return false */
            return '/login?redirect=' + to.fullPath
        }
        document.title = useLanguageStore().getWebTitle(to.fullPath)
        return true
    }
    if (to.path === '/login') {
        //已登录且链接是登录页面时，则跳到首页
        if (to.query.redirect) {
            return to.query.redirect
        }
        return '/'
    }
    /**--------判断登录状态 结束--------**/

    /**--------设置用户相关的数据（因用户在浏览器层面刷新页面，会导致pinia数据全部重置） 开始--------**/
    if (!adminStore.infoIsExist) {
        try {
            await adminStore.setInfo() //记录用户信息
            await adminStore.setMenuTree() //设置左侧菜单
            await adminStore.setActionCodeArr() //设置操作权限标识数组
        } catch (error) {
            return false
        }
    }
    /**--------设置用户相关的数据（因用户在浏览器层面刷新页面，会导致pinia数据全部重置） 结束--------**/

    /**--------设置菜单标签 开始--------**/
    if (to.meta.isAuth) {
        adminStore.pushMenuTabList({
            keepAlive: to.meta?.keepAlive ?? false,
            componentName: to.meta?.componentName,
            url: to.fullPath,
            ...to.meta?.menu,
        })
    }
    /**--------设置菜单标签 结束--------**/

    document.title = useLanguageStore().getWebTitle(to.fullPath)
    return true
})

router.afterEach((to) => {
    useKeepAliveStore().removeAppContainerExclude(to.meta?.componentName as string) //打开后重新设置成允许缓存，主要用于实现缓存刷新
})

export default router
