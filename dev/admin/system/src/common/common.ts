/**
 * 批量导入
 * @param rawImportList 导入列表。（调用import.meta.globEager或import.meta.glob方法返回的数据）
 * @param level 命名层次。特别注意：如果有不同层次文件时，默认以最浅层文件为基准开始命名。正数则表示以父级文件夹（增加相应层数）为基准开始命名；负数则表示以子级文件（减去相应层数）为基准开始命名，意味着将有部分文件不会返回。例如：-1，则最浅层文件将不返回。
 * @param type  类型，默认0。0：一维对象（键名保持原样）；1：一维对象（键名小驼峰法）；2：一维对象（键名大驼峰法）；10：多维对象（键名保持原样）；11：多维对象（键名小驼峰法）；12：多维对象（键名大驼峰法）；
 * @param noUseIndex  当文件名是index.*或者Index.*时，放弃该名称，以父级文件夹命名
 * @returns 
 */
export const batchImport = async (rawImportList: any, level: number = 0, type: number = 0, noUseIndex: boolean = false): Promise<{ [propName: string]: any }> => {
    let importList: { [propName: string]: any } = {}
    let keyArr: string[] = []
    let keyList: string[][] = []
    let importArr: any[] = []
    let levelOfMin: number = 0
    let importOne: { [propName: string]: any } = {}
    for (const path in rawImportList) {
        keyArr = path.slice(0, path.lastIndexOf('.')).split('/')
        keyList.push(keyArr)
        if (typeof rawImportList[path] === 'function') {
            importOne = await rawImportList[path]()
        } else {
            importOne = await rawImportList[path]
        }
        if (importOne.default) {  //有默认值只返回默认值
            importArr.push(importOne.default)
        } else {
            importArr.push(importOne)
        }
        if (keyArr.length < levelOfMin || levelOfMin == 0) {
            levelOfMin = keyArr.length
        }
    }
    const start: number = levelOfMin - level - 1 < 0 ? 0 : levelOfMin - level - 1;    //键名开始的位置
    switch (type) {
        case 0:
            for (const key in keyList) {
                const keyFinal = keyList[key].slice(start).reduce((keyFinalTmp, value) => {
                    return keyFinalTmp + value
                })
                importList[keyFinal] = importArr[key]
            }
            break;
        case 1:
            for (const key in keyList) {
                const keyFinal = keyList[key].slice(start).reduce((keyFinalTmp, value, index) => {
                    if (index == 0) {
                        return keyFinalTmp += value.split(/[\s-_]/).reduce((keyFinalTmp, value, index) => {
                            if (index == 0) {
                                return keyFinalTmp += value.slice(0, 1).toLowerCase() + value.slice(1)
                            }
                            return keyFinalTmp += value.slice(0, 1).toUpperCase() + value.slice(1)
                        }, '')
                    }
                    return keyFinalTmp += value.split(/[\s-_]/).reduce((keyFinalTmp, value) => {
                        return keyFinalTmp + value.slice(0, 1).toUpperCase() + value.slice(1)
                    }, '')
                }, '')
                importList[keyFinal] = importArr[key]
            }
            break;
        case 2:
            for (const key in keyList) {
                const keyFinal = keyList[key].slice(start).reduce((keyFinalTmp, value, index) => {
                    return keyFinalTmp += value.split(/[\s-_]/).reduce((keyFinalTmp, value) => {
                        return keyFinalTmp + value.slice(0, 1).toUpperCase() + value.slice(1)
                    }, '')
                }, '')
                importList[keyFinal] = importArr[key]
            }
            break;
        case 10:
            for (const key in keyList) {
                keyList[key].slice(start).reduce((importTmp, value, index, arr) => {
                    const keyFinal = value

                    if (index == arr.length - 1) {
                        importTmp[keyFinal] = importArr[key]
                    } else {
                        if (!(keyFinal in importTmp)) {
                            importTmp[keyFinal] = {}
                        }
                    }
                    return importTmp[keyFinal]
                }, importList)  //将importList作为importTmp的初始值，当importTmp改变，importList同时也会改变（js对象除非深复制，否则不管多少个变量都是指向同一个内存地址）
            }
            break;
        case 11:
            for (const key in keyList) {
                keyList[key].slice(start).reduce((importTmp, value, index, arr) => {
                    const keyFinal = value.split(/[\s-_]/).reduce((keyFinalTmp, value, index) => {
                        if (index == 0) {
                            return keyFinalTmp += value.slice(0, 1).toLowerCase() + value.slice(1)
                        }
                        return keyFinalTmp += value.slice(0, 1).toUpperCase() + value.slice(1)
                    }, '')

                    if (index == arr.length - 1) {
                        importTmp[keyFinal] = importArr[key]
                    } else {
                        if (!(keyFinal in importTmp)) {
                            importTmp[keyFinal] = {}
                        }
                    }
                    return importTmp[keyFinal]
                }, importList)  //将importList作为importTmp的初始值，当importTmp改变，importList同时也会改变（js对象除非深复制，否则不管多少个变量都是指向同一个内存地址）
            }
            break;
        case 12:
            for (const key in keyList) {
                keyList[key].slice(start).reduce((importTmp, value, index, arr) => {
                    const keyFinal = value.split(/[\s-_]/).reduce((keyFinalTmp, value, index) => {
                        return keyFinalTmp += value.slice(0, 1).toUpperCase() + value.slice(1)
                    }, '')

                    if (index == arr.length - 1) {
                        importTmp[keyFinal] = importArr[key]
                    } else {
                        if (!(keyFinal in importTmp)) {
                            importTmp[keyFinal] = {}
                        }
                    }
                    return importTmp[keyFinal]
                }, importList)  //将importList作为importTmp的初始值，当importTmp改变，importList同时也会改变（js对象除非深复制，否则不管多少个变量都是指向同一个内存地址）
            }
            break;
    }
    return importList
}

import router from '@/router'
//获取当前路由
export const getCurrentRoute = () => {
    return router.currentRoute.value
}