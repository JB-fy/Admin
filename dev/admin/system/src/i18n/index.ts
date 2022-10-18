import { createI18n } from 'vue-i18n'

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
let result: string[], language: string, languageKey: string;
for (const path in messageList) {
    //const result = path.match(/.*\/(.+)\/(.+).ts$/);
    result = path.split('/')
    language = result[result.length - 2]
    languageKey = result[result.length - 1].split('.')[0]
    if (!messages[language]) {
        messages[language] = {}
    }
    messages[language][languageKey] = messageList[path].default
}

const i18n = createI18n({
    //locale: 'zh-cn',
    locale: getLanguage(),
    fallbackLocale: ['zh-cn', 'en'],
    messages: messages,
})

export default i18n