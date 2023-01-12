<script setup lang="ts">
const { t } = useI18n()
</script>

<template>
    <div style="text-align: center; font-size: 300px; color: #409EFF;">
        <ElSpace :size="0">
            <ElIcon class="is-loading">
                <AutoiconEpOrange />
            </ElIcon>
            {{ t('view.index.welcome') }}
            <ElIcon class="is-loading">
                <AutoiconEpOrange />
            </ElIcon>
        </ElSpace>
    </div>
</template>