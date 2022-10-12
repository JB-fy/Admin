import { defineStore } from 'pinia'

export const useSettingStore = defineStore('setting', {
  state: () => {
    return {
      leftMenuFold: false,
    }
  },
  actions: {
    /**
     * 折叠左侧菜单
     */
    leftMenuFold() {
      this.leftMenuFold = (!this.leftMenuFold)
    },
  },
})
