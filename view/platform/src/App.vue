<script setup lang="tsx">
const languageStore = useLanguageStore()
</script>

<template>
    <ElConfigProvider :locale="languageStore.elementPlusLocale" size="default" :button="{ autoInsertSpace: true }" :message="{ max: 5 }">
        <RouterView />
    </ElConfigProvider>
</template>
