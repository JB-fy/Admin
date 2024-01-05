export default {
    name: {
        sceneName: '名称',
        sceneCode: '标识',
        sceneConfig: '配置',
        remark: '备注',
        isStop: '停用'
    },
    status: {},
    tip: {
        sceneConfig: 'JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：' + "{'{'}" + '"signType": "算法","signKey": "密钥","expireTime": 过期时间,...' + "{'}'}" + ''
    }
}
