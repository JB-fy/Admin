<!-------- 使用示例 开始-------->
<!-- <MyUpload v-model="saveForm.data.avatar" accept="image/*" :multiple="true" />

<MyUpload v-model="saveForm.data.avatar" :api="{ param: { type: 'common' } }" accept="video/*" :isImage="false" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="ts">
const { t } = useI18n()

const props = defineProps({
    modelValue: {   //单选传字符串，多选传数组
        type: [String, Array]
    },
    /**
     * 接口。格式：{ code: string, param: Object }
     *      code：非必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：非必须。接口函数所需参数。格式：{ [propName: string]: any }
     */
    api: {
        type: Object
    },
    acceptType: {   //需要严格限制文件格式时使用。示例：['image/png','image/jpg','image/jpeg','image/gif']
        type: Array,
        default: []
    },
    maxSize: {  //需要限制文件大小时使用，单位：字节。示例：100 * 1024 * 1024
        type: Number,
        default: 0
    },
    isImage: { //是否显示图片缩略图
        type: Boolean,
        default: true
    },
    tip: {
        type: String,
        //default: 'jpg/png files with a size less than 500kb'
    },
    multiple: {
        type: Boolean,
        default: false
    },
    limit: {
        type: Number
    },
    accept: {   //文件选择弹出框过滤用，但可被人工跳过。示例：'image/*'、'video/*'、'audio/*'、'.git,.png'等
        type: String,
        default: ''
    },
})

const emits = defineEmits(['update:modelValue', 'change'])
const upload = reactive({
    id: 'MyUpload' + new Date().getTime() + '_' + randomInt(1000, 9999) as string,   //用于判断组件是否已经销毁，防止倒计时重复执行
    ref: null as any,
    value: ((): any => {
        if (props.multiple) {
            return props.modelValue ? [...(props.modelValue as string[])] : []
        }
        return props.modelValue
    })(),
    /* //这个方式动画效果不好，但可以动态刷新组件（即组件外部改变modelValue时，会刷新）。待处理bug：多文件上传时，onSuccess内执行emits('update:modelValue', upload.value)会触发get，导致第二个文件上传被中断
    fileList: computed({
        get: () => {
            if (!props.modelValue) {
                return []
            }
            if (props.multiple) {
                return (props.modelValue as string[]).map((item) => {
                    return {
                        name: item.slice(item.lastIndexOf('/') + 1),
                        url: item
                    }
                })
            }
            return [{
                name: (props.modelValue as string).slice((props.modelValue as string).lastIndexOf('/') + 1),
                url: (props.modelValue as string)
            }]
        },
        set: (val) => {
        }
    }), */
    //这个方式动画效果好，但不能动态刷新组件（即组件外部改变modelValue时，不会刷新）。处理方法：在组件使用的地方设置key来刷新
    fileList: (() => {
        if (!props.modelValue) {
            return []
        }
        if (props.multiple) {
            return (props.modelValue as string[]).map((item) => {
                return {
                    name: item.slice(item.lastIndexOf('/') + 1),
                    url: item
                }
            })
        }
        return [{
            name: (props.modelValue as string).slice((props.modelValue as string).lastIndexOf('/') + 1),
            url: (props.modelValue as string)
        }]
    })(),
    class: computed((): string => {
        if (props.multiple) {
            return props.limit && props.limit == upload.fileList.length ? 'hide' : ''
        } else {
            return upload.fileList.length ? 'hide' : ''
        }
    }),
    action: '' as string,
    data: {} as { [propName: string]: any },
    signInfo: {} as { [propName: string]: any },    //缓存的签名信息。示例：{ uploadUrl: "https://xxxxx.com/upload", uploadData: {...}, host: "https://xxxxx.com", dir: "common/20221231/", expire: 1672471578, isRes: 1 }
    //生成保存在云服务器中的文件名及完成地址
    initSignInfo: async () => {
        const signInfo = await upload.api.getSignInfo()
        if (signInfo && Object.keys(signInfo).length) {
            upload.signInfo = { ...signInfo }
            upload.action = upload.signInfo.uploadUrl
            upload.data = { ...upload.signInfo.uploadData }
            //授权失效前，重新获取授权, 提前bufferTime更新，防止使用时失效
            let bufferTime = 10 * 1000 //缓冲时间
            let timeout = upload.signInfo.expire * 1000 - new Date().getTime() - bufferTime
            setTimeout(() => {
                //组件销毁后，倒计时还会继续执行。如果用户点击新增|编辑|复制等按钮多次，将会创建多个倒计时
                //upload.initSignInfo()
                //判断元素是否还存在，防止组件销毁后，倒计时却还在重复执行
                document.getElementById(upload.id) ? upload.initSignInfo() : null
            }, timeout)
        }
    },
    createSaveInfo: (rawFile: any) => {
        let fileName = upload.signInfo.dir + rawFile.uid + '_' + randomInt(1000, 9999) + rawFile.name.slice(rawFile.name.lastIndexOf('.'))
        let url = upload.signInfo.host + '/' + fileName
        return {
            fileName: fileName,
            url: url
        }
    },
    api: {
        loading: false,
        code: props.api?.code ?? t('config.VITE_HTTP_API_PREFIX') + '/upload/sign',
        param: {
            ...props.api?.param
        },
        getSignInfo: async () => {
            if (upload.api.loading) {
                return
            }
            upload.api.loading = true
            let signInfo = {}
            try {
                const res = await request(upload.api.code, upload.api.param)
                signInfo = res.data
            } catch (error) { }
            upload.api.loading = false
            return signInfo
        },
    },
    onPreview: (uploadFile: any) => {
        imageViewer.initialIndex = imageViewer.urlList.indexOf(uploadFile.url)
        imageViewer.visible = true
    },
    onRemove: (file: any, fileList: any) => {
        //上传前处理函数beforeUpload返回false时也会触发此函数。此时file内没有response，但是由于没上传也不会存在于props.modelValue中，故不影响删除逻辑
        let url: string = file?.response === undefined ? file.url : file.raw.saveInfo.url
        if (props.multiple) {
            upload.value.splice(upload.value.indexOf(url), 1)
        } else {
            upload.value = ''
        }
        emits('change')
        emits('update:modelValue', upload.value)
    },
    onSuccess: (res: any, file: any, fileList: any) => {
        if (upload.signInfo?.isRes) {    //如有回调服务器且有报错，则默认失败
            if (res.code !== 0) {
                ElMessage.error(t('common.tip.uploadFail'))
                fileList.splice(fileList.indexOf(file), 1)
                return
            }
            file.raw.saveInfo.url = res.data.url  //有返回以服务器返回地址为准
        }
        if (props.multiple) {
            upload.value.push(file.raw.saveInfo.url)
        } else {
            upload.value = file.raw.saveInfo.url
            upload.fileList = [file]
        }
        emits('change')
        emits('update:modelValue', upload.value)
    },
    beforeUpload: async (rawFile: any) => {
        if (props.acceptType.length > 0 && props.acceptType.indexOf(rawFile.type) === -1) {
            ElMessage.error(t('common.tip.notAcceptFileType'))
            return false
        }
        if (props.maxSize > 0 && props.maxSize < rawFile.size / 1024 / 1024) {
            ElMessage.error(t('common.tip.notWithinFileSize'))
            return false
        }
        rawFile.saveInfo = upload.createSaveInfo(rawFile)
        upload.data.key = rawFile.saveInfo.fileName //这是文件保存路径及文件名，必须唯一，否则会覆盖oss服务器同名文件
    }
})

const imageViewer = reactive({
    urlList: computed((): string[] => {
        return upload.fileList.map((item) => {
            return item.url
        })
    }),
    initialIndex: 0,
    visible: false,
    close: () => {
        imageViewer.visible = false
    }
})

upload.initSignInfo()   //初始化签名信息
</script>

<template>
    <div :id="upload.id">
        <div v-if="isImage" class="upload-container">
            <ElUpload :ref="(el: any) => { upload.ref = el }" v-model:file-list="upload.fileList" :action="upload.action"
                :data="upload.data" :before-upload="upload.beforeUpload" :on-success="upload.onSuccess"
                :on-remove="upload.onRemove" :on-preview="upload.onPreview" :multiple="multiple" :limit="limit"
                :accept="accept" list-type="picture-card" :drag="true" :class="upload.class">
                <ElIcon class="el-icon--upload">
                    <AutoiconEpUploadFilled />
                </ElIcon>
                <div class="el-upload__text" v-html="t('common.tip.uploadOrDrop')"></div>
                <template v-if="tip" #tip>
                    <div class="el-upload__tip">
                        {{ tip }}
                    </div>
                </template>
            </ElUpload>
            <ElImageViewer v-if="imageViewer.visible" :url-list="imageViewer.urlList"
                :initial-index="imageViewer.initialIndex" :hide-on-click-modal="true" @close="imageViewer.close" />
        </div>
        <ElUpload v-else :ref="(el: any) => { upload.ref = el }" v-model:file-list="upload.fileList" :action="upload.action"
            :data="upload.data" :before-upload="upload.beforeUpload" :on-success="upload.onSuccess"
            :on-remove="upload.onRemove" :multiple="multiple" :limit="limit" :accept="accept" list-type="text">
            <ElButton type="primary">{{ t('common.upload') }}</ElButton>
            <template v-if="tip" #tip>
                <div class="el-upload__tip">
                    {{ tip }}
                </div>
            </template>
        </ElUpload>
    </div>
</template>

<style scoped>
.upload-container :deep(.el-upload .el-upload-dragger) {
    border: none;
    height: 146px;
    padding: 0;
}

.upload-container :deep(.el-upload) {
    width: 146px;
    margin-right: 10px;
}

.upload-container :deep(.hide .el-upload) {
    display: none;
}
</style>