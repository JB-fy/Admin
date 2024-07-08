export default {
    name: {
        upload_type: '类型',
        upload_config: '配置',
        upload_config_0: {
            url: '本地-上传地址',
            signKey: '本地-密钥',
            fileSaveDir: '本地-保存目录',
            fileUrlPrefix: '本地-文件地址前缀',
        },
        upload_config_1: {
            host: '阿里云OSS-域名',
            bucket: '阿里云OSS-Bucket',
            accessKeyId: '阿里云OSS-AccessKeyId',
            accessKeySecret: '阿里云OSS-AccessKeySecret',
            endpoint: '阿里云OSS-Endpoint',
            roleArn: '阿里云OSS-RoleArn',
            isNotify: '阿里云OSS-回调',
        },
        remark: '备注',
        is_default: '默认',
    },
    status: {
        upload_type: [
            { value: 0, label: '本地' },
            { value: 1, label: '阿里云OSS' },
        ],
    },
    tip: {
        upload_config: '根据upload_type类型设置',
        upload_config_0: {
            fileSaveDir: '根据部署的线上环境设置。一般与nginx中设置的网站对外目录一致',
            fileUrlPrefix: '根据部署的线上环境设置。与文件保存路径拼接形成文件访问地址',
        },
        upload_config_1: {
            host: '不含Bucket部分',
            endpoint: 'APP直传需设置，用于生成STS凭证。参考：<a target="_blank" href="https://api.aliyun.com/product/Sts">https://api.aliyun.com/product/Sts</a>',
            roleArn: 'APP直传需设置，用于生成STS凭证',
        },
    },
}
