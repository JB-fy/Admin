<script setup lang="ts">
//const { t } = useI18n()
const keepAliveStore = useKeepAliveStore()
</script>

<template>
    <RouterView v-slot="{ Component, route }">
        <Transition mode="out-in" name="el-zoom-in-center">
            <KeepAlive :include="keepAliveStore.appContainerInclude" :exclude="keepAliveStore.appContainerExclude"
                :max="keepAliveStore.appContainerMax">
                <component v-if="keepAliveStore.appContainerExclude.indexOf(route.fullPath) === -1" :is="Component"
                    :key="route.fullPath" />
            </KeepAlive>
        </Transition>
    </RouterView>
    <!-- <Suspense>
        <template #default>
            <RouterView v-slot="{ Component, route }">
                <Transition mode="out-in" name="transform">
                    <KeepAlive :include="keepAliveStore.appContainerInclude"
                        :exclude="keepAliveStore.appContainerExclude" :max="keepAliveStore.appContainerMax">
                        <component v-if="keepAliveStore.appContainerExclude.indexOf(route.fullPath) === -1" :is="Component"
                            :key="route.fullPath" />
                    </KeepAlive>
                </Transition>
            </RouterView>
        </template>
        <template #fallback>{{ t('common.loading') }}</template>
    </Suspense> -->
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