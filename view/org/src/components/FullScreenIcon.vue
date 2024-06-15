<script setup lang="tsx">
import screenfull from 'screenfull'

const { t } = useI18n()

const fullScreenIcon = reactive({
    ref: null as any,
    isFullscreen: screenfull.isFullscreen,
    click: () => {
        if (!screenfull.isEnabled) {
            ElMessage.warning(t('common.tip.notFullScreen'))
            return
        }
        screenfull.toggle()
    },
})

screenfull.onchange(() => {
    fullScreenIcon.isFullscreen = screenfull.isFullscreen
})
</script>

<template>
    <van-icon :name="fullScreenIcon.isFullscreen ? 'shrink' : 'expand-o'" @click="fullScreenIcon.click" :size="16.8" />
    <!-- <autoicon-ep-aim v-if="fullScreenIcon.isFullscreen" @click="fullScreenIcon.click" />
    <autoicon-ep-full-screen v-else @click="fullScreenIcon.click" /> -->
</template>
