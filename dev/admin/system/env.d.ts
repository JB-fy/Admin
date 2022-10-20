/// <reference types="vite/client" />

declare module '*.vue' {
  import { ComponentOptions } from 'vue'
  const componentOptions: ComponentOptions
  export default componentOptions
}
// declare module '*.vue' {
//   import type { DefineComponent } from 'vue'
//   // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
//   const component: DefineComponent<{}, {}, any>
//   export default component
// }
/* interface ImportMeta {
  readonly env: ImportMetaEnv
}
interface ImportMetaEnv {
  readonly VITE_HTTP_TIMEOUT: number
} */