<script setup lang="tsx">
import screenfull from 'screenfull'

const { t } = useI18n()

const fullScreenIcon = reactive({
    isFullscreen: screenfull.isFullscreen,
    click: () => (screenfull.isEnabled ? screenfull.toggle() : ElMessage.warning(t('common.tip.notFullScreen'))),
})

screenfull.onchange(() => (fullScreenIcon.isFullscreen = screenfull.isFullscreen))
</script>

<template>
    <autoicon-material-symbols-fullscreen-exit v-if="fullScreenIcon.isFullscreen" @click="fullScreenIcon.click" />
    <autoicon-material-symbols-fullscreen v-else @click="fullScreenIcon.click" />
</template>
