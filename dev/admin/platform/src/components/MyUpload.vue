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
    acceptType: {   //需要限制文件格式时使用。示例：['image/png','image/jpg','image/jpeg','image/gif']
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
})

const emits = defineEmits(['update:modelValue', 'change'])
const upload = reactive({
    ref: null as any,
    fileList: ((): any => {
        if (!props.modelValue) {
            return []
        }
        if (props.multiple) {
            return (<string[]>props.modelValue).map((item) => {
                return {
                    name: item.slice(item.lastIndexOf('/') + 1),
                    url: item
                }
            })
        }
        return [{
            name: (<string>props.modelValue).slice((<string>props.modelValue).lastIndexOf('/') + 1),
            url: (<string>props.modelValue)
        }]
    })(),
    class: computed(() => {
        if (props.multiple) {
            return props.limit && props.limit == upload.fileList.length ? 'hide' : ''
        } else {
            return upload.fileList.length ? 'hide' : ''
        }
    }),
    action: '' as string,
    data: {} as { [propName: string]: any },
    signInfo: {} as { [propName: string]: any },    //缓存的签名信息。示例：{ accessid: "xxxx", host: "https://xxxxx.com", dir: "common/2022/12/31/1521189152_", expire: 1672471578, callback: "string", policy: "string", signature: "string" }
    //生成保存在云服务器中的文件名及完成地址
    createSaveInfo: (rawFile: any) => {
        let fileName = upload.signInfo.dir + rawFile.uid + rawFile.name.slice(rawFile.name.lastIndexOf('.'))
        let url = upload.signInfo.host + '/' + fileName
        return {
            fileName: fileName,
            url: url
        }
    },
    api: {
        loading: false,
        code: props.api?.code ?? 'upload/sign',
        param: {
            ...props.api?.param
        },
        getSign: async () => {
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
    onPreview: (uploadFiles: any) => {
        dialogImage.url = uploadFiles.url
        dialogImage.visible = true
    },
    onRemove: (file: any, fileList: any) => {
        console.log(4444)
        //上传前处理函数beforeUpload返回false时也会触发此函数。此时file内没有response，但是由于没上传也不会存在于props.modelValue中，故不影响删除逻辑
        let url: string = file?.response === undefined ? file.url : file.raw.saveInfo.url
        let value: any = props.modelValue
        if (props.multiple) {
            value.splice(value.indexOf(url), 1)
        } else {
            value = ''
        }
        emits('change')
        emits('update:modelValue', value)
    },
    onSuccess: (res: any, file: any, fileList: any) => {
        if (upload.signInfo?.callback && res.code === '00000000') {    //如有回调服务器且有报错，则默认失败
            //ElMessage.error('上传失败，请稍后再试！')
            fileList.splice(fileList.indexOf(file), 1)
            return
        }
        let value: any = props.modelValue
        if (props.multiple) {
            value.push(file.raw.saveInfo.url)
        } else {
            value = file.raw.saveInfo.url
            upload.fileList = [file]
        }
        emits('change')
        emits('update:modelValue', value)
    },
    beforeUpload: async (rawFile: any) => {
        if (props.acceptType.length > 0 && props.acceptType.indexOf(rawFile.type) === -1) {
            //ElMessage.error('文件格式不在允许范围内！')
            return false
        }
        if (props.maxSize > 0 && props.maxSize < rawFile.size / 1024 / 1024) {
            //ElMessage.error('文件大小不在允许范围内！')
            return false
        }
        //判断授权是否失效,失效则重新获取授权, 5s做为缓冲即提前3s更新授权
        if (upload.signInfo.expire > new Date().getTime() / 1000 + 5) {
            //未失效需重新设置文件名
            rawFile.saveInfo = upload.createSaveInfo(rawFile)
            upload.data.key = rawFile.saveInfo.fileName //这是文件保存路径及文件名，必须唯一，否则会覆盖oss服务器同名文件
            return true
        }

        const signInfo = await upload.api.getSign()
        if (signInfo && Object.keys(signInfo).length) {
            upload.signInfo = { ...signInfo }

            upload.action = upload.signInfo.host
            upload.data = {
                OSSAccessKeyId: upload.signInfo.accessid,
                policy: upload.signInfo.policy,
                signature: upload.signInfo.signature,
                success_action_status: '200', //让服务端返回200,不然，默认会返回204
            }
            upload.signInfo?.callback ? upload.data.callback = upload.signInfo.callback : null //是否回调服务器
            rawFile.saveInfo = upload.createSaveInfo(rawFile)
            upload.data.key = rawFile.saveInfo.fileName //文件的完整保存路径，必须唯一，否则会覆盖服务器同名文件
            return true
        }
        return false
    }
})

const dialogImage = reactive({
    url: '',
    visible: false
})
</script>

<template>
    <div v-if="isImage" class="upload-container">
        <ElUpload :ref="(el: any) => { upload.ref = el }" v-model:file-list="upload.fileList" :action="upload.action"
            :data="upload.data" :before-upload="upload.beforeUpload" :on-success="upload.onSuccess"
            :on-remove="upload.onRemove" :on-preview="upload.onPreview" :multiple="multiple" :limit="limit"
            list-type="picture-card" :drag="true" :class="upload.class">
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
        <ElDialog v-model="dialogImage.visible" :center="true" :append-to-body="true" top="50px">
            <ElImage style="width: 100%;" :src="dialogImage.url" />
        </ElDialog>
    </div>
    <ElUpload v-else :ref="(el: any) => { upload.ref = el }" v-model:file-list="upload.fileList" :action="upload.action"
        :data="upload.data" :before-upload="upload.beforeUpload" :on-success="upload.onSuccess"
        :on-remove="upload.onRemove" :multiple="multiple" :limit="limit" list-type="text">
        <ElButton type="primary">{{ t('common.upload') }}</ElButton>
        <template v-if="tip" #tip>
            <div class="el-upload__tip">
                {{ tip }}
            </div>
        </template>
    </ElUpload>

    <!-------- 使用示例 开始-------->
    <!-- <MyUpload v-model="saveForm.data.avatar" />

    <MyUpload v-model="saveForm.data.avatar" :isImage="false" /> -->
    <!-------- 使用示例 结束-------->
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