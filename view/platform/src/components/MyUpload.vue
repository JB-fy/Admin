<!-------- 使用示例 开始-------->
<!-- <my-upload v-model="saveForm.data.avatar" accept="image/*" :multiple="true" />

<my-upload v-model="saveForm.data.avatar" :api="{ param: { type: 'common' } }" accept="video/*" show-type="video" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="tsx">
import type { EpPropMergeType } from 'element-plus/es/utils/vue/props/types'
import clipboard3 from 'vue-clipboard3'
const { toClipboard } = clipboard3()
const { t } = useI18n()

const slots = useSlots()
const props = defineProps({
    modelValue: {
        //单选传字符串，多选传数组
        type: [String, Array],
    },
    /**
     * 接口。格式：{ code: string, param: Object }
     *      code：非必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：非必须。接口函数所需参数。格式：{ [propName: string]: any }
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
    showType: {
        //显示类型，只在listType=picture-card时有用
        type: String,
        default: 'image',
        validator: (value: string, props) => ['image', 'video', 'audio', 'text'].includes(value),
    },
    // 以下属性参考原el-upload组件
    multiple: {
        type: Boolean,
        default: false,
    },
    accept: {
        //文件选择弹出框过滤用，但可被人工跳过。示例：image/*; video/*; audio/*; text/*; application/*; .png,.xls,.pdf,.apk,.ipa等
        type: String,
        default: '',
    },
    listType: {
        type: String as PropType<EpPropMergeType<StringConstructor, 'picture-card' | 'text' | 'picture', unknown> | undefined>,
        default: 'picture-card',
    },
    disabled: {
        type: Boolean,
        default: false,
    },
    limit: {
        type: Number,
    },
})

const emits = defineEmits(['update:modelValue', 'change'])
const upload = reactive({
    id: ('my-upload' + new Date().getTime() + '_' + randomInt(1000, 9999)) as string, //用于判断组件是否已经销毁，防止倒计时重复执行
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
                    url: item,
                }
            })
        }
        return [
            {
                name: (props.modelValue as string).slice((props.modelValue as string).lastIndexOf('/') + 1),
                url: props.modelValue as string,
            },
        ]
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
    //生成保存在云服务器中的文件名及完成地址
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
        param: {
            ...props.api?.param,
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
        if (props.multiple) {
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
        if (props.multiple) {
            upload.value.push(file.raw.saveInfo.url)
        } else {
            upload.value = file.raw.saveInfo.url
            upload.fileList = [file]
        }
        emits('update:modelValue', upload.value)
        emits('change')
    },
    onError: (err: Error /* , file: any, fileList: any */) => {
        ElMessage.error(t('common.tip.uploadFail') + '(' + err.message + ')')
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
    },
    getUrl: (file: any): string => (file?.response === undefined ? file.url : file.raw.saveInfo.url),
    copyUrl: async (file: any) => {
        await toClipboard(upload.getUrl(file))
        ElMessage.success(t('common.tip.copySuccess'))
    },
    download: (file: any) => window.open(upload.getUrl(file)),
})

const imageViewer = reactive({
    urlList: computed((): string[] => upload.fileList.map((item) => item.url)),
    initialIndex: 0,
    visible: false,
    close: () => (imageViewer.visible = false),
})

upload.initSignInfo() //初始化签名信息
</script>

<template>
    <div :id="upload.id">
        <div v-if="listType == 'picture-card'" class="upload-container">
            <el-upload
                :ref="(el: any) => upload.ref = el"
                v-model:file-list="upload.fileList"
                :action="upload.action"
                :data="upload.data"
                :before-upload="upload.beforeUpload"
                :on-success="upload.onSuccess"
                :on-error="upload.onError"
                :on-remove="upload.onRemove"
                :multiple="multiple"
                :accept="accept"
                :list-type="listType"
                :disabled="disabled"
                :limit="limit"
                :on-preview="upload.onPreview"
                :drag="true"
                :class="upload.class"
            >
                <template #default>
                    <slot v-if="slots.default" name="default"></slot>
                    <template v-else>
                        <el-icon class="el-icon--upload">
                            <autoicon-ep-upload-filled />
                        </el-icon>
                        <div class="el-upload__text" v-html="t('common.tip.uploadOrDrop')"></div>
                    </template>
                </template>
                <template v-if="slots.trigger" #trigger>
                    <slot name="trigger"></slot>
                </template>
                <template v-if="slots.tip" #tip>
                    <slot name="tip"></slot>
                </template>
                <!-- <template v-if="slots.file" #file="{ file }">
                    <slot name="file" :file="file"></slot>
                </template> -->
                <template #file="{ file }">
                    <slot v-if="slots.file" name="file" :file="file"></slot>
                    <template v-else>
                        <label class="el-upload-list__item-status-label">
                            <el-icon class="el-icon--check">
                                <autoicon-ep-check />
                            </el-icon>
                        </label>

                        <template v-if="showType == 'image'">
                            <img class="el-upload-list__item-thumbnail" :src="file.url" />
                            <span class="el-upload-list__item-actions">
                                <span @click="upload.onPreview(file)"><autoicon-ep-zoom-in /></span>
                                <span @click="upload.copyUrl(file)"><autoicon-ep-document-copy /></span>
                                <span @click="upload.ref.handleRemove(file)"><autoicon-ep-delete /></span>
                            </span>
                            <el-image-viewer v-if="imageViewer.visible" :url-list="imageViewer.urlList" :initial-index="imageViewer.initialIndex" :hide-on-click-modal="true" @close="imageViewer.close" />
                        </template>
                        <template v-else-if="showType == 'video'">
                            <video preload="none" :controls="true" :src="file.url" style="width: 100%" />
                            <el-icon class="el-icon--close" @click="upload.ref.handleRemove(file)">
                                <autoicon-ep-close />
                            </el-icon>
                        </template>
                        <template v-else-if="showType == 'audio'">
                            <audio preload="none" :controls="true" :src="file.url" style="width: 100%; height: 40px; position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%)" />
                            <el-icon class="el-icon--close" @click="upload.ref.handleRemove(file)">
                                <autoicon-ep-close />
                            </el-icon>
                        </template>
                        <template v-else>
                            <el-icon :size="100" style="width: 100%; position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%)"><autoicon-ep-document /></el-icon>
                            <span class="el-upload-list__item-actions">
                                <!-- 刚上传的文件没必要给下载按钮 -->
                                <span v-if="file?.response === undefined" @click="upload.download(file)"><autoicon-ep-download /></span>
                                <span @click="upload.copyUrl(file)"><autoicon-ep-document-copy /></span>
                                <span @click="upload.ref.handleRemove(file)"><autoicon-ep-delete /></span>
                            </span>
                        </template>
                    </template>
                </template>
            </el-upload>
        </div>
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
            :multiple="multiple"
            :accept="accept"
            :list-type="listType"
            :disabled="disabled"
            :limit="limit"
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

.upload-container :deep(.el-upload-list--picture-card .el-upload-list__item .el-icon--close) {
    top: 5px;
    transform: translateY(0);
}

.upload-container :deep(.el-upload-list__item:hover .el-icon--close) {
    display: inline-flex;
}
</style>
