export default {
    name: {
        upload_type: '类型',
        upload_config: '配置',
        upload_config_obj: {
            uploadOfLocalUrl: '本地-上传地址',
            uploadOfLocalSignKey: '本地-密钥',
            uploadOfLocalFileSaveDir: '本地-保存目录',
            uploadOfLocalFileUrlPrefix: '本地-文件地址前缀',

            uploadOfAliyunOssHost: '阿里云OSS-域名',
            uploadOfAliyunOssBucket: '阿里云OSS-Bucket',
            uploadOfAliyunOssAccessKeyId: '阿里云OSS-AccessKeyId',
            uploadOfAliyunOssAccessKeySecret: '阿里云OSS-AccessKeySecret',
            uploadOfAliyunOssEndpoint: '阿里云OSS-Endpoint',
            uploadOfAliyunOssRoleArn: '阿里云OSS-RoleArn',
            uploadOfAliyunOssIsNotify: '阿里云OSS-回调',
        },
        remark: '备注',
        is_default: '默认',
        is_stop: '停用',
    },
    status: {
        upload_type: [
            { value: 0, label: '本地' },
            { value: 1, label: '阿里云OSS' },
        ],
    },
    tip: {
        upload_config: '根据upload_type类型设置',
        upload_config_obj: {
            uploadOfLocalFileSaveDir: '根据部署的线上环境设置。一般与nginx中设置的网站对外目录一致',
            uploadOfLocalFileUrlPrefix: '根据部署的线上环境设置。与文件保存路径拼接形成文件访问地址',
            uploadOfAliyunOssHost: '不含Bucket部分',
            uploadOfAliyunOssEndpoint: 'APP直传需设置，用于生成STS凭证。参考：<a target="_blank" href="https://api.aliyun.com/product/Sts">https://api.aliyun.com/product/Sts</a>',
            uploadOfAliyunOssRoleArn: 'APP直传需设置，用于生成STS凭证',
        },
    },
}
