import { createRouter, createWebHistory } from 'vue-router';
//import layout from '@/app/layout/default/index.vue';
//import { useUserStore } from '@/stores/user';

const initRoutes = [
    // {
    //     path: '/',  //必须设置，否则默认为'/'。在多个路由没有设置该参数时，则首页会以最后一个路由为准，会出现首页错误问题
    //     component: layout,
    //     meta: { title: '首页' },
    //     name: config('app.router.layoutName'),
    //     redirect: config('app.router.indexRedirect'),
    //     replace: true,
    //     children: [
    //         {
    //             path: '/profile',
    //             component: async () => {
    //                 let component = await import('@/views/profile')
    //                 component.default.name = '/profile'    //动态设置页面组件名称，方便清理缓存
    //                 return component
    //             },
    //             meta: { title: '个人中心', icon: 'vant-manager-o' }
    //         },
    //     ]
    // },
    {
        path: '/login',
        component: () => import('@/views/login/index.vue'),
        meta: { title: '登录' }
    },
    {
        path: '/:pathMatch(.*)*',
        component: () => import('@/views/404/index.vue'),
    },
]

const router = createRouter({
    history: createWebHistory('/view/admin/system'),
    routes: initRoutes
})

// router.beforeEach(async (to) => {
//     const userStore = useUserStore();
//     const webTitle = config('app.webTitle')
//     if (to.meta.title) {
//         document.title = webTitle + '-' + to.meta.title
//     } else {
//         document.title = webTitle
//     }

//     /**--------判断登录状态 开始--------**/
//     const accessToken = getAccessToken()
//     if (!accessToken) {
//         const whiteList = config('app.router.whiteList')
//         if (whiteList.indexOf(to.path) === -1) {
//             /* //不需要做这步，清理工作换到登录操作中执行，应变能力更好
//             await userStore.logout(to.path)
//             return false */
//             return '/login?redirect=' + to.path
//         }
//         return true
//     }
//     if (to.path === '/login') {
//         //已登录且链接是登录页面时，则跳到首页
//         return '/'
//     } else {
//         /**--------设置用户相关的数据（因用户在浏览器层面刷新页面，会导致vuex数据全部重置） 开始--------**/
//         if (!userStore.infoIsExist) {
//             let result = await userStore.setInfo()  //记录用户信息
//             if (!result) {
//                 return false
//             }
//             result = await userStore.setLeftMenuTree()  //设置左侧菜单（包含注册动态路由）
//             if (!result) {
//                 return false
//             }
//             return to.path  //由于这个路由之前不存在，会报404。此时才注册成功，需要跳转自身
//         }
//         /**--------设置用户相关的数据（因用户在浏览器层面刷新页面，会导致vuex数据全部重置） 结束--------**/
//     }
//     /**--------判断登录状态 结束--------**/

//     /**--------设置菜单标签 开始--------**/
//     if (userStore.menuTabListLength === 0) {
//         const routes = router.getRoutes()
//         const initRouteTo = routes.find((item) => {
//             return item.path === initRoutes[0].redirect
//         })
//         userStore.pushMenuTabList(Object.assign({ closable: false }, initRouteTo))
//     }
//     userStore.pushMenuTabList(to)
//     /**--------设置菜单标签 结束--------**/

//     return true
// })

// router.afterEach((to) => {
//     const userStore = useUserStore();
//     userStore.removeCacheRouteExclude(to.path)  //打开后重新设置成允许缓存，主要用于实现缓存刷新
// })

export default router

/*--------使用方式 开始--------*/
/* const route = useRoute()
route.path
route.query
route.params.lotteryId

const router = useRouter()
router.push(url) 
router.currentRoute.value.path  //当前页面地址
*/
/*--------使用方式 结束--------*/