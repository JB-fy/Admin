import router from '@/router'

//错误处理
export const errorHandle = async (err: Error) => {
    let errMsg
    switch (err.name) {
        case 'ApiError': //接口请求错误
            errMsg = JSON.parse(err.message)
            switch (errMsg.code) {
                case 39994000:
                case 39994001:
                case 39994002:
                case 39994003:
                case 39994004:
                    /* ElMessageBox.alert(errMsg.msg, '确认登出', {
                        confirmButtonText: '重新登录',
                        type: 'warning'
                    }).finally(() => useAdminStore().logout()) */
                    useAdminStore().logout(router.currentRoute.value.path)
                    ElMessage.error(errMsg.msg)
                    break
                default:
                    ElMessage.error(errMsg.msg)
                    break
            }
            break
        case 'AxiosError': //Axios插件错误
        default:
            ElMessage.error(err.message)
            break
    }
}
