export default {
    name: {
        menuName: '菜单名称',
        menuUrl: '菜单链接',
        menuIcon: '菜单图标',
        idPath: '层级路径',
        sceneId: '所属场景',
    },
    tip: {
        // menuIcon: 常用格式：Autoicon{集合}{标识}；Vant格式：Vant-{标识}
        menuIcon: '常用格式：Autoicon' + "{'{'}" + '集合' + "{'}{'}" + '标识' + "{'}'}" + '；Vant格式：Vant-' + "{'{'}" + '标识' + "{'}'}",
        // extraData: 'JSON格式。说明：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}}',
        extraData: 'JSON格式。说明：' + "{'{'}" + '"i18n（国际化设置）": ' + "{'{'}" + '"title": ' + "{'{'}" + '"语言标识":"标题",...' + "{'}}}'}",
    }
}