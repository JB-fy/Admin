export default {
    label: {
        sms: '短信',
        email: '邮箱',
        id_card: '实名认证',
        one_click: '一键登录',
        push: '推送',
        vod: '视频点播',
        wx: '微信',

        one_click_of_wx: '微信',
        one_click_of_yidun: '易盾',

        wx_gzh: '公众号',
    },
    name: {
        sms_type: '短信方式',
        sms_of_aliyun: {
            access_key_id: '阿里云-AccessKeyId',
            access_key_secret: '阿里云-AccessKeySecret',
            endpoint: '阿里云-Endpoint',
            sign_name: '阿里云-签名',
            template_code: '阿里云-模板标识',
        },

        email_code: {
            subject: '验证码标题',
            template: '验证码模板',
        },
        email_type: '邮箱方式',
        email_of_common: {
            smtp_host: '通用-SmtpHost',
            smtp_port: '通用-SmtpPort',
            from_email: '通用-邮箱',
            password: '通用-密码',
        },

        id_card_type: '实名认证方式',
        id_card_of_aliyun: {
            url: '阿里云-请求地址',
            appcode: '阿里云-Appcode',
        },

        one_click_of_wx: {
            host: '微信-域名',
            app_id: '微信-AppId',
            secret: '微信-密钥',
        },
        one_click_of_yidun: {
            secret_id: '易盾-SecretId',
            secret_key: '易盾-SecretKey',
            business_id: '易盾-BusinessId',
        },

        push_type: '推送方式',
        push_of_tx: {
            host: '腾讯移动推送-域名',
            access_id_of_android: '腾讯移动推送-AccessID(安卓)',
            secret_key_of_android: '腾讯移动推送-SecretKey(安卓)',
            access_id_of_ios: '腾讯移动推送-AccessID(苹果)',
            secret_key_of_ios: '腾讯移动推送-SecretKey(苹果)',
            access_id_of_mac_os: '腾讯移动推送-AccessID(苹果电脑)',
            secret_key_of_mac_os: '腾讯移动推送-SecretKey(苹果电脑)',
        },

        vod_type: '视频点播方式',
        vod_of_aliyun: {
            access_key_id: '阿里云-AccessKeyId',
            access_key_secret: '阿里云-AccessKeySecret',
            endpoint: '阿里云-Endpoint',
            role_arn: '阿里云-RoleArn',
        },

        wx_gzh: {
            host: '公众号-域名',
            app_id: '公众号-AppId',
            secret: '公众号-密钥',
            token: '公众号-Token',
            encoding_aes_key: '公众号-EncodingAESKey',
        },
    },
    status: {
        sms_type: [{ value: `sms_of_aliyun`, label: '阿里云' }],
        email_type: [{ value: `email_of_common`, label: '通用' }],
        id_card_type: [{ value: `id_card_of_aliyun`, label: '阿里云' }],
        push_type: [{ value: `push_of_tx`, label: '腾讯移动推送' }],
        vod_type: [{ value: `vod_of_aliyun`, label: '阿里云' }],
    },
    tip: {
        email_code: {
            template: '需保证至少拥有一个验证码占位符：' + "{'{'}" + 'code' + "{'}'}" + '',
        },
        email_of_common: {
            password: '注意：如果使用QQ邮箱，此处应填写QQ邮箱的授权码，而不是密码',
        },

        id_card_of_aliyun: {
            url: '购买地址：<a target="_blank" href="https://market.aliyun.com/products/57000002/cmapi014760.html">https://market.aliyun.com/products/57000002/cmapi014760.html</a>（也可购买阿里云市场其它接口，但需修改id_card_of_aliyun.go文件）',
        },

        wx: {
            host: '参考：<a target="_blank" href="https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Interface_field_description.html">https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Interface_field_description.html</a>',
        },

        push_of_tx: {
            host: '参考：<a target="_blank" href="https://cloud.tencent.com/document/product/548/49157">https://cloud.tencent.com/document/product/548/49157</a>',
        },

        vod_of_aliyun: {
            endpoint: '用于生成STS凭证。参考：<a target="_blank" href="https://api.aliyun.com/product/Sts">https://api.aliyun.com/product/Sts</a>',
            role_arn: '用于生成STS凭证',
        },
    },
}
