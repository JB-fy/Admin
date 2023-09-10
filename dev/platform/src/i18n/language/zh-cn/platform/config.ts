export default {
    name: {
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
    },
    status: {
        uploadType: [
            { value: `local`, label: '本地' },
            { value: `aliyunOss`, label: '阿里云' },
        ],
    },
    tip: {
        localUploadFileSaveDir: '根据部署的线上环境设置。一般与nginx中设置的网站对外目录一致',
        localUploadFileUrlPrefix: '根据部署的线上环境设置。与文件保存路径拼接形成文件访问地址',
        aliyunOssHost: '不含Bucket部分',
        aliyunOssRoleArn: '只在APP直传时使用，可不设置',
        aliyunOssCallbackUrl: '设置后开启回调，否则关闭回调',
    },
    label: {
        uploadConfig: '上传',
    },
}