<script setup lang="tsx">
import iconMap from './my-icon-dynamic'

const props = defineProps({
    icon: {
        //地址或标识。标识需在前端代码文件my-icon-dynamic.tsx中添加对应标识才能生效，格式：autoicon-{集合}-{标识}，参考：https://icones.js.org。示例：autoicon-ep-guide
        type: String,
        required: true,
    },
})
</script>

<template>
    <el-icon v-if="icon && (props.icon.indexOf('http') === 0 || iconMap[props.icon])" v-bind="$attrs">
        <img v-if="props.icon.indexOf('http') === 0" :src="props.icon" style="max-width: 100%; max-height: 100%" />
        <component v-else :is="iconMap[props.icon]" />
    </el-icon>
</template>
