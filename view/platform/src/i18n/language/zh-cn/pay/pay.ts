export default {
    name: {
        pay_name: '名称',
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
        remark: '备注',
    },
    status: {
        pay_type: [
            { value: 0, label: '支付宝' },
            { value: 1, label: '微信' },
        ],
    },
    tip: {
        pay_config: '根据pay_type类型设置',
        pay_config_0: {
            opAppId: 'JSAPI支付需设置',
        },
    },
}
