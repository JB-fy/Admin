<script setup lang="ts">
const props = defineProps({
    icon: {
        type: String,
        required: true,
        default: '', //常用格式：{前缀}{集合}{标识}。参考https://github.com/antfu/unplugin-icons。示例：AutoiconEpGuide。其中Autoicon为vite.config.ts中自定义的前缀，Ep为Element Plus集合，Guide为标识；Vant格式：Vant-{标识}。示例：Vant-like。
    },
    color: {
        type: String,
        default: '',
    },
    size: {
        type: [Number, String],
        default: '',
    },
})
const prefix = computed(() => {
    return props.icon.slice(0, props.icon.indexOf('-')).toLowerCase()
})
const iconCode = computed(() => {
    if (prefix.value === 'vant') {
        return props.icon.slice(props.icon.indexOf('-') + 1)
    }
    return props.icon
})
</script>

<template>
    <ElIcon v-if="icon" :color="color" :size="size">
        <VanIcon v-if="prefix === 'vant'" :name="iconCode" />
        <component v-else :is="iconCode" />
        <!-- <component v-else-if="icon.indexOf('ep-') === 0" :is="iconCode" />
        <component v-else-if="icon.indexOf('autoicon-') === 0" :is="iconCode" /> -->
    </ElIcon>
</template>
