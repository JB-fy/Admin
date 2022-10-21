/**
 * 批量导入
 * @param rawImportList 导入列表。（调用import.meta.globEager或import.meta.glob方法返回的数据）
 * @param level 命名层次。特别注意：如果有不同层次文件时，默认以最浅层文件为基准开始命名。正数则表示以父级文件夹（增加相应层数）为基准开始命名；负数则表示以子级文件（减去相应层数）为基准开始命名，意味着将有部分文件不会返回。例如：-1，则最浅层文件将不返回。
 * @param type  类型，默认0。0：一维对象（键名保持原样）；1：一维对象（键名小驼峰法）；2：一维对象（键名大驼峰法）；10：多维对象（键名保持原样）；11：多维对象（键名小驼峰法）；12：多维对象（键名大驼峰法）；
 * @param noUseIndex  当文件名是index.*或者Index.*时，放弃该名称，以父级文件夹命名
 * @returns 
 */
//export const batchImport = async (rawImportList: any, level: number = 0, type: number = 0, noUseIndex: boolean = false) => {
export const batchImport = (rawImportList: any, level: number = 0, type: number = 0, noUseIndex: boolean = false) => {
    let importList: { [propName: string]: any } = {}
    let keyArr: string[] = []
    let keyList: string[][] = []
    let importArr: any[] = []
    let levelOfMin: number = 0
    for (const path in rawImportList) {
        keyArr = path.slice(0, path.lastIndexOf('.')).split('/')
        keyList.push(keyArr)
        /* if (typeof rawImportList[path] === 'function') {
            importArr.push((await rawImportList[path]()).default)
        } else {
            importArr.push(rawImportList[path].default)
        } */
        if (typeof rawImportList[path] === 'object' && rawImportList[path].default) {
            importArr.push(rawImportList[path].default)
        } else {
            importArr.push(rawImportList[path])
        }
        if (keyArr.length < levelOfMin || levelOfMin == 0) {
            levelOfMin = keyArr.length
        }
    }
    switch (type) {
        case 0:
        case 1:
        case 2:
            for (const key in keyList) {
                let keyTmp: string = ''
                let iTmp: number = levelOfMin - level - 1 < 0 ? 0 : levelOfMin - level - 1; //循环参数i的初始值，不能小于0
                for (let i = iTmp; i < keyList[key].length; i++) {
                    switch (type) {
                        case 0:
                            keyTmp = keyTmp + keyList[key][i]
                            break;
                        case 1:
                            const keyArrTmp1 = keyList[key][i].split(/[\s-_]/)
                            if (i == iTmp) {
                                for (const key1 in keyArrTmp1) {
                                    if (key1 == '0') {
                                        keyTmp = keyTmp + keyArrTmp1[key1].slice(0, 1).toLowerCase() + keyArrTmp1[key1].slice(1)
                                    } else {
                                        keyTmp = keyTmp + keyArrTmp1[key1].slice(0, 1).toUpperCase() + keyArrTmp1[key1].slice(1)
                                    }
                                }
                            } else {
                                for (const key1 in keyArrTmp1) {
                                    keyTmp = keyTmp + keyArrTmp1[key1].slice(0, 1).toUpperCase() + keyArrTmp1[key1].slice(1)
                                }
                            }
                            break;
                        case 2:
                            const keyArrTmp2 = keyList[key][i].split(/[\s-_]/)
                            for (const key2 in keyArrTmp2) {
                                keyTmp = keyTmp + keyArrTmp2[key2].slice(0, 1).toUpperCase() + keyArrTmp2[key2].slice(1)
                            }
                            break;
                    }
                }
                importList[keyTmp] = importArr[key]
            }
            break;
        case 10:
        case 11:
        case 12:
            for (const key in keyList) {
                let importTmp: any = {} //利用js对象共用一个内存地址，改变这个也会影响原来的对象
                let keyTmp: string = ''
                let iTmp: number = levelOfMin - level - 1 < 0 ? 0 : levelOfMin - level - 1; //循环参数i的初始值，不能小于0
                for (let i = iTmp; i < keyList[key].length; i++) {
                    switch (type) {
                        case 10:
                            keyTmp = keyList[key][i]
                            break;
                        case 11:
                            const keyArrTmp11: string[] = keyList[key][i].split(/[\s-_]/)
                            for (const key11 in keyArrTmp11) {
                                if (key11 == '0') {
                                    keyTmp = keyTmp + keyArrTmp11[key11].slice(0, 1).toLowerCase() + keyArrTmp11[key11].slice(1)
                                } else {
                                    keyTmp = keyTmp + keyArrTmp11[key11].slice(0, 1).toUpperCase() + keyArrTmp11[key11].slice(1)
                                }
                            }
                            break;
                        case 12:
                            const keyArrTmp12: string[] = keyList[key][i].split(/[\s-_]/)
                            for (const key12 in keyArrTmp12) {
                                keyTmp = keyTmp + keyArrTmp12[key12].slice(0, 1).toUpperCase() + keyArrTmp12[key12].slice(1)
                            }
                            break;
                    }
                    if (i == iTmp) {
                        if (i == keyList[key].length - 1) {
                            importList[keyTmp] = importArr[key]
                        } else {
                            importList[keyTmp] = {
                                ...importList[keyTmp]
                            }
                            importTmp = importList[keyTmp]; //关键代码。利用js对象共用一个内存地址，改变这个也会影响原来的对象
                        }
                    } else {
                        if (i == keyList[key].length - 1) {
                            importTmp[keyTmp] = importArr[key]
                        } else {
                            importTmp[keyTmp] = {
                                ...importTmp[keyTmp]
                            }
                            importTmp = importTmp[keyTmp]
                        }
                    }
                }
            }
            break;
    }
    return importList
}

console.log(batchImport(import.meta.globEager('@/basic/*.ts')))

import router from '@/router'
/**
 * 获取当前路由路径（需在路由稳定时使用。路由切换时不能使用，此时得到的是上一个页面的路由路径，如：浏览器刷新页面；各路由对应的.vue文件内在export default之前调用。）
 * @returns 
 */
export const getCurrentPath = (): string => {
    //return useRouter().currentRoute.value.path
    return router.currentRoute.value.path
}