export default {
    name: {
        app_id: 'APP',
        pkg_type: '类型',
        pkg_name: '包名',
        is_input_pkg_file: '手动填写安装包地址',
        pkg_file: '安装包',
        ver_no: '版本号',
        ver_name: '版本名称',
        ver_intro: '版本介绍',
        extra_config: '额外配置',
        extra_config_obj: {
            market_url: '苹果-应用市场地址',
            is_plist_file: '手动填写plist文件地址',
            plist_file: '苹果-plist文件',
        },
        remark: '备注',
        is_force_prev: '强制更新',
        is_stop: '停用',
    },
    status: {
        pkg_type: [
            { value: 0, label: '安卓' },
            { value: 1, label: '苹果' },
            { value: 2, label: 'PC' },
        ],
    },
    tip: {
        extra_config_obj: {
            market_url: '示例：itms-apps://itunes.apple.com/app/idxxxxxxxxxx',
            plist_file: '企业签必须',
        },
        is_force_prev: '注意：只根据前一个版本来设置，与更早之前的版本无关',
    },
}
