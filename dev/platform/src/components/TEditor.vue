<template>
  <editor v-model="myValue" :init="init" :disabled="disabled" :id="tinymceId"></editor>
</template>
  
  
<script setup lang="ts">
//JS部分
//在js中引入所需的主题和组件
import tinymce from 'tinymce/tinymce'
import 'tinymce/skins/content/default/content.css'
import Editor from '@tinymce/tinymce-vue'
import 'tinymce/themes/silver'
import 'tinymce/themes/silver/theme'
import 'tinymce/icons/default'; //引入编辑器图标icon，不引入则不显示对应图标
import 'tinymce/models/dom' // 这里是个坑 一定要引入


//在TinyMce.vue中接着引入相关插件
import "tinymce/icons/default/icons"
import "tinymce/plugins/image" // 插入上传图片插件
// import "tinymce/plugins/media" // 插入视频插件
import "tinymce/plugins/table" // 插入表格插件
import "tinymce/plugins/lists" // 列表插件
import "tinymce/plugins/wordcount" // 字数统计插件
import "tinymce/plugins/code" // 源码
import "tinymce/plugins/fullscreen" //全屏

//接下来定义编辑器所需要的插件数据
import axios from 'axios'
// import { updateImg } from '@/api/order/order'
const emits = defineEmits(["getContent"])
//这里我选择将数据定义在props里面，方便在不同的页面也可以配置出不同的编辑器，当然也可以直接在组件中直接定义
const props = defineProps({
  value: {
    type: String,
    default: () => {
      return ""
    },
  },
  baseUrl: {
    type: String,
    default: "",
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  plugins: {
    type: [String, Array],
    default: "lists  table fullscreen image imagetools",
  },//必填
  // toolbar: {
  //   type: [String, Array],
  //   default:
  //     "codesample bold italic underline alignleft aligncenter alignright alignjustify | undo redo | formatselect | fontselect | fontsizeselect | forecolor backcolor | bullist numlist outdent indent | lists link table code  | removeformat ",
  // },//必填
  toolbar: {
    type: [String, Array],
    default:
      "image fullscreen preview codesample code forecolor backcolor styles directionality fontfamily fontsize searchreplace image  media link alignleft aligncenter alignright template table charmap pagebreak nonbreaking anchor insertdatetime advlist lists wordcount autosave"
  },//必填
})
//用于接收外部传递进来的富文本
const myValue = ref(props.value)
const tinymceId = ref("vue-tinymce-" + +new Date() + ((Math.random() * 1000).toFixed(0) + ""))
//定义一个对象 init初始化
const init = reactive({



  selector: '#' + tinymceId.value, //富文本编辑器的id,
  language_url: "/tinymce/langs/zh-CN.js", // 语言包的路径，具体路径看自己的项目，文档后面附上中文js文件
  language: "zh_CN", //语言
  skin_url: "/tinymce/skins/ui/oxide", // skin路径，具体路径看自己的项目
  height: 400, //编辑器高度
  promotion: false, //Upgrade是否开启
  branding: false, //是否禁用“Powered by TinyMCE”
  menubar: true, //顶部菜单栏显示
  image_dimensions: true, //去除宽高属性
  plugins: props.plugins,  //这里的数据是在props里面就定义好了的
  toolbar: props.toolbar, //这里的数据是在props里面就定义好了的
  font_formats: 'Arial=arial,helvetica,sans-serif; 宋体=SimSun; 微软雅黑=Microsoft Yahei; Impact=impact,chicago;', //字体
  fontsize_formats: '11px 12px 14px 16px 18px 24px 36px 48px 64px 72px', //文字大小
  // paste_convert_word_fake_lists: false, // 插入word文档需要该属性
  paste_webkit_styles: "all",
  paste_merge_formats: true,
  nonbreaking_force_tab: false,
  paste_auto_cleanup_on_paste: false,
  file_picker_types: 'file',
  content_css: '/tinymce/skins/content/default/content.css', //以css文件方式自定义可编辑区域的css样式，css文件需自己创建并引入

  //图片上传
  images_upload_handler: (blobInfo: any, progress: any) => new Promise((resolve, reject) => {

    if (blobInfo.blob().size / 1024 / 1024 > 5) {
      reject({ message: '上传失败，图片大小请控制在 5M 以内', remove: true })
      return
    } else {
      getSignInfo().then(res => {

        const uploadInfo: any = res.data

        let params = new FormData()
        let config: any = {
          headers: {
            "Content-Type": "multipart/form-data",
          }
        }
        params.append('OSSAccessKeyId', uploadInfo.accessid)
        params.append('policy', uploadInfo.policy)
        params.append('signature', uploadInfo.signature)
        params.append('success_action_status', '200')
        let key = uploadInfo.dir + Date.parse(new Date() as any) + '_' + randomInt(1000, 9999) + blobInfo.filename()
        params.append('key', key)
        uploadInfo?.callback ? params.append('callback', uploadInfo.callback) : null //是否回调服务器
        params.append('file', blobInfo.blob())

        axios.post(uploadInfo.host, params, config).then(res => {
          if (res.data.code == 0) {
            resolve(res.data.data.url)  //上传成功，在成功函数里填入图片路径
          } else {
            reject('HTTP Error: 上传失败' + res.data.code);
            return
          }
        }).catch(() => {
          reject('上传出错，服务器开小差了呢')
          return
        })
      })

    }
  }),

  // 文件上传
  file_picker_callback: (callback: any, value: any, meta: any) => {
    // Provide file and text for the link dialog
    if (meta.filetype == 'file') {
      callback('mypage.html', { text: 'My text' });
    }

    // Provide image and alt text for the image dialog
    if (meta.filetype == 'image') {
      callback('myimage.jpg', { alt: 'My alt text' });
    }

    // Provide alternative source and posted for the media dialog
    if (meta.filetype == 'media') {
      callback('movie.mp4', { source2: 'alt.ogg', poster: 'image.jpg' });
    }
  }
})

const getSignInfo = async () => {
  return await request('/upload/sign', {})
}

//监听外部传递进来的的数据变化
watch(
  () => props.value,
  () => {
    myValue.value = props.value
    emits("getContent", myValue.value)
  }
)
//监听富文本中的数据变化
watch(
  () => myValue.value,
  () => {
    emits("getContent", myValue.value)
  }
)
//在onMounted中初始化编辑器
onMounted(() => {
  tinymce.init({
  })

})
</script>

<style>
/* 在el-dialog中tinymce z-index 被太小而被遮挡时要加这两句 */
.tox-tinymce-aux {
  z-index: 99999 !important;
}

.tinymce.ui.FloatPanel {
  z-index: 99;
}</style>