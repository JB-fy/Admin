<script setup lang="tsx">
const route = useRoute()

const thirdSiteUrl = route.query.url as string
</script>

<template>
    <iframe :src="thirdSiteUrl" frameborder="0" style="width: 100%; height: calc(100vh - 194px)"></iframe>
</template>
