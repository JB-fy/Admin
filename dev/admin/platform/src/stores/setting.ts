import { defineStore } from 'pinia'

export const useSettingStore = defineStore('setting', {
  state: () => {
    return {
      leftMenuFold: false,  //左侧菜单折叠状态
      paginationSize: 20, //分页组件默认每页条数
      scrollSize: 10, //滚动加载等组件默认每页条数
      saveDrawer: { //保存组件抽屉的宽度
        size: '50%'
      },
      language: {
        'zh-cn': '中文（简体）',  //Chinese（Simplified）
        'en': 'English',  //英文
      },
    }
  },
  // actions: {
  //   /**
  //    * 折叠左侧菜单
  //    */
  //   leftMenuFold() {
  //     this.leftMenuFold = (!this.leftMenuFold)
  //   },
  // },
})
