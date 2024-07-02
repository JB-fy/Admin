<!-------- 使用示例 开始-------->
<!-- <my-editor v-model="saveForm.data.content" />

<my-editor v-model="saveForm.data.content" :api="{ param: { type: 'common' } }" :init="{width: '375px'}" :disabled="true" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="tsx">
import axios from 'axios'
import editor from '@tinymce/tinymce-vue'

const { t } = useI18n()
const languageStore = useLanguageStore()

const props = defineProps({
    modelValue: {
        type: String,
    },
    /**
     * 接口。格式：{ code: string, param: Object }
     *      code：非必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：非必须。接口函数所需参数。格式：{ [propName: string]: any }
     */
    api: {
        type: Object,
    },
    init: {
        type: Object,
        default: () => {},
    },
    disabled: {
        type: Boolean,
        default: false,
    },
})

const emits = defineEmits(['update:modelValue', 'change'])
const myEditor = reactive({
    id: ('my_editor' + new Date().getTime() + '_' + randomInt(1000, 9999)) as string, //用于判断组件是否已经销毁，防止倒计时重复执行
    ref: null as any,
    value: computed({
        get: () => {
            return props.modelValue
        },
        set: (val) => {
            emits('update:modelValue', val)
            emits('change')
        },
    }),
    init: {
        width: '100%',
        // height: 'auto',
        min_height: 500,
        language: languageStore.tinymceLocale,
        plugins: 'lists link image table code wordcount fullscreen help', //autoresize
        toolbar: 'undo redo | styles formatselect | bold italic | alignleft aligncenter alignright outdent indent bullist numlist | image fullscreen help',
        branding: false, // 右下角Tiny技术支持信息是否显示
        images_upload_handler: (blobInfo: any, progress: any) => {
            return new Promise((resolve, reject) => {
                let data: { [propName: string]: any } = { ...myEditor.signInfo.upload_data }
                const filename = blobInfo.filename()
                data.key = myEditor.signInfo.dir + blobInfo.id() + '_' + randomInt(1000, 9999) + filename.slice(filename.lastIndexOf('.'))
                data.file = blobInfo.blob()
                axios
                    .post(myEditor.signInfo.upload_url, data, { headers: { 'Content-Type': 'multipart/form-data' } })
                    .then((res) => {
                        let imgUrl = myEditor.signInfo.host + '/' + data.key
                        if (myEditor.signInfo?.is_res) {
                            if (res.data.code !== 0) {
                                reject(t('common.tip.uploadFail'))
                                return
                            }
                            imgUrl = res.data.data.url
                        }
                        resolve(imgUrl)
                    })
                    .catch((error) => {
                        reject(error.message)
                    })
                    .finally(() => {})
            })
        },
        /* file_picker_callback: (callback: any, value: any, meta: any) => {
            console.log(callback)
            console.log(value)
            console.log(meta)
        } */
        ...props.init,
    },
    signInfo: {} as { [propName: string]: any }, //缓存的签名信息。示例：{ upload_url: "https://xxxxx.com/upload", upload_data: {...}, host: "https://xxxxx.com", dir: "common/20221231/", expire: 1672471578, is_res: 1 }
    initSignInfo: async () => {
        const signInfo = await myEditor.api.getSignInfo()
        if (signInfo && Object.keys(signInfo).length) {
            myEditor.signInfo = { ...signInfo }

            //授权失效前，重新获取授权, 提前bufferTime更新，防止使用时失效
            let bufferTime = 10 * 1000 //缓冲时间
            let timeout = myEditor.signInfo.expire * 1000 - new Date().getTime() - bufferTime
            setTimeout(() => {
                //组件销毁后，倒计时还会继续执行。如果用户点击新增|编辑|复制等按钮多次，将会创建多个倒计时
                //myEditor.initSignInfo()
                //判断元素是否还存在，防止组件销毁后，倒计时却还在重复执行
                document.getElementById(myEditor.id) ? myEditor.initSignInfo() : null
            }, timeout)
        }
    },
    api: {
        loading: false,
        code: props.api?.code ?? t('config.VITE_HTTP_API_PREFIX') + '/upload/sign',
        param: {
            ...props.api?.param,
        },
        getSignInfo: async () => {
            if (myEditor.api.loading) {
                return
            }
            myEditor.api.loading = true
            let signInfo = {}
            try {
                const res = await request(myEditor.api.code, myEditor.api.param)
                signInfo = res.data
            } catch (error) {}
            myEditor.api.loading = false
            return signInfo
        },
    },
})

myEditor.initSignInfo() //初始化签名信息
</script>

<template>
    <div :id="myEditor.id" style="width: 100%">
        <editor :ref="(el: any) => myEditor.ref = el" v-model="myEditor.value" :init="myEditor.init" :disabled="disabled" />
    </div>
</template>

<!-- <style scoped> -->
<style>
.tox.tox-silver-sink.tox-tinymce-aux {
    z-index: 10000 !important;
}
</style>
