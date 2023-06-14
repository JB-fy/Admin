export default {
    name:{
        sceneName: '场景名称',
        sceneCode: '场景标识',
        sceneConfig: '场景配置',
    },
    tip: {
        //sceneConfig: 'JSON格式。后台配置说明：{"signType": "算法","signKey": "密钥","expireTime": 有效时间,...}',
        sceneConfig: 'JSON格式。后台配置说明：' + "{'{'}" + '"signType": "算法","signKey": "密钥","expireTime": 有效时间,...' + "{'}'}",
    }
}