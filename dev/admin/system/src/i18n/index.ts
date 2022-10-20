import { createI18n } from 'vue-i18n'
//import { createI18n } from 'vue-i18n/dist/vue-i18n.cjs.js'  //可以解决控制台警告（也可以在vite.config.ts中设置别名解决）：You are running the esm-bundler build of vue-i18n. It is recommended to configure your bundler to explicitly replace feature flag globals with boolean literals to get proper tree-shaking in the final bundle.

const messages: { [propName: string]: {} } = batchImport(import.meta.globEager('@/i18n/language/**/*.ts'), 1, 10, false)

const i18n = createI18n({
    legacy: false,  //解决报错：Uncaught SyntaxError: Not available in legacy mode
    //locale: 'zh-cn',
    locale: getLanguage(),
    fallbackLocale: ['zh-cn', 'en'],
    messages: messages,
})

export default i18n


/*--------使用方式 开始--------*/
/* import i18n from '@/i18n';
i18n.global.locale
i18n.global.t('common.login')

import { useI18n } from 'vue-i18n';
const { locale, t } = useI18n()
useI18n().locale
useI18n().t('common.login')

{{ $t('common.login') }}

hello: '你好，{name}！'
useI18n().t('hello', { name: '名字' })

hello: '你好，{0}！'
useI18n().t('hello', ['名字'])

hello: 'hello <br> world'
< p v-html="$t('hello')"></p> */
/*--------使用方式 结束--------*/