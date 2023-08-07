<!-------- 使用示例 开始-------->
<!-- <MyEditor v-model="saveForm.data.content" />

<MyEditor v-model="saveForm.data.content" :api="{ param: { type: 'common' } }" :disabled="true" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="ts">
import axios from 'axios'
import Editor from '@tinymce/tinymce-vue'

const { t } = useI18n()
const languageStore = useLanguageStore()

const props = defineProps({
    modelValue: {
        type: String
    },
    /**
     * 接口。格式：{ code: string, param: Object }
     *      code：非必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：非必须。接口函数所需参数。格式：{ [propName: string]: any }
     */
    api: {
        type: Object
    },
    init: {
        type: Object,
        default: {}
    },
    disabled: {
        type: Boolean,
        default: false
    },
})

const emits = defineEmits(['update:modelValue', 'change'])
const myEditor = reactive({
    id: 'MyEditor' + new Date().getTime() + '_' + randomInt(1000, 9999) as string,   //用于判断组件是否已经销毁，防止倒计时重复执行
    ref: null as any,
    value: computed({
        get: () => {
            return props.modelValue
        },
        set: (val) => {
            emits('change')
            emits('update:modelValue', val)
        }
    }),
    init: {
        width: "100%",
        language: languageStore.tinymceLocale,
        plugins: 'lists link image table code wordcount fullscreen help',
        toolbar: 'undo redo | styles formatselect | bold italic | alignleft aligncenter alignright outdent indent bullist numlist | image fullscreen help',
        branding: false, // 右下角Tiny技术支持信息是否显示
        images_upload_handler: (blobInfo: any, progress: any) => {
            return new Promise((resolve, reject) => {
                let data: { [propName: string]: any } = {
                    OSSAccessKeyId: myEditor.signInfo.accessid,
                    policy: myEditor.signInfo.policy,
                    signature: myEditor.signInfo.signature,
                    success_action_status: '200', //让服务端返回200,不然，默认会返回204
                }
                const filename = blobInfo.filename()
                data.key = myEditor.signInfo.dir + blobInfo.id() + '_' + randomInt(1000, 9999) + filename.slice(filename.lastIndexOf('.'))
                myEditor.signInfo?.callback ? data.callback = myEditor.signInfo.callback : null //是否回调服务器
                data.file = blobInfo.blob()
                axios.post(myEditor.signInfo.host, data, { headers: { "Content-Type": "multipart/form-data" } }).then((res) => {
                    if (res.data.code !== 0) {
                        reject(t('common.tip.uploadFail'))
                        return
                    }
                    let imgUrl = myEditor.signInfo.host + '/' + data.key
                    if (myEditor.signInfo?.callback) {    //如有回调服务器且有报错，则默认失败
                        imgUrl = res.data.data.url
                    }
                    resolve(imgUrl)
                }).catch((error) => {
                    reject(error.message)
                }).finally(() => {
                })
            })
        },
        /* file_picker_callback: (callback: any, value: any, meta: any) => {
            console.log(callback)
            console.log(value)
            console.log(meta)
        } */
        ...props.init,
    },
    signInfo: {} as { [propName: string]: any },    //缓存的签名信息。示例：{ accessid: "xxxx", host: "https://xxxxx.com", dir: "common/20221231/", expire: 1672471578, callback: "string", policy: "string", signature: "string" }
    //生成保存在云服务器中的文件名及完成地址
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
            }, 2000)
        }
    },
    api: {
        loading: false,
        code: props.api?.code ?? t('config.VITE_HTTP_API_PREFIX') + '/upload/sign',
        param: {
            ...props.api?.param
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
            } catch (error) { }
            myEditor.api.loading = false
            return signInfo
        },
    },
})


myEditor.initSignInfo()   //初始化签名信息
</script>

<template>
    <div :id="myEditor.id" style="width: 100%;">
        <Editor :ref="(el: any) => { myEditor.ref = el }" v-model="myEditor.value" :init="myEditor.init" :disabled="disabled" />
    </div>
</template>

<style scoped></style>