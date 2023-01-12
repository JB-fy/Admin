<script setup lang="ts">
const route = useRoute()

const thirdSiteUrl:string = route.query.url
</script>

<template>
    <iframe :src="thirdSiteUrl" frameborder="0" style="width: 100%; height: calc(100vh - 194px);"></iframe>
</template>