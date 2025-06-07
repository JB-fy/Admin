export default {
    name: {
        menu_name: '名称',
        scene_id: '场景',
        pid: '父级',
        is_leaf: '叶子',
        level: '层级',
        id_path: 'ID路径',
        name_path: '名称路径',
        is_input_menu_icon: '手动填写地址或标识',
        menu_icon: '图标',
        menu_url: '链接',
        extra_data: '额外数据',
        sort: '排序值',
        is_stop: '停用',
    },
    status: {},
    tip: {
        menu_icon: '标识需在前端代码文件my-icon-dynamic.tsx中添加对应标识才能生效，格式：autoicon-' + "{'{'}" + '集合' + "{'}-{'}" + '标识' + "{'}'}" + '，参考：<a target="_blank" href="https://icones.js.org">https://icones.js.org</a>',
        extra_data: 'JSON格式：' + "{'{'}" + '"i18n（国际化设置）": ' + "{'{'}" + '"title": ' + "{'{'}" + '"语言标识":"标题",...' + "{'}'}" + '' + "{'}'}" + '',
        sort: '从大到小排序',
    },
}
