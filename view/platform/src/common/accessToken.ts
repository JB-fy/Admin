const storage = import.meta.env.VITE_ACCESS_TOKEN_STORAGE === 'localStorage' ? localStorage : sessionStorage
const accessTokenName = import.meta.env.VITE_ACCESS_TOKEN_NAME
const activeTimeName = import.meta.env.VITE_ACCESS_TOKEN_ACTIVE_TIME_NAME
const activeTimeout = parseInt(import.meta.env.VITE_ACCESS_TOKEN_ACTIVE_TIMEOUT)

//获取accessToken
export const getAccessToken = (): string | false => {
    const accessToken = storage.getItem(accessTokenName)
    if (accessToken) {
        if (isActiveAccessToken()) {
            return accessToken
        }
        removeAccessToken()
    }
    return false
}

//设置accessToken
export const setAccessToken = (token: string): void => {
    if (activeTimeout > 0) {
        const nowTime = new Date().getTime().toString()
        storage.setItem(activeTimeName, nowTime)
    }
    storage.setItem(accessTokenName, token)
}

//删除accessToken
export const removeAccessToken = (): void => {
    if (activeTimeout > 0) {
        storage.removeItem(activeTimeName)
    }
    storage.removeItem(accessTokenName)
}

//判断accessToken是否活跃（调用getAccessToken函数的地方需要马上使用这个函数验证）
export const isActiveAccessToken = (): boolean => {
    if (activeTimeout > 0) {
        const activeTime: any = storage.getItem(activeTimeName)
        const nowTime: any = new Date().getTime().toString()
        if (nowTime - activeTime > activeTimeout) {
            return false
        }
        storage.setItem(activeTimeName, nowTime)
    }
    return true
}

/* const accessToken = {
    get: getAccessToken,
    set: setAccessToken,
    remove: removeAccessToken,
    isActive: isActiveAccessToken
}
export default accessToken */
