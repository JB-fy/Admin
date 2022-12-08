/**
 * 请求接口
 * @param apiCode   接口标识。格式：index.encryptStr
 *                      用法1：直接写接口地址。'/'要写成'.'
 *                      用法2：写apiList内的键名。键名以src/api作为根目录，用'.'拼接组成
 *                  建议：一般用法1即可。在接口改动需要更改地址时，则在src/api建立之前对应的接口，则不用到处修改之前的代码
 * @param data  请求参数
 * @param isSuccessTip  成功弹出提示
 * @param isErrorHandle 失败错误处理
 * @returns 
 */
const apiList = await batchImport(import.meta.globEager('@/api/**/*.ts'), 0, 10)
export const request = async (apiCode: string, data: { [propName: string]: any } = {}, isSuccessTip: boolean = false, isErrorHandle: boolean = true): Promise<any> => {
    //const apiList = batchImport(import.meta.globEager('@/api/**/*.ts')) //放外面去。这样每次调用都不要重新加载了
    const apiCodeList: string[] = apiCode.split('.')
    /* switch (apiCodeList[apiCodeList.length - 1]) {
        case 'delete':  //src/api内的文件不能用delete作函数名
            apiCodeList[apiCodeList.length - 1] = 'del'
            break;
    } */
    console.log(apiList)
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
            switch (apiCodeList[apiCodeList.length - 1]) {
                case 'save':
                    if (data?.id > 0) {
                        apiCodeList[apiCodeList.length - 1] = 'update'
                    } else {
                        apiCodeList[apiCodeList.length - 1] = 'create'
                    }
                    break;
                case 'del':
                    apiCodeList[apiCodeList.length - 1] = 'delete'
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