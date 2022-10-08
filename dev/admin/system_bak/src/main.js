import { createApp } from 'vue'
import App from './App.vue'
import router from '@/app/basic/router'
import store from '@/app/basic/store'

store.commit('user/setConstRoutePathList') //记录固定路由，方便删除动态路由

const app = createApp(App)
app.use(router)
app.use(store)
app.mount('#app')

import '@/styles/index.css' //全局引入css变量

/*-------- 动态加载图标 开始 --------*/
//app.component('autoiconEpLollipop', autoiconEpLollipop)
import * as epIconList from '@element-plus/icons-vue'
for (let [key, component] of Object.entries(epIconList)) {
    app.component('ep' + key, component)
    app.component('autoiconEp' + key, component)    //兼容图标插件unplugin-icons，如插件以后支持动态加载<component :is="图标标识变量"/>，不用修改代码
}
/*-------- 动态加载图标 结束 --------*/