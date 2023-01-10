import { defineStore } from 'pinia'
import router from '@/router'
import i18n from '@/i18n'

export const useLanguageStore = defineStore('language', {
  state: () => {
    return {
      language: getLanguage(),
      //elementPlusLoacleList: import.meta.globEager('@/../node_modules/element-plus/dist/locale/*.min.mjs'),
      elementPlusLoacleList: batchImport(import.meta.globEager('@/../node_modules/element-plus/dist/locale/*.min.mjs')),
    }
  },
  getters: {
    // elementPlusLocale: async (state) => {
    //   switch (state.language) {
    //     default:
    //       //return (await import(/* @vite-ignore */'../../node_modules/element-plus/dist/locale/' + state.language + '.mjs')).default
    //       return (await import(/* @vite-ignore */'/node_modules/element-plus/dist/locale/' + state.language + '.mjs')).default
    //   }
    // },
    elementPlusLocale: (state) => {
      switch (state.language) {
        default:
          //console.log(state.elementPlusLoacleList)
          //return (<any>state.elementPlusLoacleList)['/node_modules/element-plus/dist/locale/' + state.language + '.min.mjs'].default
          return state.elementPlusLoacleList[state.language]
      }
    }
  },
  actions: {
    //改变语言
    changeLanguage(language: string) {
      if (getLanguage() == language) {
        return
      }
      setLanguage(language)
      this.language = language
      //i18n.global.locale = language //当i18n设置legacy: false，要使用i18n.global.locale.value
      i18n.global.locale.value = language

      document.title = this.getWebTitle()
      /**
       * 由于许多情况不会动态刷新，故建议直接刷新页面
       *    列举以下这几种不能动态刷新的情况
       *      路由设置标题时
       *      部分接口需重新请求
       *      t函数赋值的变量。如各种表单验证
       */
      router.go(0)  //刷新页面
    },
    //获取页面标题
    getMenuTitle(menu: any) {
      if (menu) {
        return menu?.title?.[i18n.global.locale.value] ?? menu.menuName
      }
      return ''
    },
    //获取页面标题
    getPageTitle(fullPath: string = router.currentRoute.value.fullPath) {
      const menu = useAdminStore().menuList.find((item) => {
        return item.url == fullPath
      }) ?? router.currentRoute.value?.meta?.menu
      return this.getMenuTitle(menu)
    },
    //获取网站标题
    getWebTitle(fullPath: string = router.currentRoute.value.fullPath) {
      let webTitle = (<any>i18n).global.t('config.webTitle')
      const title = this.getPageTitle(fullPath)
      if (title) {
        webTitle += '-' + title
      }
      return webTitle
    },
  },
})
