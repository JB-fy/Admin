/**
 * 请求接口
 * @param apiCode   接口标识。格式：index.encryptStr
 *      用法1：直接写接口地址。'/'要写成'.'
 *      用法2：写apiList内的键名。键名以src/api作为根目录，用'.'拼接组成
 *  一般用法1即满足大部分需求。以下情况使用用法2：
 *      接口不是post请求时
 *      在接口变动时，只需在src/api目录内建立对应的文件，则不用到处修改之前的代码
 *      接口是单路径地址时，如/login，为防止以后接口变动造成麻烦，调用时传login.login
 * @param data  请求参数
 * @param isSuccessTip  成功弹出提示
 * @param isErrorHandle 失败错误处理
 * @returns 
 */
const apiList = await batchImport(import.meta.globEager('@/api/**/*.ts'), 0, 10)    //放外面。这样每次调用都不要重新加载了
export const request = async (apiCode: string, data: { [propName: string]: any } = {}, isSuccessTip: boolean = false, isErrorHandle: boolean = true): Promise<any> => {
    let apiCodeList: string[] = apiCode.split('/')
    switch (apiCodeList[apiCodeList.length - 1]) {
        case 'delete':  //src/api目录内文件不能用delete作函数名。delete为js关键字
            apiCodeList[apiCodeList.length - 1] = 'del'
            break;
    }
    let apiMethod: any = apiList
    for (const value of apiCodeList) {
        if (!(value in apiMethod)) {
            break;
        }
        apiMethod = apiMethod[value]
    }

    try {
        let res
        if (typeof apiMethod === 'function') {
            res = await apiMethod(data)
        } else {
            //未定义接口时，将apiCode转换成接口地址去请求
            apiCodeList = apiCode.split('/')    //重置apiCodeList
            switch (apiCodeList[apiCodeList.length - 1]) {
                case 'save':
                    if (data?.id > 0) {
                        apiCodeList[apiCodeList.length - 1] = 'update'
                    } else {
                        apiCodeList[apiCodeList.length - 1] = 'create'
                    }
                    break;
            }
            res = await http({
                url: '/' + apiCodeList.join('/'),
                method: 'post',
                data: data
            })
        }
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
/* request('index/index', data)
    .then((res) => {
        console.log(res)
    })
    .catch((error) => {
        errorHandle(<Error>error)   //request第四个参数为false时增加，否则已经做过错误处理
    }).finally(()=>{

    })

try {
    await request('index/index', data)
} catch (error) {
    //errorHandle(<Error>error) //request第四个参数为false时增加，否则已经做过错误处理
} */
/*--------使用方式 结束--------*/