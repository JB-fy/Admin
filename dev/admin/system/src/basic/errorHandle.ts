import router from '@/router'

//自定义错误类
export class ApiError extends Error {
    constructor(message: string) {
        super(message)
        this.name = "ApiError"
    }
}

//错误处理
export const errorHandle = async (err: Error) => {
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
