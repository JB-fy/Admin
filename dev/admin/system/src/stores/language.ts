import { defineStore } from 'pinia'
import { getLanguage } from '@/i18n'
import i18n from '@/i18n'

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
          return state.elementPlusLoacleList['/node_modules/element-plus/dist/locale/' + state.language + '.min.mjs'].default
      }
    }
  },
  actions: {
    changeLanguage(language: string) {
      localStorage.setItem('language', language)
      this.language = language
      i18n.global.locale = language
    },
  },
})
