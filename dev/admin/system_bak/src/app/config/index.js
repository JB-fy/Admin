const allConfig = getImportComponents(require.context('@/app/config/autoload', false, /\.js$/)) //放外面，不用每次调用getConfig都要加载这个目录
/**
 * 获取配置参数
 * @param {*} key   键名。以'.'分隔，格式：文件名.key1.key2...
 * @param {*} defaultValue  默认值
 * @returns 
 */
export function config(key, defaultValue = null) {
    let keyArr = key.split('.')
    let value = allConfig
    for (let i = 0; i < keyArr.length; i++) {
        if (keyArr[i] in value) {
            value = value[keyArr[i]]
        } else {
            return defaultValue
        }
    }
    return value
}