export default {
    label: {
        sms: '短信',
        email: '邮箱',
        idCard: '实名认证',
        oneClick: '一键登录',
        push: '推送',
        vod: '视频点播',
        wx: '微信',

        oneClickOfWx: '微信',
        oneClickOfYidun: '易盾',

        wxGzh: '公众号',
    },
    name: {
        smsType: '短信方式',
        smsOfAliyun: {
            accessKeyId: '阿里云-AccessKeyId',
            accessKeySecret: '阿里云-AccessKeySecret',
            endpoint: '阿里云-Endpoint',
            signName: '阿里云-签名',
            templateCode: '阿里云-模板标识',
        },

        emailCode: {
            subject: '验证码标题',
            template: '验证码模板',
        },
        emailType: '邮箱方式',
        emailOfCommon: {
            smtp_host: '通用-SmtpHost',
            smtp_port: '通用-SmtpPort',
            from_email: '通用-邮箱',
            password: '通用-密码',
        },

        idCardType: '实名认证方式',
        idCardOfAliyun: {
            host: '阿里云-域名',
            path: '阿里云-请求路径',
            appcode: '阿里云-Appcode',
        },

        oneClickOfWx: {
            host: '微信-域名',
            appId: '微信-AppId',
            secret: '微信-密钥',
        },
        oneClickOfYidun: {
            secretId: '易盾-SecretId',
            secretKey: '易盾-SecretKey',
            businessId: '易盾-BusinessId',
        },

        pushType: '推送方式',
        pushOfTx: {
            host: '腾讯移动推送-域名',
            accessIDOfAndroid: '腾讯移动推送-AccessID(安卓)',
            secretKeyOfAndroid: '腾讯移动推送-SecretKey(安卓)',
            accessIDOfIos: '腾讯移动推送-AccessID(苹果)',
            secretKeyOfIos: '腾讯移动推送-SecretKey(苹果)',
            accessIDOfMacOS: '腾讯移动推送-AccessID(苹果电脑)',
            secretKeyOfMacOS: '腾讯移动推送-SecretKey(苹果电脑)',
        },

        vodType: '视频点播方式',
        vodOfAliyun: {
            accessKeyId: '阿里云-AccessKeyId',
            accessKeySecret: '阿里云-AccessKeySecret',
            endpoint: '阿里云-Endpoint',
            roleArn: '阿里云-RoleArn',
        },

        wxGzh: {
            host: '公众号-域名',
            appId: '公众号-AppId',
            secret: '公众号-密钥',
            token: '公众号-Token',
            encodingAESKey: '公众号-EncodingAESKey',
        },
    },
    status: {
        smsType: [{ value: `smsOfAliyun`, label: '阿里云' }],
        emailType: [{ value: `emailOfCommon`, label: '通用' }],
        idCardType: [{ value: `idCardOfAliyun`, label: '阿里云' }],
        pushType: [{ value: `pushOfTx`, label: '腾讯移动推送' }],
        vodType: [{ value: `vodOfAliyun`, label: '阿里云' }],
    },
    tip: {
        emailCode: {
            template: '需保证至少拥有一个验证码占位符：' + "{'{'}" + 'code' + "{'}'}" + '',
        },
        emailOfCommon: {
            password: '注意：如果使用QQ邮箱，此处应填写QQ邮箱的授权码，而不是密码',
        },

        idCardOfAliyun: {
            host: '购买地址：<a target="_blank" href="https://market.aliyun.com/products/57000002/cmapi014760.html">https://market.aliyun.com/products/57000002/cmapi014760.html</a>（也可购买阿里云市场其它接口，但需修改id_card_of_aliyun.go文件）',
        },

        oneClickOfWx: {
            host: '参考：<a target="_blank" href="https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Interface_field_description.html">https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Interface_field_description.html</a>',
        },

        pushOfTx: {
            host: '参考：<a target="_blank" href="https://cloud.tencent.com/document/product/548/49157">https://cloud.tencent.com/document/product/548/49157</a>',
        },

        vodOfAliyun: {
            endpoint: '用于生成STS凭证。参考：<a target="_blank" href="https://api.aliyun.com/product/Sts">https://api.aliyun.com/product/Sts</a>',
            roleArn: '用于生成STS凭证',
        },
    },
}
