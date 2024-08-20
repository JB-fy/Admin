export default {
    name: {
        channel_name: '名称',
        channel_icon: '图标',
        scene_id: '场景',
        pay_id: '支付',
        method: '方法',
        sort: '排序值',
        total_amount: '总额',
        is_stop: '停用',
    },
    status: {
        method: [
            { value: 0, label: 'APP支付' },
            { value: 1, label: 'H5支付' },
            { value: 2, label: '扫码支付' },
            { value: 3, label: '小程序支付' },
        ],
    },
    tip: {
        sort: '从大到小排序',
    },
}
