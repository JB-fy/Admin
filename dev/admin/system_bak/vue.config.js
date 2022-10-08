const { defineConfig } = require('@vue/cli-service')
const AutoImport = require('unplugin-auto-import/webpack')
const Components = require('unplugin-vue-components/webpack')
const { ElementPlusResolver, VantResolver } = require('unplugin-vue-components/resolvers')
const Icons = require('unplugin-icons/webpack')
const IconsResolver = require('unplugin-icons/resolver')
//import { FileSystemIconLoader } from 'unplugin-icons/loaders'

module.exports = defineConfig({
  transpileDependencies: true,  //编译时是否兼容ie浏览器
  //lintOnSave: process.env.NODE_ENV === 'production',  //每次保存代码时都启用eslint验证。
  lintOnSave: false,  //每次保存代码时都启用eslint验证。
  publicPath: process.env.VUE_APP_BASE_PATH,  //URL部署时的目录。打包时给所有静态资源路径加上该前缀
  outputDir: '../../..' + process.env.VUE_APP_BASE_PATH, //构建文件的输出目录
  assetsDir: 'static',  //放置生成的静态资源的目录(路径相对于outputDir)
  //indexPath: 'index.html',  //指定生成的index.html的输出路径(路径相对于outputDir)。也可以是一个绝对路径
  //filenameHashing: true,  //生成静态资源的文件名中包含了hash以便更好的控制缓存。
  devServer: {
    //host: '0.0.0.0',
    //port: 8080,
    //https: true,
    //open: true, //启动后打开默认浏览器
    proxy: process.env.VUE_APP_DEV_SERVER_PROXY,  //前端和后端API服务器不同时，开发环境下将任何未知请求(没有匹配到静态文件的请求)代理到后端API服务器
  },
  configureWebpack: { //配置webpack
    plugins: [
      /*--------按需导入js需要的函数 开始--------*/
      AutoImport({
        imports: [  //加载包
          'vue',
          'vue-router',
          'vuex',
          { //自定义
            '@/app/config/index.js': [
              'config',
              //['config', 'getConfig'],
            ]
          }
        ],
        dirs: [ //目录加载，递归加载则后面加上'/**'，例：src/app/basic/**
          'src/app/basic',
        ],
        resolvers: [
          /*--------ElementPlus 开始--------*/
          ElementPlusResolver(),
          /*--------ElementPlus 结束--------*/
          /*--------Vant 开始--------*/
          VantResolver(),
          /*--------Vant 结束--------*/
          /*--------图标（格式：前缀-集合名-图标名。例：<i-ep-lock />） 开始--------*/
          IconsResolver({
            prefix: 'autoicon',  //标签前缀。默认前缀"i"，false取消前缀。（一定要设置且是唯一字符串，即除了图标用到，代码中其他地方不能以字符串开头。否则容易冲突报错。例：自定义组件right-header被认定为图标ri/ght-header；false被认定为图标fa/lse。）
            /* enabledCollections: ['ep'], //启用哪个图标集合，默认启用全部。全部可选集合：https://icones.js.org/
            alias: { //一些复杂的集合名称设置别名
              //别名: '集合名',
              park: 'icon-park',
            },
            customCollections: ['自定义集合名']  //自定义图标集合 */
          }),
          /*--------图标（格式：前缀-集合名-图标名。例：<i-ep-lock />） 结束--------*/
        ],
      }),
      /*--------按需导入js需要的函数 结束--------*/

      /*--------按需导入html需要的组件 开始--------*/
      Components({
        dirs: [
          'src/app/components',
        ],
        resolvers: [
          ElementPlusResolver(),
          VantResolver(),
          IconsResolver({
            prefix: 'autoicon'
          }),
        ],
      }),
      /*--------按需导入html需要的组件 结束--------*/

      /*--------按需自动下载图标 开始--------*/
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
      /*--------按需自动下载图标 结束--------*/
    ]
  }
})
