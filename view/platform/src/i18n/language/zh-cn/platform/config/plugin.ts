export default {
    label: {
        upload: '上传',
        pay: '支付',
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
            accessKeyId: '阿里云SMS-AccessKeyId',
            accessKeySecret: '阿里云SMS-AccessKeySecret',
            endpoint: '阿里云SMS-Endpoint',
            signName: '阿里云SMS-签名',
            templateCode: '阿里云SMS-模板标识',
        },

        emailCodeSubject: '验证码标题',
        emailCodeTemplate: '验证码模板',
        emailType: '邮箱方式',
        emailOfCommon: {
            smtpHost: '通用-SmtpHost',
            smtpPort: '通用-SmtpPort',
            fromEmail: '通用-邮箱',
            password: '通用-密码',
        },

        idCardType: '实名认证方式',
        idCardOfAliyunHost: '阿里云IdCard-域名',
        idCardOfAliyunPath: '阿里云IdCard-请求路径',
        idCardOfAliyunAppcode: '阿里云IdCard-Appcode',

        oneClickOfWxHost: '微信-域名',
        oneClickOfWxAppId: '微信-AppId',
        oneClickOfWxSecret: '微信-密钥',

        oneClickOfYidunSecretId: '易盾-SecretId',
        oneClickOfYidunSecretKey: '易盾-SecretKey',
        oneClickOfYidunBusinessId: '易盾-BusinessId',

        pushType: '推送方式',
        pushOfTxHost: '腾讯移动推送-域名',
        pushOfTxAndroidAccessID: '腾讯移动推送-AccessID(安卓)',
        pushOfTxAndroidSecretKey: '腾讯移动推送-SecretKey(安卓)',
        pushOfTxIosAccessID: '腾讯移动推送-AccessID(苹果)',
        pushOfTxIosSecretKey: '腾讯移动推送-SecretKey(苹果)',
        pushOfTxMacOSAccessID: '腾讯移动推送-AccessID(苹果电脑)',
        pushOfTxMacOSSecretKey: '腾讯移动推送-SecretKey(苹果电脑)',

        vodType: '视频点播方式',
        vodOfAliyunAccessKeyId: '阿里云VOD-AccessKeyId',
        vodOfAliyunAccessKeySecret: '阿里云VOD-AccessKeySecret',
        vodOfAliyunEndpoint: '阿里云VOD-Endpoint',
        vodOfAliyunRoleArn: '阿里云VOD-RoleArn',

        wxGzhHost: '域名',
        wxGzhAppId: 'AppId',
        wxGzhSecret: '密钥',
        wxGzhToken: 'Token',
        wxGzhEncodingAESKey: 'EncodingAESKey',
    },
    status: {
        smsType: [{ value: `smsOfAliyun`, label: '阿里云' }],
        emailType: [{ value: `emailOfCommon`, label: '通用' }],
        idCardType: [{ value: `idCardOfAliyun`, label: '阿里云' }],
        pushType: [{ value: `pushOfTx`, label: '腾讯移动推送' }],
        vodType: [{ value: `vodOfAliyun`, label: '阿里云' }],
    },
    tip: {
        emailOfCommon: {
            password: '注意：如果使用的是QQ邮箱，则此处应填写QQ邮箱的授权码，而不是密码',
        },
        emailCodeTemplate: '需保证至少拥有一个验证码占位符：' + "{'{'}" + 'code' + "{'}'}" + '',

        idCardOfAliyunHost: '购买地址：<a target="_blank" href="https://market.aliyun.com/products/57000002/cmapi014760.html">https://market.aliyun.com/products/57000002/cmapi014760.html</a>（购买其它接口，只需对代码文件做下简单修改即可）',

        wxHost: '参考：<a target="_blank" href="https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Interface_field_description.html">https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Interface_field_description.html</a>',

        pushOfTxHost: '参考：<a target="_blank" href="https://cloud.tencent.com/document/product/548/49157">https://cloud.tencent.com/document/product/548/49157</a>',

        vodOfAliyunEndpoint: '用于生成STS凭证。请参考：<a target="_blank" href="https://api.aliyun.com/product/Sts">https://api.aliyun.com/product/Sts</a>',
        vodOfAliyunRoleArn: '用于生成STS凭证',
    },
}
