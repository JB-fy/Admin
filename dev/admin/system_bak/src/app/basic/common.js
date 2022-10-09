/**
 * 处理require.context批量加载的组件
 * @param {*} requireComponents  require.context批量加载的组件
 * @param {*} prefix    是否增加前缀
 * @returns 
 */
export const getImportComponents = (requireComponents, prefix = '') => {
    //const requireComponents = require.context(path, isRecurs, /\.vue$/)   //require.context的参数只能使用字面值，不能使用变量。且第三个参数一定要传，否则会多一个空字符串''的键
    prefix = prefix ? prefix + '-' : '';
    let importComponents = {}
    let componentName;
    requireComponents.keys().forEach((fileName) => {
        //componentName = requireComponents(fileName).default.name
        /*--------去掉头部'./'和尾部'.**'，并将路径/替换成- 开始--------*/
        //componentName = fileName.replace(/^\.\/(.*)\.\w+$/, '$1').replace('/', '-')
        componentName = fileName.slice(2, fileName.lastIndexOf('.')).replace('/', '-')
        /*--------去掉头部'./'和尾部'.**'，并将路径/替换成- 结束--------*/
        importComponents[prefix + componentName] = requireComponents(fileName).default
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
    if (!(key in process.env)) {
        return def
    }
    let value = process.env[key]
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

import router from '@/app/basic/router'

/**
 * 获取当前路由路径（需在路由稳定时使用。路由切换时不能使用，此时得到的是上一个页面的路由路径，如：浏览器刷新页面；各路由对应的.vue文件内在export default之前调用。）
 * @returns 
 */
export const getCurrentPath = () => {
    return router.currentRoute.value.path
}