<script setup lang="tsx">
const { t } = useI18n()
</script>

<template>
    <div style="text-align: center; font-size: 300px; color: #409eff">
        <ElSpace :size="0">
            <ElIcon class="is-loading">
                <AutoiconEpOrange />
            </ElIcon>
            {{ t('index.welcome') }}
            <ElIcon class="is-loading">
                <AutoiconEpOrange />
            </ElIcon>
        </ElSpace>
    </div>
</template>
