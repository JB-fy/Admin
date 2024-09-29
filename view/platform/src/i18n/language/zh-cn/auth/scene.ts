export default {
    name: {
        scene_name: '名称',
        scene_code: '标识',
        scene_config: '配置',
        token_config: {
            token_type: 'Token方式',
            expire_time: '过期时间',
            is_unique: 'Token唯一',
        },
        token_config_0: {
            sign_type: 'JWT-加密方式',
            sign_key: 'JWT-密钥',
        },
        remark: '备注',
        is_stop: '停用',
    },
    status: {
        token_config: {
            token_type: [
                { value: -1, label: '不生成' },
                { value: 0, label: 'JWT' },
            ],
        },
        token_config_0: {
            sign_type: [
                { value: 'HS256', label: 'HS256' },
                { value: 'HS384', label: 'HS384' },
                { value: 'HS512', label: 'HS512' },
            ],
        },
    },
    tip: {
        scene_config: 'JSON格式，根据场景设置',
        token_config: {
            expire_time: '多少秒后token失效',
            is_unique: '开启后，可限制用户多地，多设备登录，因同时只会有一个token有效（新token生成时，旧token会失效）',
        },
    },
}
