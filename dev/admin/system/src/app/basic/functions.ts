/**
 * 自动导入文件
 * @param {*} rawFileList  加载的文件（import.meta.globEager方法）
 * @param {*} prefix    是否增加前缀
 * @returns 
 */
export const autoImport = (rawFileList, prefix = '') => {
    prefix = prefix ? prefix + '-' : '';
    let fileList = {}
    let name;
    Object.keys(rawFileList).forEach((fileName) => {
        name = fileName.slice(fileName.lastIndexOf('/') + 1, fileName.lastIndexOf('.'))
        fileList[prefix + name] = rawFileList[fileName].default
        //fileList[prefix + name] = rawFileList[fileName]().default
    })
    return fileList
}

/**
 * 获取env变量值
 * @param {*} key  键名
 * @param {*} def  默认值
 * @returns 
 */
/* export const env = (key, def = '') => {
    if (!(key in import.meta.env)) {
        return def
    }
    let value = import.meta.env[key]
    switch (value) {
        case 'true':
            return true
        case 'false':
            return false
        case 'null':
            return null
        case 'localStorage':
            return localStorage
        case 'sessionStorage':
            return sessionStorage
        default:
            let tmpValue
            //十进制字符串返回数字格式
            tmpValue = Number(value, 10)
            if (!isNaN(tmpValue)) {
                return tmpValue
            }
            //exp:开头则当作表达式运行
            if (value.indexOf('exp:') === 0) {
                tmpValue = value.slice(4)
                //return eval(tmpValue)
                return new Function('return ' + tmpValue)()
            }
            return value
    }
} */

const allConfig = autoImport(import.meta.globEager('@/app/config/*.ts')) //放外面，不用每次调用getConfig都要加载这个目录
/**
 * 获取配置参数
 * @param {*} key   键名。以'.'分隔，格式：文件名.key1.key2...
 * @param {*} defaultValue  默认值
 * @returns 
 */
export const config = (key: string, defaultValue: any = null) => {
    let keyArr = key.split('.')
    let value = allConfig
    for (let i = 0; i < keyArr.length; i++) {
        if (keyArr[i] in value) {
            value = value[keyArr[i]]
        } else {
            return defaultValue
        }
    }
    return value
}

import router from '@/router'
/**
 * 获取当前路由路径（需在路由稳定时使用。路由切换时不能使用，此时得到的是上一个页面的路由路径，如：浏览器刷新页面；各路由对应的.vue文件内在export default之前调用。）
 * @returns 
 */
export const getCurrentPath = () => {
    return router.currentRoute.value.path
}