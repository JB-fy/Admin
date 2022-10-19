<script setup lang="ts">
import { useLanguageStore } from '@/stores/language';

const { t } = useI18n()
const languageStore = useLanguageStore()
</script>

<template>
  <!-- <RouterView /> -->
  <!-- <Suspense></Suspense> -->
  <ElConfigProvider :locale="languageStore.elementPlusLocale">
    <span @click="languageStore.changeLanguage('zh-cn')">
      {{ t('config.language.zh-cn') }}
    </span>
    <span @click="languageStore.changeLanguage('en')">
      {{ t('config.language.en') }}
    </span>
    <RouterView />
  </ElConfigProvider>
</template>

<style scoped>
#app {
  /* font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;*/
  /* min-height: 100vh;
  text-align: center; 
  margin: 0 auto;
  background-color: #f7f8fa; */
}

/*--------解决用户代理样式表（浏览器提供的默认样式）设置的margin: 8px; 开始--------*/
body {
  margin: 0 !important;
}

/*--------解决用户代理样式表（浏览器提供的默认样式）设置的margin: 8px; 结束--------*/
</style>
