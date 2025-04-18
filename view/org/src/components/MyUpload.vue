<!-------- 使用示例 开始-------->
<!-- accept示例：image/*; video/*; audio/*; text/*; application/*; .png,.xls,.pdf,.apk,.ipa等 -->
<!-- <my-upload v-model="saveForm.data.avatar" accept="image/*" :multiple="true" />

<my-upload v-model="saveForm.data.avatar" :api="{ data: { scene: '指定上传场景' } }" accept="video/*" size="small" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="tsx">
import clipboard3 from 'vue-clipboard3'

const { toClipboard } = clipboard3()
const { t } = useI18n()

defineOptions({ inheritAttrs: false })
const attrs = useAttrs()
const slots = useSlots()
const emits = defineEmits(['update:modelValue', 'change'])
const props = defineProps({
    modelValue: {
        //单选传字符串，多选传数组
        type: [String, Array],
    },
    /**
     * 接口。格式：{ code: string, data: Object }
     *      code：非必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      data：非必须。接口函数所需参数。格式：{ [propName: string]: any }
     */
    api: {
        type: Object,
    },
    maxSize: {
        //需要限制文件大小时使用，单位：字节。示例：100 * 1024 * 1024
        type: Number,
        default: 0,
    },
    acceptType: {
        //需要严格限制文件格式时使用。示例：['image/png','image/jpg','image/jpeg','image/gif']
        type: Array,
        default: () => [],
    },
    size: {
        //尺寸。注意：只在listType=picture-card时有效
        type: String,
        validator: (value: string) => (value ? ['default', 'small'].includes(value) : true),
    },
    showType: {
        //显示类型，默认根据文件后缀显示，也可传值强制显示特定类型。注意：只在listType=picture-card时有效
        type: String,
        validator: (value: string /* , props */) => (value ? ['image', 'video', 'audio', 'text', 'application'].includes(value) : true),
    },
})
const showTypeMap: { [propName: string]: string } = {
    '.xbm': 'image',
    '.tif': 'image',
    '.jfif': 'image',
    '.pjp': 'image',
    '.apng': 'image',
    '.pjpeg': 'image',
    '.avif': 'image',
    '.ico': 'image',
    '.tiff': 'image',
    '.gif': 'image',
    '.svg': 'image',
    '.bmp': 'image',
    '.png': 'image',
    '.jpeg': 'image',
    '.svgz': 'image',
    '.jpg': 'image',
    '.webp': 'image',

    '.ogm': 'video',
    '.wmv': 'video',
    '.mpg': 'video',
    '.webm': 'video',
    '.ogv': 'video',
    '.mov': 'video',
    '.asx': 'video',
    '.mpeg': 'video',
    '.mp4': 'video',
    '.m4v': 'video',
    '.avi': 'video',

    '.opus': 'audio',
    '.flac': 'audio',
    '.weba': 'audio',
    '.wav': 'audio',
    '.ogg': 'audio',
    '.m4a': 'audio',
    '.oga': 'audio',
    '.mid': 'audio',
    '.mp3': 'audio',
    '.aiff': 'audio',
    '.wma': 'audio',
    '.au': 'audio',
    // '.webm': 'audio',

    '.zip': 'application',
    '.crt': 'application',
    '.docx': 'application',
    '.xlsx': 'application',
    '.ppt': 'application',
    '.xul': 'application',
    '.apk': 'application',
    '.ipa': 'application',
    '.tar': 'application',
    '.ai': 'application',
    '.ps': 'application',
    '.rss': 'application',
    '.p7s': 'application',
    '.woff': 'application',
    '.p7z': 'application',
    '.p7c': 'application',
    '.pptx': 'application',
    '.pdf': 'application',
    '.exe': 'application',
    '.rtf': 'application',
    '.bin': 'application',
    '.p7m': 'application',
    '.swf': 'application',
    '.xhtm': 'application',
    '.dot': 'application',
    '.swl': 'application',
    '.doc': 'application',
    '.xls': 'application',
    '.json': 'application',
    '.m3u8': 'application',
    '.epub': 'application',
    '.gz': 'application',
    '.com': 'application',
    '.rdf': 'application',
    '.cer': 'application',
    '.xhtml': 'application',
    '.tgz': 'application',
    '.xht': 'application',
    '.eps': 'application',
    '.crx': 'application',
    '.wasm': 'application',
    // '.js': 'application',

    '.xbl': 'text',
    '.xsl': 'text',
    '.text': 'text',
    '.xslt': 'text',
    '.txt': 'text',
    '.ehtml': 'text',
    '.sh': 'text',
    '.html': 'text',
    '.ics': 'text',
    '.mjs': 'text',
    '.js': 'text',
    '.shtml': 'text',
    '.xml': 'text',
    '.csv': 'text',
    '.css': 'text',
    '.shtm': 'text',
    '.htm': 'text',
}

const upload = reactive({
    id: ('my-upload_' + new Date().getTime() + '_' + randomInt(1000, 9999)) as string, //用于判断组件是否已经销毁，防止倒计时重复执行
    ref: null as any,
    value: ((): any => {
        if (attrs.multiple) {
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
            let valList = props.modelValue
            if (!attrs.multiple) {
                valList = [props.modelValue]
            }
            return (valList as string[]).map((item) => ({ name: item.slice(item.lastIndexOf('/') + 1), url: item }))
        },
        set: (val) => {
        }
    }), */
    //这个方式动画效果好，但不能动态刷新组件（即组件外部改变modelValue时，不会刷新）。处理方法：在组件使用的地方设置key来刷新
    fileList: (() => {
        if (!props.modelValue) {
            return []
        }
        let valList = props.modelValue
        if (!attrs.multiple) {
            valList = [props.modelValue]
        }
        return (valList as string[]).map((item) => ({ name: item.slice(item.lastIndexOf('/') + 1), url: item }))
    })(),
    class: computed((): string => {
        let classStr = 'upload-container'
        props.size == 'small' && (classStr += ' small')
        if (attrs.multiple) {
            attrs.limit && attrs.limit == upload.fileList.length && (classStr += ' hide')
        } else {
            upload.fileList.length && (classStr += ' hide')
        }
        return classStr
    }),
    action: '' as string,
    data: {} as { [propName: string]: any },
    signInfo: {} as { [propName: string]: any }, //缓存的签名信息。示例：{ upload_url: "https://xxxxx.com/upload", upload_data: {...}, host: "https://xxxxx.com", dir: "common/20221231/", expire: 1672471578, is_res: 1 }
    initSignInfo: async () => {
        const signInfo = await upload.api.getSignInfo()
        if (signInfo && Object.keys(signInfo).length) {
            upload.signInfo = { ...signInfo }
            upload.action = upload.signInfo.upload_url
            upload.data = { ...upload.signInfo.upload_data }
            //授权失效前，重新获取授权, 提前bufferTime更新，防止使用时失效
            let bufferTime = 10 * 1000 //缓冲时间
            let timeout = upload.signInfo.expire * 1000 - new Date().getTime() - bufferTime
            setTimeout(() => {
                //组件销毁后，倒计时还会继续执行。如果用户点击新增|编辑|复制等按钮多次，将会创建多个倒计时
                //upload.initSignInfo()
                //判断元素是否还存在，防止组件销毁后，倒计时却还在重复执行
                document.getElementById(upload.id) && upload.initSignInfo()
            }, timeout)
        }
    },
    //生成保存文件名和访问地址（访问地址在上传没有返回文件地址时使用。目前只在阿里云oss上传未设置回调时有用）
    createSaveInfo: (rawFile: any) => {
        let fileName = upload.signInfo.dir + rawFile.uid + '_' + randomInt(10000000, 99999999) + rawFile.name.slice(rawFile.name.lastIndexOf('.'))
        let url = upload.signInfo.host + '/' + fileName
        return {
            fileName: fileName,
            url: url,
        }
    },
    api: {
        loading: false,
        code: props.api?.code ?? t('config.VITE_HTTP_API_PREFIX') + '/upload/sign',
        data: { ...props.api?.data },
        getSignInfo: async () => {
            if (upload.api.loading) {
                return
            }
            upload.api.loading = true
            let signInfo = {}
            try {
                const res = await request(upload.api.code, upload.api.data)
                signInfo = res.data
            } finally {
                upload.api.loading = false
            }
            return signInfo
        },
    },
    onPreview: (file: any) => {
        imageViewer.initialIndex = imageViewer.urlList.indexOf(file.url)
        imageViewer.visible = true
    },
    onRemove: (file: any) => {
        if (attrs.multiple) {
            upload.value.splice(upload.value.indexOf(upload.getUrl(file)), 1)
        } else {
            upload.value = ''
        }
        emits('update:modelValue', upload.value)
        emits('change')
    },
    onSuccess: (res: any, file: any, fileList: any) => {
        if (upload.signInfo?.is_res) {
            //如有回调服务器且有报错，则默认失败
            if (res.code !== 0) {
                ElMessage.error(t('common.tip.uploadFail') + '(' + (res.msg ?? res) + ')')
                fileList.splice(fileList.indexOf(file), 1)
                return
            }
            file.raw.saveInfo.url = res.data.url //有返回以服务器返回地址为准
        }
        if (attrs.multiple) {
            upload.value.push(file.raw.saveInfo.url)
        } else {
            upload.value = file.raw.saveInfo.url
            upload.fileList = [file]
        }
        emits('update:modelValue', upload.value)
        emits('change')
    },
    onError: (err: Error /* , file: any, fileList: any */) => ElMessage.error(t('common.tip.uploadFail') + '(' + err.message + ')'),
    beforeUpload: async (rawFile: any) => {
        if (props.acceptType.length > 0 && !props.acceptType.includes(rawFile.type)) {
            ElMessage.error(t('common.tip.notAcceptFileType'))
            return false
        }
        if (props.maxSize > 0 && props.maxSize < rawFile.size / 1024 / 1024) {
            ElMessage.error(t('common.tip.notWithinFileSize'))
            return false
        }
        rawFile.saveInfo = upload.createSaveInfo(rawFile)
        upload.data.key = rawFile.saveInfo.fileName //这是文件保存路径及文件名，必须唯一，否则会覆盖oss服务器同名文件
    },
    getUrl: (file: any): string => (file?.response === undefined ? file.url : file.raw.saveInfo.url),
    copyUrl: (file: any) =>
        toClipboard(upload.getUrl(file))
            .then(() => ElMessage.success(t('common.copy') + t('common.success')))
            .catch((err) => ElMessage.error(t('common.copy') + t('common.fail') + ':' + err.message)),
    download: (file: any) => window.open(upload.getUrl(file)),
    showType: (file: any): string => {
        if (props.showType) {
            return props.showType
        }
        if (attrs.accept && (attrs.accept as string).includes('/')) {
            return (attrs.accept as string).split('/')[0]
        }
        if (file.raw && file.raw.type.includes('/')) {
            return file.raw.type.split('/')[0]
        }
        let url = upload.getUrl(file)
        let fileSuffix = url.slice(0, url.lastIndexOf('?'))
        fileSuffix = fileSuffix.slice(fileSuffix.lastIndexOf('.'))
        return showTypeMap[fileSuffix] ?? 'text'
    },
})

const imageViewer = reactive({
    urlList: computed((): string[] => upload.fileList.filter((item) => upload.showType(item) == 'image').map((item) => item.url)),
    initialIndex: 0,
    visible: false,
    close: () => (imageViewer.visible = false),
})

upload.initSignInfo() //初始化签名信息
</script>

<template>
    <div :id="upload.id">
        <el-upload
            v-if="['text' , 'picture'].includes($attrs.listType as string)"
            :ref="(el: any) => upload.ref = el"
            v-model:file-list="upload.fileList"
            :action="upload.action"
            :data="upload.data"
            :before-upload="upload.beforeUpload"
            :on-success="upload.onSuccess"
            :on-error="upload.onError"
            :on-remove="upload.onRemove"
            :on-preview="upload.onPreview"
            v-bind="$attrs"
        >
            <template #default>
                <slot v-if="slots.default" name="default"></slot>
                <el-button v-else type="primary">{{ t('common.upload') }}</el-button>
            </template>
            <template v-if="slots.trigger" #trigger>
                <slot name="trigger"></slot>
            </template>
            <template v-if="slots.tip" #tip>
                <slot name="tip"></slot>
            </template>
            <template v-if="slots.file" #file="{ file }">
                <slot name="file" :file="file"></slot>
            </template>
        </el-upload>
        <el-upload
            v-else
            :ref="(el: any) => upload.ref = el"
            v-model:file-list="upload.fileList"
            :action="upload.action"
            :data="upload.data"
            :before-upload="upload.beforeUpload"
            :on-success="upload.onSuccess"
            :on-error="upload.onError"
            :on-remove="upload.onRemove"
            :on-preview="upload.onPreview"
            v-bind="$attrs"
            list-type="picture-card"
            :drag="true"
            :class="upload.class"
        >
            <template #default>
                <slot v-if="slots.default" name="default"></slot>
                <template v-else>
                    <el-icon class="el-icon--upload"><autoicon-ep-upload-filled /></el-icon>
                    <div v-if="size != 'small'" class="el-upload__text" v-html="t('common.tip.uploadOrDrop')"></div>
                </template>
            </template>
            <template v-if="slots.trigger" #trigger>
                <slot name="trigger"></slot>
            </template>
            <template v-if="slots.tip" #tip>
                <slot name="tip"></slot>
            </template>
            <template #file="{ file }">
                <slot v-if="slots.file" name="file" :file="file"></slot>
                <template v-else>
                    <template v-if="['ready', 'uploading'].includes(file.status)">
                        <el-progress v-if="size == 'small'" type="circle" :percentage="file.percentage" :stroke-width="3" :width="45" />
                        <el-progress v-else type="circle" :percentage="file.percentage" />
                    </template>
                    <template v-else>
                        <template v-if="upload.showType(file) == 'image'">
                            <img class="el-upload-list__item-thumbnail" :src="file.url" />
                        </template>
                        <template v-else-if="upload.showType(file) == 'video'">
                            <el-icon v-if="size == 'small'" :size="38" style="width: 100%; position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%)"><autoicon-ep-film /></el-icon>
                            <video v-else class="el-upload-list__item-thumbnail" preload="none" :controls="true" :src="file.url" />
                        </template>
                        <template v-else-if="upload.showType(file) == 'audio'">
                            <el-icon v-if="size == 'small'" :size="38" style="width: 100%; position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%)"><autoicon-ep-mic /></el-icon>
                            <audio v-else preload="none" :controls="true" :src="file.url" style="width: 100%; height: 40px; position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%)" />
                        </template>
                        <template v-else-if="upload.showType(file) == 'application'">
                            <el-icon :size="size == 'small' ? 38 : 100" style="width: 100%; position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%)"><autoicon-ep-box /></el-icon>
                        </template>
                        <template v-else>
                            <el-icon :size="size == 'small' ? 38 : 100" style="width: 100%; position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%)"><autoicon-ep-document /></el-icon>
                        </template>

                        <template v-if="size == 'small'">
                            <span class="el-upload-list__item-actions">
                                <span v-if="file?.response === undefined" @click="upload.download(file)"><autoicon-ep-download /></span>
                                <span @click="upload.copyUrl(file)"><autoicon-ep-document-copy /></span>
                            </span>
                            <el-icon v-if="!$attrs.disabled" class="el-icon--close" @click="upload.ref.handleRemove(file)"><autoicon-ep-close /></el-icon>
                        </template>
                        <template v-else>
                            <label class="el-upload-list__item-status-label">
                                <el-icon class="el-icon--check"><autoicon-ep-check /></el-icon>
                            </label>

                            <template v-if="['video', 'audio'].includes(upload.showType(file))">
                                <el-icon v-if="!$attrs.disabled" class="el-icon--close" @click="upload.ref.handleRemove(file)"><autoicon-ep-close /></el-icon>
                            </template>
                            <span v-else class="el-upload-list__item-actions">
                                <span v-if="upload.showType(file) == 'image'" @click="upload.onPreview(file)"><autoicon-ep-zoom-in /></span>
                                <!-- 刚上传的文件没必要给下载按钮 -->
                                <span v-else-if="file?.response === undefined" @click="upload.download(file)"><autoicon-ep-download /></span>
                                <span @click="upload.copyUrl(file)"><autoicon-ep-document-copy /></span>
                                <span v-if="!$attrs.disabled" @click="upload.ref.handleRemove(file)"><autoicon-ep-delete /></span>
                            </span>
                        </template>
                    </template>
                </template>
            </template>
        </el-upload>

        <el-image-viewer v-if="imageViewer.visible" :url-list="imageViewer.urlList" :initial-index="imageViewer.initialIndex" :hide-on-click-modal="true" @close="imageViewer.close" />
    </div>
</template>

<style scoped>
.upload-container :deep(.el-upload .el-upload-dragger) {
    border: none;
    height: 100%;
    padding: 0;
}

.upload-container.hide :deep(.el-upload) {
    display: none;
}

.upload-container :deep(.el-upload-list__item .el-icon--close) {
    background-color: var(--el-color-danger);
    border-radius: 50%;
    top: 5px;
    transform: translateY(0);
}

.upload-container :deep(.el-upload-list__item:hover .el-icon--close) {
    display: inline-flex;
}

.upload-container.small {
    --my-upload-container-small-wg: 50px;
}

.upload-container.small :deep(.el-upload) {
    width: var(--my-upload-container-small-wg);
    height: var(--my-upload-container-small-wg);
}

.upload-container.small :deep(.el-upload-dragger) {
    height: 100%;
}

.upload-container.small :deep(.el-upload-dragger .el-icon--upload) {
    font-size: 38px;
    margin-bottom: 0;
}

.upload-container.small :deep(.el-upload-list__item) {
    width: var(--my-upload-container-small-wg);
    height: var(--my-upload-container-small-wg);
}

.upload-container.small :deep(.el-upload-list__item:hover) {
    overflow: visible;
}

.upload-container.small :deep(.el-upload-list__item:hover .el-icon--close) {
    top: -7px;
    right: -7px;
}

.upload-container.small :deep(.el-progress) {
    width: auto;
}

.upload-container.small :deep(.el-progress__text) {
    min-width: auto;
    font-size: 12px !important;
}

.upload-container.small :deep(.el-upload-list__item-actions span) {
    width: 20px;
    height: 20px;
}

.upload-container.small :deep(.el-upload-list__item-actions span + span) {
    margin-left: 5px;
}
</style>
