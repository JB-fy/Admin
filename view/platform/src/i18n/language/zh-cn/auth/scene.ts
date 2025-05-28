export default {
    name: {
        scene_id: '场景ID',
        scene_name: '名称',
        token_config: {
            token_type: 'Token方式',
            expire_time: '过期时间',
            active_time: '失活时间',
            is_ip: '验证IP',
            is_unique: 'Token唯一',
        },
        token_config_0: {
            sign_type: 'JWT-加密方式',
            key: 'JWT-密钥',
            private_key: 'JWT-私钥',
            public_key: 'JWT-公钥',
        },
        sign_config: {
            sign_type: '签名方式',
        },
        sign_config_0: {
            method: '通用-加密方式',
            key: '通用-密钥',
            key_name: '通用-密钥名称',
            key_sep: '通用-多字段拼接符',
            val_sep: '通用-单字段拼接符',
        },
        scene_config: '配置',
        remark: '备注',
        is_stop: '停用',
    },
    status: {
        token_config: {
            token_type: [
                { value: -1, label: '不用' },
                { value: 0, label: 'JWT' },
            ],
        },
        token_config_0: {
            sign_type: [
                { value: 'HS256', label: 'HS256' },
                { value: 'HS384', label: 'HS384' },
                { value: 'HS512', label: 'HS512' },
                { value: 'RS256', label: 'RS256' },
                { value: 'RS384', label: 'RS384' },
                { value: 'RS512', label: 'RS512' },
                { value: 'ES256', label: 'ES256' },
                { value: 'ES384', label: 'ES384' },
                { value: 'ES512', label: 'ES512' },
            ],
        },
        sign_config: {
            sign_type: [
                { value: -1, label: '不用' },
                { value: 0, label: '通用' },
            ],
        },
        sign_config_0: {
            method: [
                { value: 'md5', label: 'MD5' },
                { value: 'hmac-md5', label: 'HMAC-MD5' },
                { value: 'hmac-sha1', label: 'HMAC-SHA1' },
                { value: 'hmac-sha256', label: 'HMAC-SHA256' },
                { value: 'hmac-sha512', label: 'HMAC-SHA512' },
            ],
        },
    },
    tip: {
        scene_config: 'JSON格式，根据场景设置',
        token_config: {
            expire_time: '多少秒后Token失效',
            active_time: '大于0生效，防止长时间无操作（人离开）时，被他人趁机而入（一段秒数内Token未使用，判定失活）',
            is_ip: '开启后，可防止Token被盗用（验证使用Token时的IP与生成Token时的IP是否一致）',
            is_unique: '开启后，可限制用户多地、多设备登录（同时只会有一个Token有效，生成新Token时，旧Token失效）',
        },
        sign_config_0: {
            key_name: '不为空时，会在签名的字符串后面追加密钥名称和密钥',
            key_sep: '多个字段之间的拼接符',
            val_sep: '单个字段名与值的拼接符',
        },
    },
}
