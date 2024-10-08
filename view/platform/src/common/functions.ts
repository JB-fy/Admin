/* common.ts与functions.ts的区别：
common.ts：基于当前框架封装的常用函数（与框架耦合）
functions.ts：基于原生js封装的常用函数（不与框架耦合） */

/**
 * 获取浏览器当前语言
 * @param defaultValue
 * @returns
 */
export const getBrowserLanguage = (defaultValue: string = 'zh-cn'): string => {
    return (navigator.language || defaultValue).toLowerCase()
}

/**
 * 随机整数
 * @param min
 * @param max
 * @returns
 */
export const randomInt = (min: number, max: number) => {
    return Math.round(Math.random() * (max - min)) + min
}

/**
 * 是否空对象
 * @param obj
 * @returns
 */
export const isEmptyObj = (obj: any) => {
    if (Array.isArray(obj) && obj.length === 0) {
        return true
    }
    if (obj instanceof Object && Object.keys(obj).length === 0) {
        return true
    }
    return false
}

/**
 * 清理对象空值属性
 * @param obj
 * @param isClearStr    清理空字符串：''
 * @param isClearObj    清理空对象：[]，{}
 * @returns
 */
export const removeEmptyOfObj = (obj: { [propName: string]: any }, isClearStr: boolean = false, isClearObj: boolean = false) => {
    const temp: { [propName: string]: any } = {}
    Object.keys(obj).forEach((item) => {
        if (!(obj[item] === undefined || obj[item] === null || (isClearStr && obj[item] === '') || (isClearObj && isEmptyObj(obj[item])))) {
            temp[item] = obj[item]
        }
    })
    return temp
}

/**
 * JSON反序列化
 * @param value
 * @returns
 */
export const jsonDecode = (value: string, def: any = undefined): any => {
    try {
        return JSON.parse(value)
    } catch (error) {
        return def === undefined ? value : def
    }
}

/**
 * JSON序列化
 * @param value
 * @returns
 */
export const jsonEncode = (value: any, def: string | undefined = undefined): string => {
    try {
        return JSON.stringify(value)
    } catch (error) {
        return def === undefined ? value : def
    }
}
