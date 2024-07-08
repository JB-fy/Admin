export default {
    name: {
        upload_type: '类型',
        upload_config: '配置',
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
    },
}
