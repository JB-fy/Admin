/**
 * 批量导入
 * @param {*} rawImportList  导入列表（调用import.meta.globEager或import.meta.glob方法返回的数据）
 * @param {*} level    命名层次。特别注意：如果有不同层次文件时，与最深层相差大于等于level的所有浅层文件将不返回（默认以所有文件中的最深层文件为基准开始计算）
 * @param {*} isParentName    是否以父级文件夹命名。特别注意：此时命名层次以父级文件夹开始计算
 * @param {*} type    类型，默认0。0：一维对象（键名保持原样）；1：一维对象（键名小驼峰法）；2：一维对象（键名大驼峰法）；11：多维对象（键名保持原样）；
 * @returns 
 */
export const batchImport1 = async (rawImportList: any, level: number = 1, isParentName: boolean = false, type: number = 0) => {
    let importList: { [propName: string]: any } = {}
    let keyArr: string[] = []
    let keyList: string[][] = []
    let importArr: any[] = []
    let levelOfMax: number = 0
    for (const path in rawImportList) {
        console.log(path)
        keyArr = isParentName ? path.slice(0, path.lastIndexOf('/')).split('/') : path.slice(0, path.lastIndexOf('.')).split('/')
        keyList.push(keyArr)
        if (typeof rawImportList[path] === 'function') {
            importArr.push((await rawImportList[path]()).default)
        } else {
            importArr.push(rawImportList[path].default)
        }
        if (keyArr.length > levelOfMax || levelOfMax == 0) {
            levelOfMax = keyArr.length
        }
    }
    switch (type) {
        case 0:
        case 1:
        case 2:
            for (const key in keyList) {
                if (levelOfMax - keyList[key].length >= level) {
                    continue
                }
                let lengthTmp: number = keyList[key].length
                let keyTmp: string = ''
                for (let i = 0; i < level; i++) {
                    switch (type) {
                        case 0:
                            keyTmp = keyTmp + keyList[key][lengthTmp - level + i]
                            break;
                        case 1:
                        case 2:
                            const keyArrTmp = keyList[key][lengthTmp - level + i].split(/[\s-_]/)
                            if (type == 1 && i == 0) {
                                for (const key1 in keyArrTmp) {
                                    keyTmp = keyTmp + keyArrTmp[key1].slice(0, 1).toLowerCase() + keyArrTmp[key1].slice(1)
                                }
                            } else {
                                for (const key1 in keyArrTmp) {
                                    keyTmp = keyTmp + keyArrTmp[key1].slice(0, 1).toUpperCase() + keyArrTmp[key1].slice(1)
                                }
                            }
                            break;
                    }
                }
                importList[keyTmp] = importArr[key]
            }
            break;
        case 11:
            break;
    }
    return importList
}
console.log(await batchImport1(import.meta.globEager('@/app/config/*.ts'), 2, false, 1))
console.log(await batchImport1(import.meta.glob('@/app/config/*.ts'), 2, false, 2))
// const messages: any = batchImport1(import.meta.globEager('@/i18n/language/**/*.ts'), 2)
// console.log(messages)

/**
 * 批量导入
 * @param {*} rawImportList  导入列表（调用import.meta.globEager方法返回的数据）
 * @param {*} prefix    是否增加前缀
 * @returns 
 */
export const batchImport = (rawImportList: any, level: number = 1) => {
    let importList: any = {}
    let keyList: string[] = []
    for (const path in rawImportList) {
        keyList = path.slice(0, path.lastIndexOf('.')).split('/')
        importList[keyList[keyList.length - 1]] = rawImportList[path].default
    }
    return importList
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

const allConfig = batchImport(import.meta.globEager('@/app/config/*.ts')) //放外面，不用每次调用getConfig都要加载这个目录
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
export const getCurrentPath = (): string => {
    //return useRouter().currentRoute.value.path
    return router.currentRoute.value.path
}