<script setup lang="ts">
const props = defineProps({
    icon: {
        type: String,
        required: true,
        default: ''
    }
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
    <ElIcon v-if="icon">
        <VanIcon v-if="prefix === 'vant'" :name="iconCode" />
        <component v-else :is="iconCode" />
        <!-- <component v-else-if="icon.indexOf('ep-') === 0" :is="iconCode" />
        <component v-else-if="icon.indexOf('autoicon-') === 0" :is="iconCode" /> -->
    </ElIcon>
</template>