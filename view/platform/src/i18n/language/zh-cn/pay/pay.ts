export default {
    name: {
        pay_name: '名称',
        pay_icon: '图标',
        pay_type: '类型',
        pay_config: '配置',
        pay_config_0: {
            appId: '支付宝-AppID',
            privateKey: '支付宝-私钥',
            publicKey: '支付宝-公钥',
            opAppId: '支付宝-小程序AppID',
        },
        pay_config_1: {
            appId: '微信-AppID',
            mchid: '微信-商户ID',
            serialNo: '微信-证书序列号',
            apiV3Key: '微信-APIV3密钥',
            privateKey: '微信-私钥',
        },
        pay_rate: '费率',
        total_amount: '总额',
        balance: '余额',
        sort: '排序值',
        remark: '备注',
        pay_scene_arr: '支付场景',
        is_stop: '停用',
    },
    status: {
        pay_type: [
            { value: 0, label: '支付宝' },
            { value: 1, label: '微信' },
        ],
        pay_scene_arr: [
            { value: 0, label: 'APP' },
            { value: 1, label: 'H5' },
            { value: 2, label: '扫码' },
            { value: 10, label: '微信小程序' },
            { value: 11, label: '微信公众号' },
            { value: 20, label: '支付宝小程序' },
        ],
    },
    tip: {
        pay_config: '根据pay_type类型设置',
        pay_config_0: {
            opAppId: 'JSAPI支付需设置',
        },
        sort: '从大到小排序',
    },
}
