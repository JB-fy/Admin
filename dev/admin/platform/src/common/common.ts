/**
 * 请求接口
 * @param apiCode   接口标识。格式：index/encryptStr
 *      用法1：直接写接口地址
 *      用法2：写apiList内的键名。键名以src/api作为根目录，用'.'拼接组成
 *  一般用法1即满足大部分需求。需要特殊处理才使用用法2：
 *      接口不是post请求时
 *      在接口变动时，只需在src/api目录内建立对应的文件，则不用到处修改之前的代码
 * @param data  请求参数
 * @param isSuccessTip  成功弹出提示
 * @param isErrorHandle 失败错误处理
 * @returns 
 */
const apiList = batchImport(import.meta.globEager('@/api/**/*.ts'), 0, 10)    //放外面。这样每次调用都不要重新加载了
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
            res = await http({
                url: apiCode,
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
/* request('index/index', data).then((res) => {
    console.log(res)
}).catch((error) => {
    errorHandle(<Error>error)   //request第四个参数为false时增加，否则已经做过错误处理
}).finally(() => {
})

try {
    const res = await request('index/index', data)
    console.log(res)
} catch (error) {
    //errorHandle(<Error>error) //request第四个参数为false时增加，否则已经做过错误处理
} finally {
} */
/*--------使用方式 结束--------*/

/**
 * 导出excel
 * @param sheetList 
 * @param fileName 
 */
//import * as XLSX from 'xlsx'
import { utils, writeFile } from 'xlsx'
export const exportExcel = (sheetList: { data: any[][] | { [propName: string]: any }[], sheetName?: string }[], fileName: string = 'excel.xlsx') => {
    const workbook = utils.book_new()   //生成工作簿
    sheetList.forEach((item, index) => {
        let sheet
        if (item.data.length > 0 && Array.isArray(item.data[0])) {
            //生成工作表。格式：[[表头1,...],[数据1,...],...]。示例：[["周一", "周二"],["语文", "数学"]]
            sheet = utils.aoa_to_sheet(<any[][]>item.data)
        } else {
            //生成工作表。格式：[{"表头1":"数据1",...},...]。示例：[{周一: '语文',周二: '数学'}]
            sheet = utils.json_to_sheet(<{ [propName: string]: any }[]>item.data)
        }
        utils.book_append_sheet(workbook, sheet, item.sheetName ?? 'sheet' + (index + 1))   //工作簿中添加工作表
    })
    writeFile(workbook, fileName);  //输出工作表，由文件名决定的输出格式
}