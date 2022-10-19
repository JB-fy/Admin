import { useUserStore } from '@/stores/user'

/**
 * 错误处理
 * @param {*} err 错误信息。格式：JSON；包含字段：{ code: 9999, msg: '失败', data: {} }
 */
export const errorHandle = async (err: Error) => {
    try {
        const errMsg = JSON.parse(err.message)
        if (typeof errMsg.code !== 'undefined') {
            switch (errMsg.code) {
                //case '000404':
                case '001400':
                    /* ElMessageBox.alert(errMsg.msg, '确认登出', {
                        confirmButtonText: '重新登录',
                        type: 'warning'
                    }).then(async () => {
                        await useUserStore().logout()
                    }).catch(async () => {
                        await useUserStore().logout()
                    }) */
                    useUserStore().logout(getCurrentPath())
                    ElMessage.error(errMsg.msg)
                    break
                default:
                    ElMessage.error(errMsg.msg)
                    break
            }
        }
    } catch (e) {
        ElMessage.error((<Error>e).message);
    }
}
