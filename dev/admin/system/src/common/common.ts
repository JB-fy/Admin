//获取当前路由
import router from '@/router'
export const getCurrentRoute = () => {
    return router.currentRoute.value
}

//翻译字符串
import i18n from '@/i18n'
export const getI18n = (): any => {
    return i18n.global
}
export const trans = (str: string): any => {
    return (<any>i18n).global.tm(str)
}

/**
 * 请求接口
 * @param apiCode   接口标识。apiList内的键用'.'拼接组成。例如：login.encryptStr
 * @param data  请求参数
 * @param isErrorHandle 错误处理，默认true。传false则会抛出错误，可在调用处捕获错误进行特殊处理。以下几种情况需要传false：调用接口报错需要特殊处理；多接口调用时，有接口报错，后续接口不再请求。（多接口请求也可以传true，用返回参数是否false判断是否进行后续接口请求）
 * @returns 
 */
const apiList = await batchImport(import.meta.globEager('@/api/**/*.ts'))
//export const request = async (apiCode: string, data?: {}, isErrorHandle: boolean = true): Promise<{ [propName: string]: any } | false | void> => {
export const request = async (apiCode: string, data?: {}, isErrorHandle: boolean = true): Promise<any> => {
    //const apiList = batchImport(import.meta.globEager('@/api/**/*.ts')) //放外面去。这样每次调用都不要重新加载了
    let apiMethod: any = apiList;
    for (const value of apiCode.split('.')) {
        if (!(value in apiMethod)) {
            break;
        }
        apiMethod = apiMethod[value]
    }

    try {
        if (typeof apiMethod !== 'function') {
            throw new Error(trans('error.apiFunctionNoFind'))
        }
        return await apiMethod(data)
    } catch (error) {
        if (isErrorHandle) {
            errorHandle(<Error>error)
            return false
        } else {
            throw error
        }
    }
}
/*--------使用方式 开始--------*/
/* request('index.index', data)

request('index.index', undefined, false)
    .then((data) => {
        console.log(data)
    })
    .catch((error) => {
        errorHandle(<Error>error)
    })

try {
    await request('index.index', data, false)
} catch (error) {
    errorHandle(<Error>error)
} */
/*--------使用方式 结束--------*/