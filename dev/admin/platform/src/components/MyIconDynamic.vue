<script setup lang="ts">
const props = defineProps({
    icon: {
        type: String,
        required: true,
        default: '' //参数示例：Element Plus(AutoiconEpXxxx或EpXxxx)；Vant(Vant-xxxx)；其他(AutoiconXxxx，参考https://github.com/antfu/unplugin-icons)
    },
    color: {
        type: String,
        default: ''
    },
    size: {
        type: [Number, String],
        default: ''
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