import { createI18n } from 'vue-i18n'
//import { createI18n } from 'vue-i18n/dist/vue-i18n.cjs.js'  //可以解决控制台警告（也可以在vite.config.ts中设置别名解决）：You are running the esm-bundler build of vue-i18n. It is recommended to configure your bundler to explicitly replace feature flag globals with boolean literals to get proper tree-shaking in the final bundle.

const i18n = createI18n({
    legacy: false,  //会导致不能动态刷新，需强制刷新页面。但可解决以下报错。当使用useI18n()时会报错：Uncaught SyntaxError: Not available in legacy mode。也可以自定义一个getI18n函数替代useI18n。
    //locale: 'zh-cn',
    locale: getLanguage(),
    fallbackLocale: ['zh-cn', 'en'],
    messages: await batchImport(import.meta.globEager('@/i18n/language/**/*.ts'), 1, 10, false),
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