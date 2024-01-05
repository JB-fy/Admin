export default {
    name: {
        menuName: '名称',
        sceneId: '场景',
        pid: '父级',
        level: '层级',
        idPath: '层级路径',
        menuIcon: '图标',
        menuUrl: '链接',
        extraData: '额外数据',
        sort: '排序值',
        isStop: '停用'
    },
    status: {},
    tip: {
        menuIcon: '常用格式：Autoicon' + "{'{'}" + '集合' + "{'}{'}" + '标识' + "{'}'}" + '；Vant格式：Vant-' + "{'{'}" + '标识' + "{'}'}",
        extraData: 'JSON格式：' + "{'{'}" + '"i18n（国际化设置）": ' + "{'{'}" + '"title": ' + "{'{'}" + '"语言标识":"标题",...' + "{'}'}" + '' + "{'}'}" + '',
        sort: '从小到大排序，默认50，范围0-100'
    }
}
