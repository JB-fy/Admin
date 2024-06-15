/// <reference types="vite/client" />

declare module '*.vue' {
    import { ComponentOptions } from 'vue'
    const componentOptions: ComponentOptions
    export default componentOptions
    /* import { DefineComponent } from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component */
}
declare module 'js-md5'