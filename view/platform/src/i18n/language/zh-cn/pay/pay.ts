export default {
    name: {
        pay_name: '名称',
        pay_type: '类型',
        pay_config: '配置',
        pay_config_0: {
            app_id: '支付宝-AppID',
            private_key: '支付宝-私钥',
            public_key: '支付宝-公钥',
            op_app_id: '支付宝-小程序AppID',
        },
        pay_config_1: {
            app_id: '微信-AppID',
            mch_id: '微信-商户ID',
            serial_no: '微信-证书序列号',
            api_v3_key: '微信-APIV3密钥',
            private_key: '微信-私钥',
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
        pay_config: 'JSON格式，根据类型设置',
        pay_config_0: {
            op_app_id: 'JSAPI支付需设置',
        },
    },
}
