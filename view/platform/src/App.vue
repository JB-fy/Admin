<script setup lang="tsx">
const languageStore = useLanguageStore()
</script>

<template>
    <el-config-provider :locale="languageStore.elementPlusLocale" size="default" :button="{ autoInsertSpace: true }" :message="{ max: 5 }">
        <router-view />
    </el-config-provider>
</template>
