<script setup lang="ts">
const props = defineProps({
    modelValue: {
        type: Array
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
})

const emits = defineEmits(['update:modelValue', 'change'])
//emits('change')
//emits('update:modelValue', val)
const upload = reactive({
    ref: null as any,
    action: '' as string,
    data: {} as { [propName: string]: any },
    signInfo: {} as { [propName: string]: any },
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
    beforeUpload: async (rawFile: any) => {
        /* if (this.mimeTypes.length > 0 && this.mimeTypes.indexOf(rawFile.type) === -1) {
          this.$message({
            type: 'error',
            message: '文件格式不在允许范围内！'
          })
          return false;
        }
        if (this.maxSize < rawFile.size / 1024 / 1024) {
          this.$message({
            type: 'error',
            message: '文件大小不在允许范围内！'
          })
          return false;
        } */
        //判断授权是否失效,失效则重新获取授权, 5s做为缓冲即提前3s更新授权
        if (upload.signInfo.expire > Date.parse(new Date()) / 1000 + 5) {
            //未失效需重新设置文件名
            upload.data.key = upload.signInfo.dir + rawFile.uid + rawFile.name.substring(rawFile.name.lastIndexOf('.')) //这是文件保存路径及文件名，必须唯一，否则会覆盖oss服务器同名文件
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
                //callback: upload.signInfo.callback,
                success_action_status: '200', //让服务端返回200,不然，默认会返回204
            }
            upload.data.key = upload.signInfo.dir + rawFile.uid + rawFile.name.substring(rawFile.name.lastIndexOf('.')) //这是文件保存路径及文件名，必须唯一，否则会覆盖oss服务器同名文件
            return true
        }
        return false
    },
})
</script>

<template>
    <ElUpload :ref="(el: any) => { upload.ref = el }" :action="upload.action" :data="upload.data"
        :beforeUpload="upload.beforeUpload">
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
    </ElUpload>

    <!-- <el-upload
      ref="upload"
      :class="value.length==limit ? 'hide' : ''"
      :style="limit==1 ? 'height: 148px;' : ''"
      list-type="picture-card"
      :multiple="multiple"
      :limit="limit"
      :drag="true"
      :file-list="existFileList"
      :action="aliyun.url"
      :data="aliyun.query"
      :before-upload="beforeUpload"
      :on-success="onSuccess"
      :on-remove="onRemove"
      :on-preview="onPreview"
      :on-exceed="onExceed">
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
    </el-upload>
    <el-dialog :visible.sync="dialogImage.visible" center top="50px" :append-to-body="true">
      <el-image style="width: 100%;" :src="dialogImage.url"></el-image>
    </el-dialog> -->

    <!-------- 使用示例 开始-------->

    <!-------- 使用示例 结束-------->
</template>