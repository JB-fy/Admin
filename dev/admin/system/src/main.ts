import { createApp } from "vue";
import App from "./App.vue";

import "./assets/main.css";

const app = createApp(App);

import router from './router';
app.use(router);

import { createPinia } from 'pinia';
app.use(createPinia());

import i18n from './i18n';
app.use(i18n);

app.mount('#app');

/*-------- 错误处理 开始 --------*/
//import { AxiosError } from "axios"; //这个错误类导入会报错，所以只能用err.name来识别错误类型
//import { ApiError } from "@/basic/http";
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus/es'
app.config.errorHandler = (err: any, instance, info) => {
    //Error是所有错误类的父类。所以(err instanceof Error)一定是true，不能用于识别错误类型
    /* if (err instanceof ApiError) {
        console.log(1111)
    } else if (err instanceof AxiosError) {
        console.log(1111)
    } else {
    } */
    switch (err.name) {
        case 'ApiError':    //接口请求错误
            const errMsg = JSON.parse(err.message)
            switch (errMsg.code) {
                //case '000404':
                case '001400':
                case '001401':
                case '001402':
                    /* ElMessageBox.alert(errMsg.msg, '确认登出', {
                        confirmButtonText: '重新登录',
                        type: 'warning'
                    }).then(async () => {
                        useUserStore().logout()
                    }).catch(async () => {
                        useUserStore().logout()
                    }) */
                    useUserStore().logout(router.currentRoute.value.path)
                    ElMessage.error(errMsg.msg)
                    break
                default:
                    ElMessage.error(errMsg.msg)
                    break
            }
            break;
        case 'AxiosError':  //Axios插件错误
        default:
            ElMessage.error(err.message)
            break;
    }
}
/*-------- 错误处理 结束 --------*/

/*-------- 动态加载图标 开始 --------*/
//app.component('autoiconEpLollipop', autoiconEpLollipop)
import * as epIconList from '@element-plus/icons-vue'
for (let [key, component] of (<any>Object).entries(epIconList)) {
    app.component('Ep' + key, component)
    app.component('AutoiconEp' + key, component)    //兼容图标插件unplugin-icons，如插件以后支持动态加载<component :is="图标标识变量"/>，不用修改代码
}
/*-------- 动态加载图标 结束 --------*/