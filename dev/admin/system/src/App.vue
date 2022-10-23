<script setup lang="ts">
const languageStore = useLanguageStore()
</script>

<template>
  <ElConfigProvider :locale="languageStore.elementPlusLocale">
    <RouterView />
  </ElConfigProvider>
</template>