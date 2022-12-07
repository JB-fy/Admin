import { defineStore } from 'pinia'

export const useSettingStore = defineStore('setting', {
  state: () => {
    return {
      leftMenuFold: false,  //左侧菜单折叠状态
      paginationSize: 20, //全局分页组件默认每页条数
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
