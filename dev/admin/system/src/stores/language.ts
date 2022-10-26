import { defineStore } from 'pinia'
import router from '@/router'
//import i18n from '@/i18n'

const elementPlusLoacleList = await batchImport(import.meta.globEager('@/../node_modules/element-plus/dist/locale/*.min.mjs'))
export const useLanguageStore = defineStore('language', {
  state: () => {
    return {
      language: getLanguage(),
      //elementPlusLoacleList: import.meta.globEager('@/../node_modules/element-plus/dist/locale/*.min.mjs'),
      elementPlusLoacleList: elementPlusLoacleList,
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
    changeLanguage(language: string) {
      if (getLanguage() == language) {
        return
      }
      setLanguage(language)
      //this.language = language
      //i18n.global.locale = language
      /**
       * 下面这几种情况，需要使用router.go(0)，强制刷新页面
       *    当i18n设置legacy: false时，虽能解决报错问题，但会导致不能动态刷新。但又必须设置，理由使用了useI18n()会报错
       *    路由设置标题时，不能动态刷新
       *    接口数据，不能动态刷新
       */
      router.go(0)
    },
  },
})
