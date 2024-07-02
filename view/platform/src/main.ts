import { createApp } from 'vue'
import App from './App.vue'

import './assets/css/main.css'

import i18n from './i18n'
import router from './router'
import { createPinia } from 'pinia'

const app = createApp(App)
app.use(i18n)
app.use(router)
app.use(createPinia())
app.mount('#app')

/*-------- 动态加载图标 开始 --------*/
import * as epIconList from '@element-plus/icons-vue'
for (const [key, component] of Object.entries(epIconList)) {
    app.component('AutoiconEp' + key, component) //兼容图标插件unplugin-icons，如插件以后支持动态加载<component :is="图标标识变量"/>，不用修改代码
}
/*-------- 动态加载图标 结束 --------*/

/*-------- 错误处理 开始 --------*/
//不好用，很多地方无法触发这个机制。例如：路由router.beforeEach前置导航守护; element plus的表单验证提交formRef.validate((valid) => { throw new Error('你大爷') })
// import { AxiosError } from "axios"; //这个错误类导入会报错，所以只能用err.name来识别错误类型
// import { ApiError } from "@/basic/http";
// import { useAdminStore } from '@/stores/admin'
// import { ElMessage } from 'element-plus';
// app.config.errorHandler = (err: any, instance, info) => {
//     //Error是所有错误类的父类。所以(err instanceof Error)一定是true，不能用于识别错误类型
//     /* if (err instanceof ApiError) {
//         console.log(1111)
//     } else if (err instanceof AxiosError) {
//         console.log(1111)
//     } else {
//     } */
//     let errMsg
//     switch (err.name) {
//         case 'ApiError':    //接口请求错误
//             errMsg = JSON.parse(err.message)
//             switch (errMsg.code) {
//                 case 39994000:
//                 case 39994001:
//                 case 39994002:
//                 case 39994003:
//                 case 39994004:
//                     /* ElMessageBox.alert(errMsg.msg, '确认登出', {
//                         confirmButtonText: '重新登录',
//                         type: 'warning'
//                     }).then(async () => {
//                         useAdminStore().logout()
//                     }).catch(async () => {
//                         useAdminStore().logout()
//                     }) */
//                     useAdminStore().logout(router.currentRoute.value.path)
//                     ElMessage.error(errMsg.msg)
//                     break
//                 default:
//                     ElMessage.error(errMsg.msg)
//                     break
//             }
//             break;
//         case 'AxiosError':  //Axios插件错误
//         default:
//             ElMessage.error(err.message)
//             break;
//     }
// }
/*-------- 错误处理 结束 --------*/
