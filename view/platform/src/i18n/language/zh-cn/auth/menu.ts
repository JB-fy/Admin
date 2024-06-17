export default {
    name: {
        menu_name: '名称',
        scene_id: '场景',
        pid: '父级',
        level: '层级',
        id_path: '层级路径',
        menu_icon: '图标',
        menu_url: '链接',
        extra_data: '额外数据',
        sort: '排序值',
        is_stop: '停用',
    },
    status: {},
    tip: {
        menu_icon: '常用格式：autoicon-' + "{'{'}" + '集合' + "{'}-{'}" + '标识' + "{'}'}" + '；vant格式：vant-' + "{'{'}" + '标识' + "{'}'}",
        extra_data: 'JSON格式：' + "{'{'}" + '"i18n（国际化设置）": ' + "{'{'}" + '"title": ' + "{'{'}" + '"语言标识":"标题",...' + "{'}'}" + '' + "{'}'}" + '',
        sort: '从大到小排序',
    },
}
