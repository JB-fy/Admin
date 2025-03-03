export default {
    name: {
        upload_type: '类型',
        upload_config: '配置',
        upload_config_0: {
            signKey: '本地-密钥',
            url: '本地-上传地址',
            fileSaveDir: '本地-保存目录',
            isCluster: '本地-集群服务',
            serverList_label: '本地-服务器列表',
            serverList: {
                ip: '外网IP',
                host: '域名',
            },
            isSameServer: '本地-单次多文件上传相同服务器',
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
            url: '默认：域名/upload/upload。可填写完整的http地址指定上传服务器',
            fileSaveDir: '默认：../public/。需根据部署环境设置，一般填写服务启动文件所在目录与域名对外目录的相对路径',
            serverList: '集群服务时，默认文件地址：外网IP:端口/文件路径（需开放端口）。<br>如果不想暴露IP和端口，各服务器可设置不同域名，并将外网IP和域名设置到此列表中，未设置的服务器还是返回默认地址',
            isSameServer: '集群服务时，启用后上传地址的域名将被替换成：外网IP:端口。如设置服务器列表，则替换成各服务器域名',
        },
        upload_config_1: {
            host: '不含Bucket部分',
            endpoint: 'APP直传需设置，用于生成STS凭证。参考：<a target="_blank" href="https://api.aliyun.com/product/Sts">https://api.aliyun.com/product/Sts</a>',
            roleArn: 'APP直传需设置，用于生成STS凭证',
        },
    },
}
