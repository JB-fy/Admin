import { createRouter, createWebHistory } from 'vue-router';
import layout from '@/app/layout/default/Index.vue';
import { useUserStore } from '@/stores/user';
import { useKeepAliveStore } from '@/stores/keepAlive';


/*
meta说明：
    title: '主页',  //标题
    keepAlive: true,    //是否可以缓存
    isAuth: true,   //是否需要权限验证
 */
const initRouteList = [
    {
        path: '/layout',  //必须设置，否则默认为'/'。在多个路由没有设置该参数时，则首页会以最后一个路由为准，会出现首页错误问题
        component: layout,
        redirect: '/',
        replace: true,
        children: [
            {
                path: '/',
                component: async () => {
                    //let componentPath='../views/index/Index.vue'
                    //const component = await import(componentPath)
                    const component = await import('@/views/index/Index.vue')
                    component.default.name = '/authAction'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { title: '主页', keepAlive: true, isAuth: true, isIndexMenuTab: true }
            },
            {
                path: '/authAction',
                component: async () => {
                    const component = await import('@/views/auth/action/Index.vue')
                    component.default.name = '/authAction'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { title: '操作列表', keepAlive: true, isAuth: true }
            },
            {
                path: '/authMenu',
                component: async () => {
                    const component = await import('@/views/auth/menu/Index.vue')
                    component.default.name = '/authMenu'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { title: '菜单列表', keepAlive: true, isAuth: true }
            },
            {
                path: '/authRole',
                component: async () => {
                    const component = await import('@/views/auth/role/Index.vue')
                    component.default.name = '/authRole'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { title: '角色列表', keepAlive: true, isAuth: true }
            },
            {
                path: '/authScene',
                component: async () => {
                    const component = await import('@/views/auth/scene/Index.vue')
                    component.default.name = '/authScene'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { title: '场景列表', keepAlive: true, isAuth: true }
            },
            {
                path: '/systemAdmin',
                component: async () => {
                    const component = await import('@/views/system/admin/Index.vue')
                    component.default.name = '/systemAdmin'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { title: '系统管理员', keepAlive: true, isAuth: true }
            },
            {
                path: '/systemConfig',
                component: async () => {
                    const component = await import('@/views/system/config/Index.vue')
                    component.default.name = '/systemConfig'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { title: '系统配置', keepAlive: true, isAuth: true }
            },
            {
                path: '/systemLogOfRequest',
                component: async () => {
                    const component = await import('@/views/system/logOfRequest/Index.vue')
                    component.default.name = '/systemLogOfRequest'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { title: '请求日志', keepAlive: true, isAuth: true }
            },
            {
                path: '/profile',
                component: async () => {
                    const component = await import('@/views/profile/Index.vue')
                    component.default.name = '/profile'    //设置页面组件name为path，方便清理缓存
                    return component
                },
                meta: { title: '个人中心', keepAlive: true, isAuth: true }
            },
        ]
    },
    {
        path: '/login',
        component: () => import('@/views/login/Index.vue'),
        meta: { title: '登录' }
    },
    {
        path: '/:pathMatch(.*)*',
        component: () => import('@/views/404/Index.vue'),
    },
]

const router = createRouter({
    //history: createWebHistory(import.meta.env.VITE_BASE_PATH),
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: initRouteList
})

router.beforeEach(async (to) => {
    const userStore = useUserStore();
    const webTitle = config('app.webTitle')
    if (to.meta.title) {
        document.title = webTitle + '-' + to.meta.title
    } else {
        document.title = webTitle
    }

    /**--------判断登录状态 开始--------**/
    const accessToken = getAccessToken()
    if (!accessToken) {
        if (to.meta.isAuth) {
            /* //不需要做这步，清理工作换到登录操作中执行，应变能力更好
            await userStore.logout(to.path)
            return false */
            return '/login?redirect=' + to.path
        }
        return true
    }
    if (to.path === '/login') {
        //已登录且链接是登录页面时，则跳到首页
        return '/'
    } else {
        /**--------设置用户相关的数据（因用户在浏览器层面刷新页面，会导致vuex数据全部重置） 开始--------**/
        if (!userStore.infoIsExist) {
            let result = await userStore.setInfo()  //记录用户信息
            if (!result) {
                return false
            }
            result = await userStore.setMenuTree()  //设置左侧菜单（包含注册动态路由）
            if (!result) {
                return false
            }
            return to.path  //由于这个路由之前不存在，会报404。此时才注册成功，需要跳转自身
        }
        /**--------设置用户相关的数据（因用户在浏览器层面刷新页面，会导致vuex数据全部重置） 结束--------**/
    }
    /**--------判断登录状态 结束--------**/

    /**--------设置菜单标签 开始--------**/
    userStore.pushMenuTabList({
        title: <string>to.meta.title ?? '',
        path: to.path,
        icon: <string>to.meta.icon ?? ''
    })
    /**--------设置菜单标签 结束--------**/

    return true
})

router.afterEach((to) => {
    useKeepAliveStore().removeAppContainerExclude(to.path)  //打开后重新设置成允许缓存，主要用于实现缓存刷新
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
router.currentRoute.value.path  //当前页面地址 */
/*--------使用方式 结束--------*/