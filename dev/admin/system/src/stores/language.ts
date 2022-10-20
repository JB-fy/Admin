import { defineStore } from 'pinia'
import router from '@/router'
//import i18n from '@/i18n'

export const useLanguageStore = defineStore('language', {
  state: () => {
    return {
      language: getLanguage(),
      elementPlusLoacleList: import.meta.globEager('@/../node_modules/element-plus/dist/locale/*.min.mjs')
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
          return (<any>state.elementPlusLoacleList)['/node_modules/element-plus/dist/locale/' + state.language + '.min.mjs'].default
      }
    }
  },
  actions: {
    changeLanguage(language: string) {
      if (getLanguage() == language) {
        return
      }
      setLanguage(language)
      //this.language = language
      //i18n.global.locale = language
      router.go(0)  //直接刷新页面。接口数据也要刷新
    },
  },
})
