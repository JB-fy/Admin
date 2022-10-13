<script setup lang="ts">
import { useUserStore } from '@/stores/user';

const userStore = useUserStore()
</script>

<template>
    <RouterView v-slot="{ Component, route }">
        <Transition mode="out-in" name="transform">
            <KeepAlive :include="userStore.cacheRouteInclude" :exclude="userStore.cacheRoute.exclude"
                :max="userStore.cacheRoute.max">
                <component v-if="userStore.cacheRoute.exclude.indexOf(route.path) === -1" :is="Component"
                    :key="route.path" />
            </KeepAlive>
        </Transition>
    </RouterView>
    <!-- <suspense>
        <template #default>
            <component v-if="userStore.cacheRoute.exclude.indexOf(route.path) === -1" :is="Component"
                :key="route.path" />
        </template>
        <template #fallback> Loading... </template>
    </suspense> -->
</template>

<style scoped>
.transform-enter-active,
.transform-leave-active {
    transition: opacity 0.3s ease;
}

/* .transform-enter, */
.transform-enter-from,
.transform-leave-to {
    opacity: 0;
}
</style>