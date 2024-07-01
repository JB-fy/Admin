export default {
    name: {
        pay_name: '名称',
        pay_icon: '图标',
        pay_type: '类型',
        pay_config: '配置',
        pay_config_obj: {
            payOfAliAppId: '支付宝-AppID',
            payOfAliPrivateKey: '支付宝-私钥',
            payOfAliPublicKey: '支付宝-公钥',
            payOfAliOpAppId: '支付宝-小程序AppID',

            payOfWxAppId: '微信-AppID',
            payOfWxMchid: '微信-商户ID',
            payOfWxSerialNo: '微信-证书序列号',
            payOfWxApiV3Key: '微信-APIV3密钥',
            payOfWxPrivateKey: '微信-私钥',
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
        pay_config_obj: {
            payOfAliOpAppId: 'JSAPI支付需设置',
        },
        sort: '从大到小排序',
    },
}
