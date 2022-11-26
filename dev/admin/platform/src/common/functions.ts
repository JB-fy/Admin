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
 * 对象的非空属性组成一个新对象返回
 * @param obj 
 * @returns 
 */
export const removeObjectNullValue = (obj: { [propName: string]: any }) => {
    const temp: { [propName: string]: any } = {}
    Object.keys(obj).forEach(item => {
        if (!(obj[item] === '' ||
            obj[item] === undefined ||
            obj[item] === null ||
            obj[item] === 'null')
        ) {
            temp[item] = obj[item]
        }
    })
    return temp
}