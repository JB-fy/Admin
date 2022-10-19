/**
 * 获取accessToken
 */
export const getAccessToken = () => {
    const storage = config('app.accessToken.storage')
    const accessTokenName = config('app.accessToken.name')

    let accessToken = storage.getItem(accessTokenName)
    if (accessToken && !isActiveAccessToken()) {
        removeAccessToken()
        return null
    }
    return accessToken
}

/**
 * 设置accessToken
 * @param {*} token 
 */
export const setAccessToken = (token: string) => {
    const storage = config('app.accessToken.storage')
    const accessTokenName = config('app.accessToken.name')
    const activeTimeName = config('app.accessToken.activeTimeName')
    const activeTimeout = config('app.accessToken.activeTimeout')
    if (activeTimeout > 0) {
        let nowTime = new Date().getTime().toString()
        storage.setItem(activeTimeName, nowTime)
    }
    storage.setItem(accessTokenName, token)
}

/**
 * 删除accessToken
 */
export const removeAccessToken = () => {
    const storage = config('app.accessToken.storage')
    const accessTokenName = config('app.accessToken.name')
    const activeTimeName = config('app.accessToken.activeTimeName')
    const activeTimeout = config('app.accessToken.activeTimeout')
    if (activeTimeout > 0) {
        storage.removeItem(activeTimeName)
    }
    storage.removeItem(accessTokenName)
}

/**
 * 判断accessToken是否活跃（调用getAccessToken函数的地方需要马上使用这个函数验证）
 */
export const isActiveAccessToken = () => {
    const storage = config('app.accessToken.storage')
    const activeTimeName = config('app.accessToken.activeTimeName')
    const activeTimeout = config('app.accessToken.activeTimeout')
    if (activeTimeout > 0) {
        let activeTime = storage.getItem(activeTimeName)
        //let nowTime = new Date().getTime().toString()
        let nowTime: number = new Date().getTime()
        if (nowTime - activeTime > activeTimeout) {
            return false
        }
        storage.setItem(activeTimeName, nowTime)
    }
    return true
}