export default {
    label: {
        websiteConfig: '网站',
        uploadConfig: '上传',
        smsConfig: '短信',
        idCardConfig: '实名认证',
    },
    name: {
        userAgreement: '用户协议',
        privacyAgreement: '隐私协议',

        uploadType: '上传方式',
        localUploadUrl: '本地-上传地址',
        localUploadSignKey: '本地-密钥',
        localUploadFileSaveDir: '本地-文件保存目录',
        localUploadFileUrlPrefix: '本地-文件地址前缀',
        aliyunOssHost: '阿里云OSS-域名',
        aliyunOssBucket: '阿里云OSS-Bucket',
        aliyunOssAccessKeyId: '阿里云OSS-AccessKeyId',
        aliyunOssAccessKeySecret: '阿里云OSS-AccessKeySecret',
        aliyunOssRoleArn: '阿里云OSS-RoleArn',
        aliyunOssCallbackUrl: '阿里云OSS-回调地址',

        smsType: '短信方式',
        aliyunSmsAccessKeyId: '阿里云SMS-AccessKeyId',
        aliyunSmsAccessKeySecret: '阿里云SMS-AccessKeySecret',
        aliyunSmsEndpoint: '阿里云SMS-Endpoint',
        aliyunSmsSignName: '阿里云SMS-签名',
        aliyunSmsTemplateCode: '阿里云SMS-模板标识',

        idCardType: '实名认证方式',
        aliyunIdCardHost: '阿里云IdCard-域名',
        aliyunIdCardPath: '阿里云IdCard-请求路径',
        aliyunIdCardAppcode: '阿里云IdCard-Appcode',
    },
    status: {
        uploadType: [
            { value: `local`, label: '本地' },
            { value: `aliyunOss`, label: '阿里云' },
        ],
        smsType: [
            { value: `aliyunSms`, label: '阿里云' },
        ],
        idCardType: [
            { value: `aliyunIdCard`, label: '阿里云' },
        ],
    },
    tip: {
        localUploadFileSaveDir: '根据部署的线上环境设置。一般与nginx中设置的网站对外目录一致',
        localUploadFileUrlPrefix: '根据部署的线上环境设置。与文件保存路径拼接形成文件访问地址',
        aliyunOssHost: '不含Bucket部分',
        aliyunOssRoleArn: '只在APP直传时使用，可不设置',
        aliyunOssCallbackUrl: '设置后开启回调，否则关闭回调',
    },
}