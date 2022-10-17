import { createApp } from "vue";
import App from "./App.vue";

import "./assets/main.css";

const app = createApp(App);

import router from './router';
app.use(router);

import { createPinia } from 'pinia';
const pinia = createPinia();
app.use(pinia);

app.mount('#app');

/*-------- 动态加载图标 开始 --------*/
//app.component('autoiconEpLollipop', autoiconEpLollipop)
import * as epIconList from '@element-plus/icons-vue'
for (let [key, component] of Object.entries(epIconList)) {
    app.component('Ep' + key, component)
    app.component('AutoiconEp' + key, component)    //兼容图标插件unplugin-icons，如插件以后支持动态加载<component :is="图标标识变量"/>，不用修改代码
}
/*-------- 动态加载图标 结束 --------*/