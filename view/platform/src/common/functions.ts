//common.ts：基于当前框架的封装函数。functions.ts：基于原生js的封装函数

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
    if (Array.prototype.isPrototypeOf(obj) && obj.length === 0) {
        return true
    }
    if (Object.prototype.isPrototypeOf(obj) && Object.keys(obj).length === 0) {
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
