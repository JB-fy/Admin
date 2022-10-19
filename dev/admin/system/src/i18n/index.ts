import { createI18n } from 'vue-i18n'
//import { createI18n } from 'vue-i18n/dist/vue-i18n.cjs.js'  //可以解决控制台警告（也可以在vite.config.ts中设置别名解决）：You are running the esm-bundler build of vue-i18n. It is recommended to configure your bundler to explicitly replace feature flag globals with boolean literals to get proper tree-shaking in the final bundle.

export const getLanguage = () => {
    const language = localStorage.getItem('language')
    if (language) {
        return language
    }
    return (navigator.language || 'zh-cn').toLowerCase()
}

const messages: any = {
    'zh-cn': {},
    'en': {}
}
const messageList = import.meta.globEager('@/i18n/language/**/*.ts')
let keyList: string[], key1: string, key2: string;
for (const path in messageList) {
    //keyList = path.match(/.*\/(.+)\/(.+).ts$/);
    keyList = path.slice(0, path.lastIndexOf('.')).split('/')
    key1 = keyList[keyList.length - 2]
    key2 = keyList[keyList.length - 1]
    if (!messages[key1]) {
        messages[key1] = {}
    }
    messages[key1][key2] = (<any>messageList[path]).default
}

const i18n = createI18n({
    //locale: 'zh-cn',
    locale: getLanguage(),
    fallbackLocale: ['zh-cn', 'en'],
    messages: messages,
})

export default i18n