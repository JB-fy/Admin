import { createRouter, createWebHistory } from 'vue-router'
import Layout from '@/layout/default/Index.vue'

/**
 * meta说明：（当路由在后端数据库菜单表中未记录时，menu必须设置，反之不用设置。例如：个人中心不在用户菜单中，则需要设置）
 *      keepAlive: true,    //是否可以缓存
 *      isAuth: true,   //是否需要权限验证  
 *      menu: { menuName: '菜单名称', title: {'en': 'homepage', 'zh-cn': '主页',...}, icon:'图标'}，    //菜单配置
 *          menuName: '菜单名称', //菜单名称。
 *          title: {'en': 'homepage', 'zh-cn': '主页'},  //标题，多语言时设置，未设置以menuName为准。
 *          icon: '图标',  //图标。
 */
const initRouteList = [
    {
        path: '/layout',  //必须设置，否则默认为'/'。在多个路由没有设置该参数时，则首页会以最后一个路由为准，会出现首页错误问题
        component: Layout,
        redirect: '/',
        replace: true,
        children: [
            {
                path: '/',
                component: async () => {
                    /*  import说明
                            参数为静态（无变量）时，可以使用任意方式导入。即能使用@路径
                            参数为动态（有变量）时，必须是相对路径或绝对路径方式，其他方式不允许。（查看文档绝对路径也不允许，但实际可以使用）
                    */
                    //let componentPath='../views/index/Index.vue'
                    //const component = await import(componentPath)
                    const component = await import('@/views/index/Index.vue')
                    component.default.name = '/'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { keepAlive: true, isAuth: true }
            },
            {
                path: '/authAction',
                component: async () => {
                    const component = await import('@/views/auth/action/Index.vue')
                    component.default.name = '/authAction'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { keepAlive: true, isAuth: true }
            },
            {
                path: '/authMenu',
                component: async () => {
                    const component = await import('@/views/auth/menu/Index.vue')
                    component.default.name = '/authMenu'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { keepAlive: true, isAuth: true }
            },
            {
                path: '/authRole',
                component: async () => {
                    const component = await import('@/views/auth/role/Index.vue')
                    component.default.name = '/authRole'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { keepAlive: true, isAuth: true }
            },
            {
                path: '/authScene',
                component: async () => {
                    const component = await import('@/views/auth/scene/Index.vue')
                    component.default.name = '/authScene'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { keepAlive: true, isAuth: true }
            },
            {
                path: '/systemAdmin',
                component: async () => {
                    const component = await import('@/views/system/admin/Index.vue')
                    component.default.name = '/systemAdmin'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { keepAlive: true, isAuth: true }
            },
            {
                path: '/systemConfig',
                component: async () => {
                    const component = await import('@/views/system/config/Index.vue')
                    component.default.name = '/systemConfig'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { keepAlive: true, isAuth: true }
            },
            {
                path: '/systemLogOfRequest',
                component: async () => {
                    const component = await import('@/views/system/logOfRequest/Index.vue')
                    component.default.name = '/systemLogOfRequest'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { keepAlive: true, isAuth: true }
            },
            {
                path: '/profile',
                component: async () => {
                    const component = await import('@/views/profile/Index.vue')
                    component.default.name = '/profile'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { keepAlive: true, isAuth: true, menu: { menuName: '个人中心', title: { 'en': 'Profile', 'zh-cn': '个人中心' }, icon: 'AutoiconEpUserFilled' } }
            },
            {
                path: '/thirdUrl/:userId?',
                component: {
                    template: '<iframe :src="$route.query.url" frameborder="0" style="width: 100%; height: calc(100vh - 194px);"></iframe>',
                },
                meta: { keepAlive: false, isAuth: true, menu: { menuName: '第三方网站', title: { 'en': 'thridWebsite', 'zh-cn': '第三方网站' }, icon: 'AutoiconEpChromeFilled' } }
            },
        ]
    },
    {
        path: '/login',
        component: () => import('@/views/login/Index.vue'),
        meta: { keepAlive: false, isAuth: false, menu: { menuName: '登录', title: { 'en': 'Login', 'zh-cn': '登录' } } }
    },
    {
        path: '/:pathMatch(.*)*',
        component: () => import('@/views/404/Index.vue'),
        meta: { keepAlive: false, isAuth: false, menu: { menuName: '404', title: { 'en': '404', 'zh-cn': '404' } } }
    },
]

const router = createRouter({
    //history: createWebHistory(import.meta.env.VITE_BASE_PATH),
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: initRouteList
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

    /**--------设置用户相关的数据（因用户在浏览器层面刷新页面，会导致vuex数据全部重置） 开始--------**/
    if (!adminStore.infoIsExist) {
        try {
            await adminStore.setInfo()  //记录用户信息
            await adminStore.setMenuTree()  //设置左侧菜单
        } catch (error) {
            await errorHandle(<Error>error)
            return false
        }
    }
    /**--------设置用户相关的数据（因用户在浏览器层面刷新页面，会导致vuex数据全部重置） 结束--------**/

    /**--------设置菜单标签 开始--------**/
    //404不放入菜单标签中
    if (to.meta.isAuth) {
        adminStore.pushMenuTabList({
            url: to.fullPath,
            ...to.meta?.menu,
            keepAlive: to.meta?.keepAlive ?? false,
        })
    }
    /**--------设置菜单标签 结束--------**/

    document.title = useLanguageStore().getWebTitle(to.fullPath)
    return true
})

router.afterEach((to) => {
    useKeepAliveStore().removeAppContainerExclude(to.fullPath)  //打开后重新设置成允许缓存，主要用于实现缓存刷新
})

export default router

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