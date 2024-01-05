import { defineStore } from 'pinia'

export const useSettingStore = defineStore('setting', {
    state: () => {
        return {
            language: {
                'zh-cn': '中文（简体）', //Chinese（Simplified）
                en: 'English' //英文
            },
            leftMenuFold: false, //左侧菜单折叠状态
            pagination: {
                //分页组件的配置
                size: 20, //默认每页条数
                sizeList: [10, 20, 50, 100, 500, 1000, 5000], //可选每页条数量
                layout: 'total, sizes, prev, pager, next, jumper' //样式
            },
            scrollSize: 20, //滚动加载等组件默认每页条数
            saveDrawer: {
                //保存组件抽屉的配置
                isTipClose: true, //退出是否提示
                size: '60%' //宽度
            },
            exportButton: {
                //导出按钮组件的配置
                limit: 50000 //单文件最大导出数量
            }
        }
    }
    // actions: {
    //   /**
    //    * 折叠左侧菜单
    //    */
    //   leftMenuFold() {
    //     this.leftMenuFold = (!this.leftMenuFold)
    //   },
    // },
})
