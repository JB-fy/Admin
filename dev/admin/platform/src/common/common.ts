import i18n from '@/i18n'


/**
 * 请求接口
 * @param apiCode   接口标识。apiList内的键用'.'拼接组成。例如：login.encryptStr
 * @param data  请求参数
 * @param isSuccessTip  成功弹出提示
 * @param isErrorHandle 失败错误处理
 * @returns 
 */
const apiList = await batchImport(import.meta.globEager('@/api/**/*.ts'), 0, 10)
export const request = async (apiCode: string, data: object = {}, isSuccessTip: boolean = false, isErrorHandle: boolean = true): Promise<any> => {
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
            throw new Error(i18n.global.t('error.apiFunctionNoFind'))
        }
        const res = await apiMethod(data)
        if (isSuccessTip) {
            ElMessage.success(res.msg)
        }
        return res
    } catch (error) {
        if (isErrorHandle) {
            errorHandle(<Error>error)
        }
        throw error
    }
}
/*--------使用方式 开始--------*/
/* request('index.index', data)
    .then((res) => {
        console.log(res)
    })
    .catch((error) => {
        errorHandle(<Error>error)   //request第四个参数为false时增加，否则已经做过错误处理
    }).finally(()=>{

    })

try {
    await request('index.index', data)
} catch (error) {
    //errorHandle(<Error>error) //request第四个参数为false时增加，否则已经做过错误处理
} */
/*--------使用方式 结束--------*/