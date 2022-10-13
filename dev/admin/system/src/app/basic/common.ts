/**
 * 处理require.context批量加载的组件
 * @param {*} requireComponents  require.context批量加载的组件
 * @param {*} prefix    是否增加前缀
 * @returns 
 */
export const getImportComponents = (requireComponents, prefix = '') => {
    prefix = prefix ? prefix + '-' : '';
    let importComponents = {}
    let componentName;
    Object.keys(requireComponents).forEach((fileName) => {
        componentName = fileName.slice(fileName.lastIndexOf('/') + 1, fileName.lastIndexOf('.'))
        importComponents[prefix + componentName] = requireComponents[fileName].default
        //importComponents[prefix + componentName] = requireComponents[fileName]().default
    })
    return importComponents
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

import router from '@/router'

/**
 * 获取当前路由路径（需在路由稳定时使用。路由切换时不能使用，此时得到的是上一个页面的路由路径，如：浏览器刷新页面；各路由对应的.vue文件内在export default之前调用。）
 * @returns 
 */
export const getCurrentPath = () => {
    return router.currentRoute.value.path
}