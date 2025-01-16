<script setup lang="tsx">
//const { t } = useI18n()
const keepAliveStore = useKeepAliveStore()
</script>

<template>
    <router-view v-slot="{ Component, route }">
        <transition mode="out-in" name="el-zoom-in-center">
            <keep-alive :include="keepAliveStore.appContainerInclude" :exclude="keepAliveStore.appContainerExclude" :max="keepAliveStore.appContainerMax">
                <component v-if="!keepAliveStore.appContainerExclude.includes(route.meta.componentName as string)" :is="Component" :key="route.fullPath" />
            </keep-alive>
        </transition>
    </router-view>
    <!-- <suspense>
        <template #default>
            <router-view v-slot="{ Component, route }">
                <transition mode="out-in" name="transform">
                    <keep-alive :include="keepAliveStore.appContainerInclude"
                        :exclude="keepAliveStore.appContainerExclude" :max="keepAliveStore.appContainerMax">
                        <component v-if="!keepAliveStore.appContainerExclude.includes(route.meta.componentName as string)" :is="Component"
                            :key="route.fullPath" />
                    </keep-alive>
                </transition>
            </router-view>
        </template>
        <template #fallback>{{ t('common.loading') }}</template>
    </suspense> -->
</template>

<!-- <style scoped>
.transform-enter-active,
.transform-leave-active {
    transition: opacity 0.3s ease;
}

/* .transform-enter, */
.transform-enter-from,
.transform-leave-to {
    opacity: 0;
}
</style> -->
