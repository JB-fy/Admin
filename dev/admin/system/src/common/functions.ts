//common.ts：基于当前框架的封装函数。functions.ts：基于原生js的封装函数

//获取语言
export const getBrowserLanguage = (defaultValue: string = 'zh-cn'): string => {
    return (navigator.language || defaultValue).toLowerCase()
}