import { useUserStore } from '@/stores/user'
import router from '@/router'

/**
 * 错误处理
 * @param {*} err 错误信息。格式：JSON；包含字段：{ code: 9999, msg: '失败', data: {} }
 */
export const errorHandle = async (err: any) => {
    const errMsg = JSON.parse(err.message)
    if (typeof errMsg.code !== 'undefined') {
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
    }
}
