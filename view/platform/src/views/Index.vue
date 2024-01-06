<script setup lang="tsx">
const { t } = useI18n()
</script>

<template>
    <div style="text-align: center; font-size: 300px; color: #409eff">
        <el-space :size="0">
            <el-icon class="is-loading">
                <autoicon-ep-orange />
            </el-icon>
            {{ t('index.welcome') }}
            <el-icon class="is-loading">
                <autoicon-ep-orange />
            </el-icon>
        </el-space>
    </div>
</template>
