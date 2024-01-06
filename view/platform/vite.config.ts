import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver, VantResolver } from 'unplugin-vue-components/resolvers'
import Icons from 'unplugin-icons/vite'
import IconsResolver from 'unplugin-icons/resolver'

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  /* if (command === 'serve') {  //dev独有配置
    return {}
  } else if (command === 'build') {//build独有配置
    return {}
  } */
  const env = loadEnv(mode, process.cwd(), ''); //设置第三个参数为 '' 来加载所有环境变量，而不管是否有 `VITE_` 前缀。
  return {
    define: {
      // 'process.env': env
      __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: command === 'serve' ? true : false
    },
    base: env.VITE_BASE_PATH, //URL部署时的目录。打包时给所有静态资源路径加上该前缀,
    build: {
      outDir: env.VITE_OUT_DIR, //构建文件的输出目录
      assetsDir: 'static',  //放置生成的静态资源的目录(路径相对于outDir)
      target: ['edge90', 'chrome90', 'firefox90', 'safari15']
    },
    server: {
      host: '0.0.0.0',
      //port: 5173,
      //https: true,
      //open: true, //启动后打开默认浏览器
      //vite会代理所有地址。而以前的webpack是找不到路由才做代理（故本地开发需要按下面说明做设置）
      proxy: {
        [env.VITE_DEV_API_PREFIX]: {
          target: env.VITE_DEV_SERVER_PROXY,
          changeOrigin: true,
          rewrite: (path) => path.replace(env.VITE_DEV_API_PREFIX, '')
        },
      }
    },
    resolve: {
      alias: {
        "@": fileURLToPath(new URL("./src", import.meta.url)),
        "vue-i18n": 'vue-i18n/dist/vue-i18n.cjs.js',  //可以解决控制台警告（也可以在i18n/index.ts中直接更改引用解决）：You are running the esm-bundler build of vue-i18n. It is recommended to configure your bundler to explicitly replace feature flag globals with boolean literals to get proper tree-shaking in the final bundle.
      },
    },
    plugins: [
      vue(),
      vueJsx(),
      /*--------按需导入函数 开始--------*/
      AutoImport({
        imports: [  //加载包
          'vue',
          'vue-router',
          'vue-i18n',
          /* { //自定义
            '@/basic/functions.ts': [
              'config',
              //['config', 'getConfig'],
            ]
          } */
        ],
        dirs: [ //目录加载，递归加载则后面加上'/**'
          'src/stores',
          'src/common/**',
        ],
        resolvers: [
          ElementPlusResolver(),  //ElementPlus
          VantResolver(), //Vant
          IconsResolver({ //图标（格式：前缀-集合名-图标名。例：<i-ep-lock />）
            prefix: 'autoicon',  //标签前缀。默认前缀"i"，false取消前缀。（必须设置且是唯一字符串。代码中除图标使用外，其它地方不能以该字符串开头，否则会被当作图标处理。比如该值未设置，自定义组件rightHeader、布尔值false等会被当作图标处理）
            /* enabledCollections: ['ep'], //启用哪个图标集合，默认启用全部。全部可选集合：https://icones.js.org/
            alias: { //一些复杂的集合名称设置别名
              //别名: '集合名',
              park: 'icon-park',
            },
            customCollections: ['自定义集合名']  //自定义图标集合 */
          }),
        ],
      }),
      /*--------按需导入函数 结束--------*/

      /*--------按需导入组件 开始--------*/
      Components({
        /* dirs: [
          'src/app/components',
        ], */
        resolvers: [
          ElementPlusResolver(),
          VantResolver(),
          IconsResolver({
            prefix: 'autoicon'
          }),
        ],
      }),
      /*--------按需导入组件 结束--------*/

      /*--------按需下载图标 开始--------*/
      Icons({
        autoInstall: true,  //自动下载图标
        /* compiler: 'vue3',
        customCollections: {  //自定义图标集合
          '自定义集合名': FileSystemIconLoader(
            '@/assets/icons',
            svg => svg.replace(/^<svg /, '<svg fill="currentColor" '),
          ),
        }, */
      }),
      /*--------按需下载图标 结束--------*/
    ],
  }
});
