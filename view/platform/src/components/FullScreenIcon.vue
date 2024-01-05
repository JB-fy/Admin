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
    }
})

screenfull.onchange(() => {
    fullScreenIcon.isFullscreen = screenfull.isFullscreen
})
</script>

<template>
    <VanIcon :name="fullScreenIcon.isFullscreen ? 'shrink' : 'expand-o'" @click="fullScreenIcon.click" :size="16.8" />
    <!-- <AutoiconEpAim v-if="fullScreenIcon.isFullscreen" @click="fullScreenIcon.click" />
    <AutoiconEpFullScreen v-else @click="fullScreenIcon.click" /> -->
</template>
