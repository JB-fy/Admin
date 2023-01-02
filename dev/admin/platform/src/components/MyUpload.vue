<script setup lang="ts">
import type { UploadUserFile } from 'element-plus';
import { slice } from 'lodash';

const props = defineProps({
    modelValue: {
        type: [String, Array]
    },
    /**
     * 接口。格式：{ code: string, param: object }
     *      code：非必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：非必须。接口函数所需参数。格式：{ [propName: string]: any }
     */
    api: {
        type: Object,
        default: {
            code: 'upload/sign',
            param: {}
        }
    },
    acceptType: {
        type: Array,
        default: []
    },
    maxSize: {
        type: Number,
        default: 0
    },
    multiple: {
        type: Boolean,
        default: true
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
                return { url: item }
            })
        }
        return [{ url: (<string>props.modelValue) }]
    })(),
    action: '' as string,
    data: {} as { [propName: string]: any },
    signInfo: {} as { [propName: string]: any },    //缓存的签名信息
    save: { //保存的文件名及文件路径
        fileName: (rawFile: any) => {
            return upload.signInfo.dir + rawFile.uid + rawFile.name.slice(rawFile.name.lastIndexOf('.'))
        },
        url: (rawFile: any) => {
            return upload.signInfo.host + '/' + upload.save.fileName(rawFile)
        }
    },
    /* //示例
    signInfo: {
        accessid: "xxxx",
        host: "http://xxxxx.oss-cn-hongkong.aliyuncs.com",
        dir: "common/2022/12/31/1521189152_",
        expire: 1672471578,
        callback: "string",
        policy: "string",
        signature: "string"
    }, */
    api: {
        loading: false,
        getSign: async () => {
            if (upload.api.loading) {
                return
            }
            upload.api.loading = true
            let signInfo = {}
            try {
                const res = await request(props.api.code, props.api.param)
                signInfo = res.data
            } catch (error) { }
            upload.api.loading = false
            return signInfo
        },
    },
    onPreview: (file: any) => {
        dialogImage.url = file.url
        dialogImage.visible = true
    },
    onRemove: (file: any, fileList: UploadUserFile) => {
        console.log(4444)
        //上传前处理函数beforeUpload返回false时也会触发此函数。此时file内没有response，但是由于没上传也不会存在于props.modelValue中，故不影响删除逻辑
        //let url: string = file?.response === undefined ? file.url : file.response.data.url
        let url: string = file?.response === undefined ? file.url : upload.save.url(file.raw)
        if (props.multiple) {
            (<string[]>props.modelValue).splice((<string[]>props.modelValue).indexOf(url), 1)
        } else {
            (<string>props.modelValue) = ''
        }
        emits('change')
        emits('update:modelValue', props.modelValue)
    },
    onSuccess: (res: any, file: any, fileList: any) => {
        console.log(2222)
        /* //file示例：
        {
            "name": "ico_kong.3fd7d5f.png",
            "percentage": 100,
            "status": "success",
            "size": 26131,
            "raw": {
                "uid": 1672479985126
            },
            "uid": 1672479985126,
            "url": "blob:http://192.168.200.200:5173/726cbaa0-dae6-4a35-90fc-31802b74af40",
            "response": "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<Error>\n  <Code>CallbackFailed</Code>\n  <Message>Error status : 502.</Message>\n  <RequestId>63B004F1A96699383854257A</RequestId>\n  <HostId>gamemeta.oss-cn-hangzhou.aliyuncs.com</HostId>\n</Error>\n"
        } */
        if (res.code !== '00000000') {
            if (props.multiple) {
                //(<string[]>props.modelValue).push(res.data.url)
                (<string[]>props.modelValue).push(upload.save.url(file.raw))
            } else {
                (<string>props.modelValue) = upload.save.url(file.raw)
            }
            emits('change')
            emits('update:modelValue', props.modelValue)
        } else {
            ElMessage.error('上传失败，请稍后再试！')
            fileList.splice(fileList.indexOf(file), 1)
        }
    },
    beforeUpload: async (rawFile: any) => {
        if (props.acceptType.length > 0 && props.acceptType.indexOf(rawFile.type) === -1) {
            ElMessage.error('文件格式不在允许范围内！')
            return false
        }
        if (props.maxSize > 0 && props.maxSize < rawFile.size / 1024 / 1024) {
            ElMessage.error('文件大小不在允许范围内！')
            return false
        }
        //判断授权是否失效,失效则重新获取授权, 5s做为缓冲即提前3s更新授权
        if (upload.signInfo.expire > new Date().getTime() / 1000 + 5) {
            //未失效需重新设置文件名
            upload.data.key = upload.save.fileName(rawFile) //这是文件保存路径及文件名，必须唯一，否则会覆盖oss服务器同名文件
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
                callback: upload.signInfo.callback,
                success_action_status: '200', //让服务端返回200,不然，默认会返回204
            }
            upload.data.key = upload.save.fileName(rawFile) //文件的完整保存路径，必须唯一，否则会覆盖服务器同名文件
            return true
        }
        return false
    },
    onExceed: (file: any, fileList: any) => {
        ElMessage.error('最多允许上传（' + props.limit + '）个文件')
    }
})

const dialogImage = reactive({
    url: '',
    visible: false
})
</script>

<template>
    <div class="upload-container">
        <ElUpload :ref="(el: any) => { upload.ref = el }" v-model:file-list="upload.fileList" :action="upload.action"
            :data="upload.data" :before-upload="upload.beforeUpload" :on-success="upload.onSuccess"
            :on-remove="upload.onRemove" :on-preview="upload.onPreview" :on-exceed="upload.onExceed"
            :multiple="multiple" :limit="limit" :drag="true" list-type="picture-card"
            :class="limit > 0 && limit == modelValue?.length ? 'hide' : ''" :style="multiple ? '' : 'height: 148px;'">
            <ElIcon class="el-icon--upload">
                <AutoiconEpUploadFilled />
            </ElIcon>
            <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
            <template #tip>
                <div class="el-upload__tip">
                    jpg/png files with a size less than 500kb
                </div>
            </template>
        </ElUpload>
        <ElDialog v-model="dialogImage.visible" :center="true" :append-to-body="true" top="50px">
            <ElImage style="width: 100%;" :src="dialogImage.url" />
        </ElDialog>
    </div>

    <!-------- 使用示例 开始-------->

    <!-------- 使用示例 结束-------->
</template>

<style scoped>
.upload-container :deep(.el-upload .el-upload-dragger) {
    border: none;
    height: 146px;
    line-height: 18px;
}

.upload-container :deep(.el-upload) {
    width: 146px;
    margin-right: 10px;
}

.upload-container :deep(.hide .el-upload) {
    display: none;
}
</style>