<script setup lang="tsx">
const props = defineProps({
    icon: {
        //常用格式：{前缀}-{集合}-{标识}。参考https://github.com/antfu/unplugin-icons。示例：autoicon-ep-guide。其中autoicon为vite.config.ts中自定义的前缀，ep为Element Plus集合，guide为标识；vant格式：vant-{标识}。示例：vant-like。
        type: String,
        required: true,
    },
})

const prefix = computed(() => props.icon.slice(0, props.icon.indexOf('-')).toLowerCase())

const iconCode = computed(() => (prefix.value === 'vant' ? props.icon.slice(props.icon.indexOf('-') + 1) : props.icon))
</script>

<template>
    <el-icon v-if="icon">
        <van-icon v-if="prefix === 'vant'" :name="iconCode" />
        <component v-else :is="iconCode" />
        <!-- <component v-else-if="icon.indexOf('ep-') === 0" :is="iconCode" />
        <component v-else-if="icon.indexOf('autoicon-') === 0" :is="iconCode" /> -->
    </el-icon>
</template>
